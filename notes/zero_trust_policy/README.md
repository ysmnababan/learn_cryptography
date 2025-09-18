# Zero Trust Architecture

##🔑 Core Principles of Zero Trust:

Verify Explicitly
Always authenticate and authorize based on all available data points: identity, device health, location, service, workload, etc.

Least Privilege Access
Give users, devices, and applications the minimum access needed, and enforce just-in-time and just-enough-access.

Assume Breach
Design security controls as if an attacker is already inside your network:

Segment networks and services.

Monitor continuously.

Contain threats by limiting lateral movement.

---

## High Level Architecture

```
+------------------------ Cloud / Internet ------------------------+
|                                                                  |
|  Users / Clients  -->  API Gateway  -->  Ingress (mTLS to mesh)  |
|                                 |                                 |
+---------------------------------v---------------------------------+
                                  |
                     +------------+-------------+
                     |   Kubernetes Cluster     |
                     |  (Service Mesh Enabled)  |
                     +------------+-------------+
                                  |
     +-----------------------------+-------------------------------+
     |                             |                               |
+----v----+                  +-----v-----+                   +-----v-----+
| Service |                  |  Service  |                   |  Service  |
|  A      |<--mTLS/ACLs---->|    B      |<--mTLS/ACLs------>|    C      |
| (API)   |                  | (AuthZ)   |                   | (Worker)  |
+----+----+                  +-----+-----+                   +-----+-----+
     |                              |                               |
     | OPA/Envoy ExtAuthZ           | JWT/Policy                    |
     |                               \                              |
     |                                \                             |
     |                                 v                            |
     |                         +--------------+                     |
     |                         |  OPA Agent   |<---- Rego Policies  |
     |                         +--------------+                     |
     |                                                             |
     |                         +--------------+                    |
     +-----------------------> |  HashiCorp   |  <--- PKI/Secrets  |
                               |   Vault      |                    |
                               +-------+------+                    |
                                       |                           |
                                       | Dynamic DB creds / TLS    |
                                       v                           v
                               +--------------+            +---------------+
                               |   Database   |            |  External API |
                               |  (Postgres)  |            |   Providers   |
                               +--------------+            +---------------+
```

**Trust anchors:**

* **Workload identity:** SPIFFE/SPIRE or mesh-issued identities (e.g., Istio Citadel).
* **Network trust:** mTLS everywhere (mesh), deny‑by‑default L4/L7 ACLs.
* **App trust:** JWT/OIDC for end‑user identity, OPA for ABAC/RBAC.
* **Secrets trust:** Vault for short‑lived DB creds & PKI.

---

## 🔐 mTLS vs TLS Decision Matrix

| Communication Pair            | Use Case                                    | Recommended Security Setup                                                                                                       |
| ----------------------------- | ------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| **Service ↔ Service**         | Microservices / API calls inside your infra | ✅ **mTLS** (short-lived certs via cert-manager, Vault, or SPIRE).<br>Strong service identity, encrypted traffic.                 |
| **Frontend ↔ Backend**        | Browser / Mobile App → API                  | ❌ No mTLS.<br>✅ Use **TLS (server cert only)** + **JWT/OAuth2/API keys** for client auth.                                        |
| **Service ↔ Database**        | App → DB (Postgres, MySQL, etc.)            | ⚖️ Either:<br> - ✅ **TLS + password/role** (simpler, common).<br> - ✅ **mTLS** if you have PKI automation (Vault, cert-manager). |
| **Service ↔ Message Broker**  | Kafka, RabbitMQ, NATS, etc.                 | ⚖️ Often support mTLS.<br>✅ Use mTLS if your PKI is automated (e.g., Kafka + Vault).<br>Otherwise TLS + SASL/username.           |
| **Service ↔ Third-party API** | Stripe, AWS, Google, etc.                   | ❌ Almost never mTLS (providers don’t support it).<br>✅ TLS (server cert only) + **API key / OAuth token**.                       |
| **Service ↔ B2B Partner API** | Bank, legacy enterprise API                 | ✅ Sometimes **mTLS required** (they issue you a client cert).<br>Usually static certs + manual renewal.                          |

---

### ✅ Rules of Thumb

* **mTLS is best for internal service-to-service auth** where you control both sides.
* **Frontend clients (users) should never manage certs** → use tokens instead.
* **Databases/message brokers** → mTLS is “defense in depth”, but not strictly required unless compliance needs it.
* **3rd-party APIs** → mTLS only if explicitly required. Otherwise, tokens/keys.

---

⚡ Shortcut:

* **Inside cluster (trust boundary):** mTLS.
* **Outside cluster (users, SaaS APIs):** TLS + tokens.
* **Infra (DB/broker):** TLS + password/role is enough unless you’re in high-security or regulated environments → then mTLS.


