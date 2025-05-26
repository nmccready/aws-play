package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/nmccready/aws-play/aws/kms/args"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kms/kms_iface"
)

func EncryptWithClient(client kms_iface.IClient, text string, args *args.Args) (string, error) {
	keyId := os.Getenv("KMS_ID")
	if args != nil && args.KeyId != "" {
		keyId = args.KeyId
	}
	input := &kms.EncryptInput{
		KeyId:     &keyId,
		Plaintext: []byte(text),
	}
	out, err := client.Encrypt(context.Background(), input)
	if err != nil {
		return "", err
	}
	if args == nil || args.Encoding == "" {
		return string(out.CiphertextBlob), nil
	}
	if args.Encoding == "base64" {
		return base64.StdEncoding.EncodeToString(out.CiphertextBlob), nil
	}
	return hex.EncodeToString(out.CiphertextBlob), nil
}

func Encrypt(text string, args *args.Args) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}
	client := kms.NewFromConfig(cfg)
	return EncryptWithClient(client, text, args)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	args := args.GetArgs()
	out, err := Encrypt(text, args)
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}
