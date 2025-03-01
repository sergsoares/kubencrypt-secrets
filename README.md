# kubencrypt-secrets

CLI app for backup secrets from a K8S Cluster and encyrpt using [age - secure encryption tool](https://github.com/FiloSottile/age).

## Features

- One single binary for K8S connection and age encryption.
- Support for env vars and CLI Flags being flexible on how to use.

## Instalation

``` bash
$ curl -LO https://github.com/sergsoares/kubencrypt-secrets/releases/download/v0.1.0/kubencrypt-secrets_Linux_x86_64.zip

$ unzip kubencrypt-secrets_Linux_x86_64.zip 
```

## Parameters

```
$ ./kubencrypt-secrets --help
Usage:
  kubencrypt-secrets [OPTIONS]

Application Options:
  -p, --password=   Password to encrypt zip file [$PASSWORD]
      --kubeconfig= Absolute path to kubeconfig (default: ~/.kube/config) [$KUBECONFIG]

Help Options:
  -h, --help        Show this help message
```

## Usage to backup
```
$ ./kubencrypt-secrets --password test
There are 6 secrets to be saved

$ age -d secrets-20250301140544.zip.age > secrets-20250301140544.zip

$ unzip -l secrets-20250301140544.zip
Archive:  secrets-20250301140544.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     3133  1980-00-00 00:00   secrets-20250301140544/kube-system-chart-values-traefik.json
     2166  1980-00-00 00:00   secrets-20250301140544/kube-system-chart-values-traefik-crd.json
      779  1980-00-00 00:00   secrets-20250301140544/kube-system-k3d-krew-bkp-server-0.node-password.k3s.json
     4460  1980-00-00 00:00   secrets-20250301140544/kube-system-k3s-serving.json
   158327  1980-00-00 00:00   secrets-20250301140544/kube-system-sh.helm.release.v1.traefik-crd.v1.json
   228963  1980-00-00 00:00   secrets-20250301140544/kube-system-sh.helm.release.v1.traefik.v1.json
---------                     -------
   397828                     6 files

```
