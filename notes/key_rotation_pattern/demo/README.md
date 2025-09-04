# Vault Simulation Demo

## Basic syntax
run the vault as a container
```bash
docker run -d \
  --cap-add=IPC_LOCK \
  --name vault \
  -p 8200:8200 \
  -e 'VAULT_DEV_ROOT_TOKEN_ID=dev-only-token' \
  hashicorp/vault:1.19
```

```bash
docker exec -it vault sh
```
```bash
# inside vault
export VAULT_ADDR='http://127.0.0.1:8200' # set endpoint for accessing it 
export VAULT_TOKEN='dev-only-token'  # set token for login
vault status
vault login # input the vault token
```

```bash
# basic syntax
vault version
vault read
vault write
```


```bash 
# WRITE VALUE
vault kv put secret/foo bar=baz 
# put keyvalue `bar = vaz` to the path `secret/foo`,
# this will update the versioning, and also replace any kv inside the path

# create 2 value, and delete all the previous 
vault kv put secret/foo bar=baz hello=world 
```

```bash
# READ VALUE
vault kv get secret/foo
```

```bash
# DELETE VALUE
vault kv delete secret/foo

# UNDELETE VALUE based on version
vault kv undelete -versions=5 secret/foo
```