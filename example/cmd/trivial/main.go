package main

import (
	"github.com/rmrobinson-textnow/scratch/proto"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Debug("test")

	msg := &proto.OuterMessage{
		ImportantString: "asdf",
		Inner: &proto.InnerMessage{
			SomeInteger: 1,
		},
	}
	err := msg.Validate()
	if err != nil {
		logger.Warn("error validating proto",
			zap.Error(err),
		)
	} else {
		logger.Debug("cool")
	}
}
