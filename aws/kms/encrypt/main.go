package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	. "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func Encrypt(text string) (string, error) {
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

	return string(out.CiphertextBlob), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	out, err := Encrypt(text)

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
