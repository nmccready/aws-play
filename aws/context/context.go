package context

import (
	"context"
	"time"
)

var awsPlayContext = context.TODO()
var awsPlayTimeoutSeconds = 3

func GetContext() context.Context {
	return awsPlayContext
}

// allow custom override
func SetContext(ctx context.Context) {
	awsPlayContext = ctx
}

func GetTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(GetContext(), time.Duration(awsPlayTimeoutSeconds)*time.Second)
}

// allow custom override
func SetTimeoutSeconds(seconds int) {
	awsPlayTimeoutSeconds = seconds
}
