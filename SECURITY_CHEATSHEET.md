# Security Cheat Sheet (Go Edition)

A quick reference for choosing the right security tool or approach in Go projects.  
Use this alongside the full roadmap in `README.md`.

---

## 🔑 Core Concepts

- **Hashing** = one-way, irreversible  
  - Use for: password storage, integrity checks
- **Encryption** = reversible, with a key  
  - Use for: protecting confidential data
- **Signing** = proof of authenticity + integrity  
  - Use for: JWTs, verifying updates, digital signatures
- **Symmetric vs Asymmetric**:
  - Symmetric (AES): same key to encrypt/decrypt
  - Asymmetric (RSA/ECDSA): public/private key pair

---

## 📝 Hashing

| Use Case             | Algorithm            | Notes |
|----------------------|----------------------|-------|
| Password storage     | `bcrypt` or `argon2id` | Argon2id recommended, bcrypt still widely used. |
| Integrity check      | `SHA256`, `SHA512`   | Fast hashes, not safe for passwords. |
| HMAC (auth hash)     | `HMAC-SHA256`        | Used in JWT HS256, API key signing. |

**Go libs**:
```go
import "golang.org/x/crypto/bcrypt"
import "golang.org/x/crypto/argon2"
```

---

## 🔒 Symmetric Encryption (AES)

| Algorithm  | Use Case             | Notes |
|------------|----------------------|-------|
| AES-GCM    | Preferred choice     | Provides confidentiality + integrity. |
| AES-CBC    | Legacy only          | Needs extra integrity (HMAC). Avoid new usage. |

**Go libs**:
```go
import "crypto/aes"
import "crypto/cipher"
```

---

## 🔑 Asymmetric Encryption

| Algorithm | Use Case                | Notes |
|-----------|-------------------------|-------|
| RSA       | JWT signing, signatures | Slower, older, widely supported. |
| ECDSA     | JWT signing, signatures | Faster, smaller keys. |
| Ed25519   | Modern alternative      | Safer defaults, simpler API. |

**Go libs**:
```go
import "crypto/rsa"
import "crypto/ecdsa"
import "crypto/ed25519"
```

---

## 🪙 JWT (JSON Web Tokens)

| Algorithm | When to Use | Notes |
|-----------|-------------|-------|
| HS256     | Simple, one secret key shared | Both parties must keep the secret. |
| RS256     | Public/private keys | Server signs with private key, clients verify with public key. |
| ES256     | Uses ECDSA | More modern, smaller keys than RSA. |

**Guidelines**:
- Don’t put sensitive data in JWT payload (it’s just base64).
- Set short expiry times.
- Always validate `alg` field.
- Prefer RS256/ES256 in distributed systems.

---

## 🌐 TLS / HTTPS

- Always use TLS 1.2+ (ideally 1.3).  
- Use Go’s `crypto/tls` instead of writing your own.  
- Certificates:
  - Dev: self-signed
  - Prod: Let’s Encrypt or trusted CA  

**Go code**:
```go
import "crypto/tls"
srv := &http.Server{
    Addr:      ":443",
    TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
}
srv.ListenAndServeTLS("cert.pem", "key.pem")
```

---

## 🛡️ Secure Coding in Go

- Use `crypto/rand` for keys/tokens, **not** `math/rand`.
- Compare secrets with `subtle.ConstantTimeCompare`.
- Don’t roll your own crypto — use `crypto/*` packages.
- Always check errors (crypto errors are critical).
- Sanitize all input → protect against SQLi, XSS, etc.
- Use context timeouts to avoid DoS.

---

## ✅ Safe Defaults (Quick Decisions)

- **Passwords** → `argon2id` (preferred) or `bcrypt` with cost ≥ 12  
- **API auth (shared secret)** → `HMAC-SHA256`  
- **Distributed auth (JWT)** → `RS256` or `ES256`  
- **Encrypting data** → `AES-GCM` with 32-byte key  
- **Digital signatures** → `Ed25519` (modern), or `ECDSA` if required  
- **Transport** → TLS 1.3 with strong ciphers  

---

## 📚 References

- OWASP Cheat Sheets → <https://cheatsheetseries.owasp.org/>
- Crypto 101 → <https://crypto101.io/>
- Go crypto docs → <https://pkg.go.dev/crypto>
- Cryptopals challenges → <https://cryptopals.com/>
