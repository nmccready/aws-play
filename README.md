# AWS KMS

## Node

`npm install -g @znemz/aws-play`

Expects KMS_ID to be available / exported

```bash
echo abcd | awsKmsEncrypt -e base64 -k alias/demo | awsKmsDecrypt -e base64
abcd

echo abcd | awsKmsEncrypt -e base64 -k alias/demo
AQICAHj/6a1KHdB7qaXDbeWQ9K48M0vQfukO9weGdqwlCJ2ehQE2GJx31AA8adTIcCOKmJf9AAAAYzBhBgkqhkiG9w0BBwagVDBSAgEAME0GCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMltAgB55xMNmhaLR3AgEQgCAJVDr5cdtjELQwHbtzWbAwwWD/iTVWPpoD9qpnqW2iZg==%
```

## Go

`go get github.com/nmccready/aws-play/aws/kms/encrypt`
`go get github.com/nmccready/aws-play/aws/kms/decrypt`

```bash
$ echo abcd | encrypt -e base64 | decrypt -e base64
abcd
```
