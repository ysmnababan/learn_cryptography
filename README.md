# Security Learning Roadmap (for Go Engineers)

This roadmap is for backend engineers who currently "just use" security libraries (JWT, bcrypt, AES, RSA) as boilerplate, but want to **understand the concepts** and **choose the right tools confidently**.  

It’s structured as a **6-month plan** with resources, weekly goals, and expected outcomes.

---

## 📆 6-Month Study Plan

### Phase 1 — Foundations (Weeks 1-4)

**Goal:** Understand core cryptography terms and categories.

- **Week 1: Security Basics**
  - Read: OWASP Top Ten overview → <https://owasp.org/www-project-top-ten/>
  - Watch: Computerphile videos (hashing, encryption, RSA, signatures)
  - Task: Write a personal glossary of security terms

- **Week 2: Hashing**
  - Read: *Crypto 101* (Ch. 1–2)
  - Practice: `crypto/sha256` and `bcrypt` in Go
  - Compare: SHA256 vs bcrypt for passwords

- **Week 3: Symmetric Encryption**
  - Read: *Crypto 101* (Ch. 3)
  - Practice: Encrypt/decrypt with AES-GCM in Go
  - Explore: AES-GCM vs AES-CBC

- **Week 4: Asymmetric Encryption & Signing**
  - Read: *Crypto 101* (Ch. 4–5)
  - Practice: RSA sign/verify in Go
  - Try: ECDSA in Go and compare

---

### Phase 2 — Applying to Web Systems (Weeks 5-8)

**Goal:** Learn how primitives are applied in real systems.

- **Week 5: Password Storage**
  - Read: OWASP Password Storage Cheat Sheet
  - Practice: bcrypt vs argon2id in Go
  - Project: CLI to hash & verify passwords

- **Week 6: JWT and Session Security**
  - Read: Auth0 JWT Best Practices
  - Practice: Implement JWT with HS256 and RS256
  - Compare: HMAC vs RSA signing

- **Week 7: TLS and HTTPS**
  - Read: Cloudflare blog “How HTTPS Works”
  - Practice: HTTPS server in Go with self-signed cert
  - Learn: certificate chains and trust

- **Week 8: OAuth2 & OpenID**
  - Read: *OAuth 2.0 Simplified* (Aaron Parecki)
  - Watch: Okta Dev YouTube tutorials
  - Project: Small OAuth2 client (GitHub login)

---

### Phase 3 — Secure Design Thinking (Weeks 9-12)

**Goal:** Move from "I know tools" to "I can design securely."

- **Week 9: Threat Modeling**
  - Read: OWASP Threat Modeling Cheat Sheet
  - Task: Threat model a sample API (e.g., banking app)

- **Week 10: Vulnerabilities**
  - Read: OWASP cheat sheets (SQL Injection, CSRF, XSS)
  - Project: Build a vulnerable Go API, then fix it

- **Week 11: Secure Coding in Go**
  - Read: Go Security Best Practices → <https://golang.org/doc/security>
  - Practice: `math/rand` vs `crypto/rand`
  - Learn: timing attacks and `subtle.ConstantTimeCompare`

- **Week 12: Defensive Strategies**
  - Study: least privilege, defense in depth, fail secure
  - Add: rate limiting, CSRF protection, secure headers to an API

---

### Phase 4 — Advanced & Practice (Weeks 13-20)

**Goal:** Build & break real cryptosystems.

- **Weeks 13-14: File & Data Encryption**
  - Project: CLI to encrypt/decrypt files with AES-GCM
  - Add: password-based key derivation (PBKDF2/argon2id)

- **Weeks 15-16: Digital Signatures**
  - Project: “secure updater” → sign a file with RSA/ECDSA, verify before run
  - Compare: RSA vs ECDSA performance

- **Weeks 17-18: Attack & Defense**
  - Tool: hashcat (crack MD5/SHA1 passwords)
  - Lesson: why bcrypt/argon2id resist brute force

- **Weeks 19-20: CTF Challenges**
  - Site: <https://cryptopals.com/>
  - Goal: Solve at least first 2–3 sets of challenges

---

### Phase 5 — System-Level Knowledge (Weeks 21-24)

**Goal:** Think like a security engineer.

- **Week 21: Case Studies**
  - Read: postmortems (JWT “alg:none”, Heartbleed, Equifax)
  - Lesson: how crypto misuse leads to real breaches

- **Week 22: System Security Principles**
  - Read: *Security Engineering* (Ross Anderson), focus on protocols & identity

- **Week 23: Full-System Hardening**
  - Project: Harden an old Go API:
    - TLS
    - secure headers
    - JWT
    - password hashing
    - rate limiting

- **Week 24: Review & Summarize**
  - Task: Write a personal security cheat sheet
  - Include: glossary, safe defaults, library choices

---

## 🎯 Outcomes (What You’ll Gain)

By the end of this roadmap you will:

- **Understand**:
  - Hashing vs encryption vs signing
  - When to use bcrypt vs argon2
  - Why AES-GCM > AES-CBC
  - RSA vs ECDSA vs HMAC
  - How TLS/HTTPS really works

- **Be able to choose tools**:
  - Decide symmetric vs asymmetric encryption
  - Select proper password hashing algorithm & cost factors
  - Pick JWT signing method safely

- **Build securely in Go**:
  - Password hashing (bcrypt, argon2)
  - JWT signing & verification
  - AES encryption/decryption
  - HTTPS server with certs
  - Secure Go API (XSS/CSRF/SQLi defense)

- **Think like a security engineer**:
  - Do threat modeling
  - Spot insecure defaults
  - Understand CVEs & advisories

- **Portfolio Projects**:
  - Password hasher CLI
  - File encryption tool (AES-GCM + KDF)
  - JWT demo (HS256 vs RS256)
  - Hardened Go API
  - Crypto CTF challenge solutions

---

## 📚 Key Resources

- Free book: [Crypto 101](https://crypto101.io/)  
- Book: *Cryptography Engineering* (Ferguson, Schneier, Kohno)  
- Book: *Web Security for Developers* (Malcolm McDonald)  
- Book: *Security Engineering* (Ross Anderson)  
- OWASP Cheat Sheet Series → <https://cheatsheetseries.owasp.org/>  
- Go-specific: *Practical Cryptography With Go* (Kyle Isom)  
- Crypto challenges: <https://cryptopals.com/>  

---

## 🚀 Next Steps After This Roadmap

- Join CTFs to keep practicing
- Follow Cryptography StackExchange
- Read real-world postmortems of breaches
- Contribute to open-source security libraries in Go
