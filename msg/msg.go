package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"GameMockClient/proto/user"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(3001, &user.LoginRequest{})
	Processor.Register(3002, &user.LoginSuccessResponse{})
	Processor.Register(3003, &user.LoginFailedResponse{})
}
