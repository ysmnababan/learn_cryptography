# Heartbleed Simulation Demo

This project provides a **safe, self-contained simulation** of the "Heartbleed" class of bug.  
It does **not** use TLS/OpenSSL and **does not** attack real software.  
Instead, it demonstrates how an unchecked length in a heartbeat-like echo protocol can cause an out-of-bounds read that leaks adjacent memory.

## What is Heartbleed?

**Heartbleed** is a famous security vulnerability that affected OpenSSL in 2014.  
It exploited a flaw in the TLS "heartbeat" extension, where a client could request more data than it actually sent.  
Because the server did not properly check the length, it would respond with extra bytes from its memory—potentially leaking sensitive information.

The "heartbeat" protocol is designed to keep connections alive by echoing back data.  
Heartbleed showed how failing to validate input lengths in such protocols can lead to serious data leaks.

## How It Works

The simulation consists of a server and a client:

- The **server** echoes back data sent by the client.
- If run in vulnerable mode, the server does **not** check the requested echo length, potentially leaking extra memory.
- If run in safe mode, the server clamps the echo length to the actual payload size, preventing leaks.

## Usage

Open **two terminals**:

### 1. Start the Vulnerable Server

```sh
go run heartbleed_sim.go -mode=server
```

### 2. Run the Client (Requesting Oversized Echo)

```sh
go run heartbleed_sim.go -mode=client -len=200 -payload=hi
```

You should see the client print the original payload plus extra bytes, including the simulated "secret".  
This demonstrates the leak.

### 3. Start the Safe Server (with Bounds Checking)

```sh
go run heartbleed_sim.go -mode=server -safe
```

Re-run the client.  
The leak should **disappear**.

## Flags

| Flag                | Description                                 |
|---------------------|---------------------------------------------|
| `-mode=server|client` | Run as server or client                    |
| `-addr=127.0.0.1:4444` | TCP address to listen/dial                 |
| `-safe`               | (server) Enable fix (length clamped)       |
| `-len=200`            | (client) Requested echo length             |
| `-payload=hi`         | (client) Payload bytes actually sent       |

## Security & Ethics

> **For EDUCATIONAL DEMONSTRATION ONLY** in a closed lab environment.  
> Do **not** point at real systems or networks you do not own or have explicit permission to test.  
> This is **not** an exploit and does **not** interact with TLS.

---

**Note:**  
This simulation is intended to help understand how unchecked input can lead to memory disclosure vulnerabilities like Heartbleed.  
Always follow responsible security practices.