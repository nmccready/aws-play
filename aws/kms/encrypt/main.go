package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	playContext "github.com/nmccready/aws-play/aws/context"
	"github.com/nmccready/aws-play/aws/kms/args"
)

func Encrypt(text string, args *args.Args) (string, error) {
	cfg, err := config.LoadDefaultConfig(playContext.GetContext())

	if err != nil {
		return "", err
	}

	svc := kms.NewFromConfig(cfg)

	keyId := os.Getenv("KMS_ID")

	if args.KeyId != "" {
		keyId = args.KeyId
	}

	tCtx, cancel := playContext.GetTimeoutContext()
	defer cancel()
	out, err := svc.Encrypt(tCtx, &kms.EncryptInput{
		KeyId:     aws.String(keyId),
		Plaintext: []byte(text),
	})

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
