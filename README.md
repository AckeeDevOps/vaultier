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