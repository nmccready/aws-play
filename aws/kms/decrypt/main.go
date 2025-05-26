package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/nmccready/aws-play/aws/kms/args"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kms/kms_iface"
)

func DecryptWithClient(client kms_iface.IClient, text string, args *args.Args) (string, error) {
	var err error
	var b []byte

	if args != nil && args.Encoding != "" {
		if args.Encoding == "base64" {
			b, err = base64.StdEncoding.DecodeString(text)
			if err != nil {
				return "", err
			}
		} else {
			b, err = hex.DecodeString(text)
			if err != nil {
				return "", err
			}
		}
		text = string(b)
	}

	var keyId *string
	if args.ForceKeyId {
		maybeKey := os.Getenv("KMS_ID")
		if args.KeyId != "" {
			maybeKey = args.KeyId
		}
		if maybeKey != "" {
			keyId = &maybeKey
		}
	}

	input := &kms.DecryptInput{
		KeyId:          keyId,
		CiphertextBlob: []byte(text),
	}
	out, err := client.Decrypt(context.Background(), input)
	if err != nil {
		return "", err
	}
	return string(out.Plaintext), nil
}

func Decrypt(text string, args *args.Args) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}
	client := kms.NewFromConfig(cfg)
	return DecryptWithClient(client, text, args)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	args := args.GetArgs()
	out, err := Decrypt(text, args)
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}
