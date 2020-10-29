package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func Decrypt(text string) (string, error) {
	session, err := NewSession()

	if err != nil {
		return "", err
	}

	svc := kms.New(session)

	out, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: []byte(text)})

	if err != nil {
		return "", err
	}

	return string(out.Plaintext), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	out, err := Decrypt(text)

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
