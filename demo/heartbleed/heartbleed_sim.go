package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"strings"
	"time"
)

// Simple heartbeat-like message:
// 1 byte type (1=request, 2=response)
// 2 bytes big-endian length (N)
// N bytes payload

const (
	msgHeartbeatReq  = 1
	msgHeartbeatResp = 2
)

var (
	mode     = flag.String("mode", "server", "server or client")
	addr     = flag.String("addr", "127.0.0.1:4444", "listen/dial address")
	safe     = flag.Bool("safe", false, "server: enable bounds checking (the fix)")
	clen     = flag.Int("len", 200, "client: requested echo length (can exceed payload)")
	cpayload = flag.String("payload", "hi", "client: payload string actually sent")
)

// In a real bug, the leaked bytes come from adjacent heap memory. Here we
// simulate a process heap that has some secrets in it.
var processHeap []byte

func initHeap() {
	secret := []byte("TOP-SECRET: api_key=sk_demo_1234; password=hunter2; session=abcd\n")
	filler := make([]byte, 8*1024)
	if _, err := rand.Read(filler); err != nil {
		panic(err)
	}
	// Put the secret somewhere near the start so it's likely to be leaked.
	processHeap = append([]byte{}, secret...)
	processHeap = append(processHeap, filler...)
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	initHeap()

	switch *mode {
	case "server":
		if err := runServer(*addr, *safe); err != nil {
			log.Fatal(err)
		}
	case "client":
		if err := runClient(*addr, *clen, []byte(*cpayload)); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown mode: %s", *mode)
	}
}

func runServer(addr string, safe bool) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Printf("[server] listening on %s (safe=%v)", addr, safe)
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handleConn(conn, safe)
	}
}

func handleConn(c net.Conn, safe bool) {
	defer c.Close()
	_ = c.SetDeadline(time.Now().Add(5 * time.Minute))
	br := bufio.NewReader(c)

	for {
		// Read header: type (1 byte)
		typ, err := br.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Printf("[server] read type error: %v", err)
			}
			return
		}
		// Read length (2 bytes)
		lenBuf := make([]byte, 2)
		if _, err := io.ReadFull(br, lenBuf); err != nil {
			log.Printf("[server] read len error: %v", err)
			return
		}
		N := int(binary.BigEndian.Uint16(lenBuf))

		// Read payload (actual payload size = M)
		payload := make([]byte, 0, N)
		// For the request we don't know M from header; we have to delimit by N? In heartbleed bug,
		// the server trusts N and reads M exactly as sent by TCP framing. Here we simulate: client
		// sends M bytes; we read only what's available up to M (<= N).
		// To keep it simple, read the next M bytes based on available data in the socket buffer.
		// We'll first peek how much data is available; but bufio doesn't expose it reliably.
		// Instead, read up to min(N, 4096) and allow short read to simulate M < N.
		M := br.Buffered()
		if M == 0 {
			// if nothing buffered, read at least something with a small timeout
			_ = c.SetReadDeadline(time.Now().Add(50 * time.Millisecond))
			b := make([]byte, int(math.Min(float64(N), 4096)))
			n, _ := br.Read(b)
			_ = c.SetReadDeadline(time.Time{})
			payload = append(payload, b[:n]...)
		} else {
			b := make([]byte, int(math.Min(float64(N), float64(M))))
			n, _ := br.Read(b)
			payload = append(payload, b[:n]...)
		}
		M = len(payload)

		if typ != msgHeartbeatReq {
			log.Printf("[server] unknown msg type: %d", typ)
			return
		}

		// VULNERABLE behavior: respond with N bytes, copying only M and then blindly
		// appending bytes from adjacent memory to satisfy N. The fix is to clamp N to M.
		respLen := N
		if safe {
			if respLen > M {
				respLen = M
			}
		}

		resp := make([]byte, 1+2+respLen)
		resp[0] = msgHeartbeatResp
		binary.BigEndian.PutUint16(resp[1:3], uint16(respLen))

		copy(resp[3:], payload)
		if respLen > len(payload) {
			leak := overreadFromHeap(respLen - len(payload))
			copy(resp[3+len(payload):], leak)
		}

		if _, err := c.Write(resp); err != nil {
			log.Printf("[server] write error: %v", err)
			return
		}
	}
}

// overreadFromHeap returns bytes from our simulated process heap.
func overreadFromHeap(n int) []byte {
	if n <= 0 {
		return nil
	}
	if n > len(processHeap) {
		n = len(processHeap)
	}
	return processHeap[:n]
}

func runClient(addr string, reqLen int, payload []byte) error {
	if reqLen < 0 || reqLen > 65535 {
		return fmt.Errorf("len must be 0..65535")
	}
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Printf("[client] connected to %s", addr)

	// Build request: type=1, len=reqLen, payload=payload (shorter than reqLen)
	msg := make([]byte, 1+2+len(payload))
	msg[0] = msgHeartbeatReq
	binary.BigEndian.PutUint16(msg[1:3], uint16(reqLen))
	copy(msg[3:], payload)

	if _, err := conn.Write(msg); err != nil {
		return err
	}

	// Read response header
	hdr := make([]byte, 3)
	if _, err := io.ReadFull(conn, hdr); err != nil {
		return err
	}
	if hdr[0] != msgHeartbeatResp {
		return fmt.Errorf("unexpected response type: %d", hdr[0])
	}
	respLen := int(binary.BigEndian.Uint16(hdr[1:3]))
	buf := make([]byte, respLen)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return err
	}

	// Pretty-print: show printable runes, escape others
	printable := sanitize(buf)
	log.Printf("[client] requested len=%d, sent payload=%q (%d bytes)", reqLen, payload, len(payload))
	fmt.Println("------ leaked bytes ------")
	fmt.Println(printable)
	fmt.Println("--------------------------")
	return nil
}

func sanitize(b []byte) string {
	sb := strings.Builder{}
	for _, c := range b {
		if c >= 32 && c <= 126 { // printable ASCII
			sb.WriteByte(c)
		} else if c == '\n' || c == '\t' || c == '\r' {
			sb.WriteByte(c)
		} else {
			sb.WriteString("\\x")
			sb.WriteString(fmt.Sprintf("%02x", c))
		}
	}
	return sb.String()
}
