package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	. "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	. "github.com/nmccready/aws-play/aws/kms/args"
)

func Encrypt(text string, args *Args) (string, error) {
	session, err := NewSession()

	if err != nil {
		return "", err
	}

	svc := kms.New(session)

	out, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(os.Getenv("KMS_ID")),
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

	args := GetArgs()
	out, err := Encrypt(text, args)

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
