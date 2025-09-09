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

