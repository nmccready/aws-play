package logger

import debug "github.com/nmccready/go-debug"

var RootDebug = debug.Debug("@znemz/aws-play")

var Spawn = RootDebug.Spawn
