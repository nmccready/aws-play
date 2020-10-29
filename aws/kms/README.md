# AWS KMS

Example:

```bash
$ echo abcd | aws/kms/encrypt/index.js |  aws/kms/decrypt/index.js
abcd

echo abcd | aws/kms/encrypt/index.js -e base64 | aws/kms/decrypt/index.js -e base64
abcd

$ echo abcd | aws/kms/encrypt/index.js -e base64
AQICAHj/6a1KHdB7qaXDbeWQ9K48M0vQfukO9weGdqwlCJ2ehQE1oq//vPbhrBUdFOe8Gl1LAAAAYzBhBgkqhkiG9w0BBwagVDBSAgEAME0GCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMW8Nhd/r1THSG7Z0xAgEQgCC+lkYTTnyKgRGnsup2EF1pPFqubT7hg/w+KRBkNB7DUA==%

$ echo abcd | aws/kms/encrypt/index.js -e hex
0102020078ffe9ad4a1dd07ba9a5c36de590f4ae3c334bd07ee90ef7078676ac25089d9e850154dda86285de1497204cab5888e4047300000063306106092a864886f70d010706a0543052020100304d06092a864886f70d010701301e060960864801650304012e3011040c54bbc5ef8344f6f6254d51c202011080206b8b615ff5bb007bbe16f28e745a1315adf1a3bc0b94f46570218452e56ff058%
```

```bash
$ echo hi | go run ./aws/kms/encrypt/main.go | go run ./aws/kms/decrypt/main.go
hi
```
