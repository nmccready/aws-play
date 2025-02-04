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

func Decrypt(text string, args *args.Args) (string, error) {
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
	cfg, err := config.LoadDefaultConfig(playContext.GetContext())

	if err != nil {
		return "", err
	}

	svc := kms.NewFromConfig(cfg)

	var keyId *string

	if args.ForceKeyId {
		maybeKey := os.Getenv("KMS_ID")

		if args.KeyId != "" {
			maybeKey = args.KeyId
		}

		if maybeKey != "" {
			keyId = aws.String(maybeKey)
		}
	}

	tCtx, cancel := playContext.GetTimeoutContext()
	defer cancel()
	out, err := svc.Decrypt(tCtx, &kms.DecryptInput{KeyId: keyId, CiphertextBlob: []byte(text)})

	if err != nil {
		return "", err
	}

	return string(out.Plaintext), nil
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
