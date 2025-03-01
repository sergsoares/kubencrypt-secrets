# kubencrypt-secrets

CLI app for backup secrest from a K8S Cluster and encyrpt using [age - secure encryption tool](https://github.com/FiloSottile/age).

## Features

- One single binary for K8S connection and age encryption.
- Support for env vars and CLI Flags being flexible on how to use.

## Instalation

``` bash
$ curl -LO https://github.com/sergsoares/kubencrypt-secrets/releases/download/v0.1.0/kubencrypt-secrets_Linux_x86_64.zip

$ unzip kubencrypt-secrets_Linux_x86_64.zip 

$ ./kubencrypt-secrets --password test
There are 6 secrets to be saved 
```

## Usage

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