# Vaultier: CI/CD Vault helper

Vaultier will be simple helper for Container native CI/CD pipelines. 
The main goal is to parse specs file, obtain secrets from the Vault 
instance and create basic (one level deep) JSON file for the further 
processing by backend applications or Helm.

> Please note that using of this tool can be potentially dangerous - 
> in the certain phase of CI/CD process, there are secrets stored as 
> a plain text. Hence always use Vaultier in the environment you trust.

## Supported output formats

Output formats are controlled via `VAULTIER_OUTPUT_FORMAT` environment variable 
in following manner:

### Helm

`helm` produces key:value pairs nested to `secrets` property. The last property is 
always `cfg.json` which contains the whole JSON with all provided keys. This is useful 
when you want to provide the whole configuration file to your application.

```json
{
  "secrets": {
    "VAR1": "aGVzbG9KZVZlc2xvCg==",
    "VAR2": "bmVjb0tsZXNsbwo=",
    "cfg.json": "ewogICJWQVIxIjogImhlc2xvSmVWZXNsbyIsCiAgIlZBUjIiOiAibmVjb0tsZXNsbyIKfQo="
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
{{- range $k, $v := .Values.secrets }}
  {{ $k }}: {{ $v }}
{{- end }}
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: mypod
    image: redis
    volumeMounts:
    - name: app-cfg-volume
      mountPath: "/etc/secrets"
      readOnly: true
  volumes:
  - name: app-cfg-volume
    secret:
      secretName: mysecret
```

and use them with Helm CLI

```bash
helm install -n release01 -f /tmp/secrets.json /path/to/chart
```

### .env (JSON)

`dotenv` produces top-level key:value structure which is meant to be 
used with tools like [env2](https://www.npmjs.com/package/env2) or 
[dotenv-json](https://www.npmjs.com/package/dotenv-json).

```json
{
  "VAR1": "VALUE1",
  "VAR2": "VALUE2"
}
```

## Configuration options

**`VAULTIER_VAULT_ADDR`**

url of your Vault instance. Typically it's something 
like `http://yourvault.yourdomain.co.uk`

**`VAULTIER_VAULT_TOKEN`** 

vault token, follow 
[official Vault documentation](https://www.vaultproject.io/docs/commands/token/create.html) 
to get one

**`VAULTIER_BRANCH`**

A branch you want to retrieve secrets for. Same branch has to be 
specified in the specs file. 

**`VAULTIER_RUN_CAUSE`**

This option influences whether `branches` or `testConfig` will  
be selected. Currently supported options are `delivery` or `test`.

**`VAULTIER_OUTPUT_FORMAT`**
As mentioned, this option influences the output format. 
Currently supported options are `helm` or `dotenv`.

**`VAULTIER_SECRET_SPECS_PATH`**

Path to the specs file. If not set, it defaults to `secrets.yaml`.

**`VAULTIER_SECRET_OUTPUT_PATH`**

Path to the output file. Absolute or relative.

## Secrets specification file

```yaml
---
# will be used only when 'delivery' is selected
branches:
  # secrets will be selected based on VAULTIER_BRANCH env. variable
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

## Example use in CD pipeline

in progress.