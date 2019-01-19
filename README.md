# Vaultier: CI/CD Vault helper

Vaultier will be simple helper for Container native CI/CD pipelines. 
The main goal is to parse specs file, obtain secrets from the Vault 
instance and create basic (one level deep) JSON file for the further 
processing by dotenv-like libraries or Helm.

## Supported output formats

Output formats are controlled via `PLUGIN_RUN_CAUSE` environment variable 
in following manner:

### Delivery

`delivery` produces key:value pairs nested to `secrets` property:

```json
{
  "secrets": {
    "VAR1": "VALUE1",
    "VAR2": "VALUE2"
  }
}
```

Then, you can just refer these values in Helm templates

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  var1: {{ .Values.secrets.VAR1 | b64enc }}
```

and use them with Helm CLI

```bash
helm install -n release01 -f /tmp/secrets.json /path/to/chart
```

### Test

`test` produces top-level key:value structure which is meant to be 
used with tools like [env2](https://www.npmjs.com/package/env2) or 
[dotenv-json](https://www.npmjs.com/package/dotenv-json).

```json
{
  "VAR1": "VALUE1",
  "VAR2": "VALUE2"
}
```

## Configuration options

**`PLUGIN_VAULT_ADDR`**

url of your Vault instance. Typically it's something 
like `http://yourvault.yourdomain.co.uk`

**`PLUGIN_VAULT_TOKEN`** 

vault token, follow 
[official Vault documentation](https://www.vaultproject.io/docs/commands/token/create.html) 
to get one

**`PLUGIN_BRANCH`**

**`PLUGIN_RUN_CAUSE`**

**`PLUGIN_SECRET_SPECS_PATH`**

**`PLUGIN_SECRET_OUTPUT_PATH`**

##Secrets specification file
```yaml
---
# will be used only when 'delivery' is selected
branches:
  # secrets will be selected based on PLUGIN_BRANCH env. variable
  - name: master
    secrets:
      - path: secret/data/blah/production/config1
        # each Vault path can have more secrets you
        # waint to obtain so so can specify more
        # keyMap entries
        keyMap:
          - vaultKey: vaultVariableName1
            localKey: VAR1
      - path: secret/data/blah/production/config2
        keyMap:
          - vaultKey: vaultVariableName2
            localKey: VAR2

# will be used only when 'test' is selected
testConfig:
  secrets:
    - path: secret/data/blah/test/config1
      keyMap:
        - vaultKey: vaultVariableName1
          localKey: VAR1
    - path: secret/data/blah/test/config2
      keyMap:
        - vaultKey: vaultVariableName2
          localKey: VAR2
```