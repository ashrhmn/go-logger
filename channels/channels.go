package channels

import "github.com/ashrhmn/go-logger/types"

type AppChannel struct {
	LogStream chan types.AppLog
}

func NewAppChannel() AppChannel {
	return AppChannel{
		LogStream: make(chan types.AppLog),
	}
}
