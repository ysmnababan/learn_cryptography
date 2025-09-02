# Key Rotation Pattern

## 1. What Kinds of Keys/Passwords Require Rotation—and Which Don’t?

### Keys/Passwords You Should Rotate

- **Cryptographic keys:**  
    Symmetric keys (used for bulk data encryption) and asymmetric keys (used for signing/encryption) benefit from rotation to limit exposure and cryptanalysis risk.
- **Certificates and SSH keys:**  
    Especially those in long-term use.
- **API keys and tokens:**  
    These can leak, so frequent rotation mitigates long-lived exposure.
- **System/user passwords:**  
    Depending on sensitivity and context, rotation may still be good practice.

### Keys/Passwords You May Not Need to Rotate Frequently

- **Static passwords (e.g., for console access):**  
    NIST now recommends rotation only if there's suspicion of compromise, considering the burdens of frequent mandatory changes.
- **Master/root keys:**  
    Well-protected master/root keys (but still need strong protection controls).

---

## 2. Integrating a New Key into Other Systems

**With a centralized Key Management System (KMS):**

1. Rotate or generate the new key via your KMS.
2. Update aliases or key references in your systems/app configs to point to the new key.
3. Many KMS (e.g., AWS KMS) handle decryption automatically using the correct key version—you don’t need to handle it manually.

---

## 3. Recommended Automated Key Management Tools (Considering Cost & Community Support)

**Top choices:**

- **HashiCorp Vault:**  
    Open-source, self-hosted, strong community adoption.
- **AWS KMS:**  
    Fully managed, automatic rotation, deep integration with AWS ecosystem.
- **Azure Key Vault, Google Cloud KMS:**  
    Similarly integrated with their respective cloud ecosystems.
- **Venafi Trust Protection Platform:**  
    Enterprise-grade, feature-rich.

**Price Snapshot:**

- **Google Cloud KMS:** ~$0.06–$3/month per active key version (priced by protection level), and some crypto op fees. Rotation admin ops are free.
- **AWS KMS:** Modest rotation-specific costs (billed for first and second rotation), but free afterward.
- **Vault:** Open-source with no licensing cost—but self-hosting incurs ops overhead.
- **Venafi:** Likely costly, enterprise-focused.

---

## 4. Automation Scope: What Can Be Automated?

**Automation includes:**

- Key generation and rotation, based on schedules or triggers (time-based or on-demand).
- Updating applications—when using a KMS with alias support, the app just points to alias; rotation happens seamlessly.
- Notification & logging—monitoring is often built-in.

**What might remain manual:**

- Legacy systems lacking API integration may require manual key updates.
- Final configuration validation or redeployment might need manual approval.

---

## 5. End-to-End Key Rotation Flow

**Step-by-step flow:**

1. **Define Policy:** Determine which keys rotate, frequency, triggers (time-based or event-based).
2. **Select Tool:** Choose KMS/Vault considering budget, integration, and community support.
3. **Set Up KMS:** Configure rotation schedules or triggers.
4. **Configure Applications:** Use key aliases or secrets fetch logic.
5. **Automate:** Set rotation schedule and pipeline integration.
6. **Rotate:** KMS issues a new key; apps fetch it via alias.
7. **Validate:** Ensure new key works for encryption/decryption.
8. **Old Key Handling:** Keep old key versions for decryption, then retire when safe.
9. **Audit & Monitor:** Monitor logs, set alerts on anomalies.
10. **Disaster Recovery:** Ensure keys are backed up securely, with recovery procedures tested.

---

## 6. Example for rotation Policy

| **Key / Secret**                       | **Needs Rotation?** | **Best Practice**                   | **Auto or Manual**           |
| -------------------------------------- | ------------------- | ----------------------------------- | ---------------------------- |
| `jwt.secret` (JWT signing key)         | Yes                 | Rotate every 60–90 days or on event | Manual (graceful transition) |
| `dbpass` (each database)               | Yes                 | Frequent rotation                   | Automatic via secrets tools  |
| `appkey`, `appsecret`, etc.            | Yes                 | Rotate frequently                   | Automatic if supported       |
| `redis.password`                       | Yes                 | Rotate frequently                   | Automatic if supported       |
| `accesskeyid`, `secretkeyid` (storage) | Yes                 | Frequent rotation                   | Automatic if supported       |
| `server`, `port`, `env`, etc.          | No                  | Not sensitive                       | N/A                          |


--- 

## 7. Vault Model and Sequence Diagram
```
                ┌──────────────────────────┐
                │      HashiCorp Vault     │
                │   (centralized secrets)  │
                └─────────────▲────────────┘
                              │
                        (secure channel)
                              │
          ┌───────────────────┴───────────────────┐
          │                                       │
┌────────────────────┐                  ┌─────────────────────┐
│    Vault Agent     │                  │     Vault Agent     │
│   (sidecar/local)  │                  │   (sidecar/local)   │
└─────────▲──────────┘                  └──────────▲──────────┘
          │                                        │
   ┌──────┴──────┐                          ┌──────┴──────┐
   │   App A     │                          │   App B     │
   │ (JWT, Redis)│                          │ (DB client) │
   └─────────────┘                          └─────────────┘
```

    --- CASE 1: DB DYNAMIC SECRETS ---
```mermaid
sequenceDiagram
    
    autonumber
    Note over App,DB: Case 1: Database Dynamic Secrets
    App->>Agent: Request DB credentials
    Agent->>Vault: Auth to Vault (JWT/K8s/Token)
    Vault->>DB: Create ephemeral user + password (valid X min)
    DB-->>Vault: Returns user/password
    Vault-->>Agent: Return ephemeral credentials
    Agent-->>App: App connects with temp DB creds
    Note over App,DB: Credentials expire automatically after TTL<br/>Vault can renew if app still needs them
```

    --- CASE 2: DB STATIC SECRETS ---
```mermaid
sequenceDiagram
    
    autonumber
    Vault->>DB: ALTER USER app_user WITH PASSWORD 'new-pass'
    DB-->>Vault: Password updated
    Vault-->>Agent: Provide updated password
    Agent-->>App: App reconnects using new password
    Note over App,DB: Vault periodically rotates static passwords<br/>App must reload them via Agent
```

    --- CASE 3: REDIS SECRETS ---
```mermaid
sequenceDiagram

    autonumber
    App->>Agent: Request Redis password
    Agent->>Vault: Fetch Redis creds (dynamic or KV)
    Vault-->>Agent: Return password/creds
    Agent-->>App: App connects to Redis with fresh creds
    Note over App,Redis: Similar to DB static/dynamic pattern
```

    --- CASE 4: CERTIFICATES & KEYS ---
```mermaid
sequenceDiagram

    autonumber
    App->>Agent: Request TLS cert or key
    Agent->>Vault: Request new cert/key pair
    Vault->>PKI: Generate new X.509 cert + private key (short-lived)
    PKI-->>Vault: Cert and key
    Vault-->>Agent: Return cert and key
    Agent-->>App: App uses cert/key for TLS
    Note over App,PKI: For symmetric/asymmetric keys:<br/>App sends encrypt/decrypt/sign requests to Vault Transit engine<br/>Vault never releases raw master key
```