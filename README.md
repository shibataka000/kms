# KMS

Encrypt and decrypt file by envelope encryption using AWS KMS.

## Requirement
- OpenSSL

## Usage

### Encrypt
```
$ kms encrypt --help
Encrypt file by envelope encryption using AWS KMS.

Usage:
  kms encrypt [flags]

Flags:
  -h, --help            help for encrypt
      --in string       The path to plaintext file
      --iter uint       Iteration count for PBKDF2. Default is 100000. (default 100000)
      --key-id string   The symmetric encryption KMS key ID that encrypts the data key
      --out string      The path to ciphertext file written into
```

### Decrypt
```
$ kms decrypt --help
Decrypt file by envelope encryption using AWS KMS.

Usage:
  kms decrypt [flags]

Flags:
  -h, --help         help for decrypt
      --in string    The path to ciphertext file
      --out string   The path to plaintext file written into
```

## Install
```
go install github.com/shibataka000/kms@master
```
