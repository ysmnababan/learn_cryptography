# Manual mTLS Demonstration

This project demonstrates Mutual TLS (mTLS) authentication by manually creating certificates.

## Prerequisites

- [Go](https://golang.org/) installed
- Bash shell (for running scripts)

## Steps to Run

1. **Generate Certificates**

    Run the following command to create the necessary certificates:
    ```bash
    bash key.sh
    ```

2. **Start the Server**

    Launch the server using:
    ```bash
    go run . --mode=server
    ```

3. **Run the Client**

    In a separate terminal, start the client:
    ```bash
    go run . --mode=client
    ```

## Expected Output

### Server

The server will display details about the incoming request and TLS state:
```text
2025/09/10 14:28:29 >>>>>>>>>>>>>>>> Header <<<<<<<<<<<<<<<<
2025/09/10 14:28:29 Accept-Encoding:gzip
2025/09/10 14:28:29 User-Agent:Go-http-client/1.1
2025/09/10 14:28:29 >>>>>>>>>>>>>>>> State <<<<<<<<<<<<<<<<
2025/09/10 14:28:29 Version: 304
2025/09/10 14:28:29 HandshakeComplete: true
2025/09/10 14:28:29 DidResume: false
2025/09/10 14:28:29 CipherSuite: 1301
2025/09/10 14:28:29 NegotiatedProtocol:
2025/09/10 14:28:29 NegotiatedProtocolIsMutual: true
2025/09/10 14:28:29 Certificate chain:
2025/09/10 14:28:29  0 s:/C=[SO]/ST=[Earth]/L=[Mountain]/O=[Client-A]/OU=[Client-A-OU]/CN=localhost
2025/09/10 14:28:29    i:/C=[SO]/ST=[Earth]/L=[Mountain]/O=[MegaEase]/OU=[MegaCloud]/CN=localhost
```

### Client

The client will display the server's response:
```text
2025/09/10 14:28:29 Response body: nice !!!
```

## Notes

- Ensure both server and client use the generated certificates for successful mTLS authentication.
- Modify certificate details in `key.sh` as needed for your environment.
