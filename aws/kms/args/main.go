package args

import (
	"errors"
	"flag"
	. "github.com/nmccready/aws-play/aws"
)

var debug = Spawn("args")

type Args struct {
	Encoding   string
	KeyId      string
	ForceKeyId bool
}

const encodingDesc = `incoming or outgoing additional encoding / decoding
	IE:
		- encrypt -> will take the decimal output and encode in base64
		- decrypt -> will take base64 input and decode base64 to decimal`

const keyIdDesc = "aws kms id or alias defaults to proccess.env.KMS_ID"

const forceKeyDesc = "for decrypt which defaults to false IE uses first key that works, this is to foce a specific key usage"

func GetArgs() *Args {
	args := Args{}

	flag.StringVar(&args.Encoding, "e", "", encodingDesc)
	flag.StringVar(&args.Encoding, "encoding", args.Encoding, encodingDesc)

	if args.Encoding != "" && args.Encoding != "hex" && args.Encoding != "base64" {
		panic(errors.New("encoding choices: hex | base64"))
	}

	flag.StringVar(&args.KeyId, "k", "", keyIdDesc)
	flag.StringVar(&args.KeyId, "key-id", args.KeyId, keyIdDesc)

	flag.BoolVar(&args.ForceKeyId, "fk", false, forceKeyDesc)
	flag.BoolVar(&args.ForceKeyId, "forceKeyId", args.ForceKeyId, forceKeyDesc)

	flag.Parse()

	debug.Log("args: %+v", args)

	return &args
}
