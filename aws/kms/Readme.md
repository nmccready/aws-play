# AWS KMS

Example:

```bash
$ echo abcd | aws/kms/encrypt/index.js |  aws/kms/decrypt/index.js
abcd
```

```bash
$ echo hi | go run ./aws/kms/encrypt/main.go | go run ./aws/kms/decrypt/main.go
hi
```
