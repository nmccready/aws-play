package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	. "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	. "github.com/nmccready/aws-play/aws/kms/args"
)

func Decrypt(text string, args *Args) (string, error) {
	var err error
	var b []byte

	if args != nil && args.Encoding != "" {
		if args.Encoding == "base64" {
			b, err = base64.StdEncoding.DecodeString(text)
			if err != nil {
				fmt.Println("===========")
				fmt.Println(text)
				fmt.Println("===========")
				return "", err
			}
		} else {
			b, err = hex.DecodeString(text)
			if err != nil {
				return "", err
			}
		}
		text = string(b)
		// fmt.Printf("text: %s\n", text)
	}
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

	args := GetArgs()
	out, err := Decrypt(text, args)

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
