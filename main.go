package main

import (
	"GameMockClient/msg"
	"GameMockClient/network"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"GameMockClient/proto/user"
)

func main() {
	var mc = network.NewMockClient("127.0.0.1:3563", 1024)
	mc.SetProcessor(msg.Processor)
	// register response handler
	mc.Register(&user.LoginSuccessResponse{}, handleLoginSuccessResponse_S2C)
	mc.Register(&user.LoginFailedResponse{}, handleLoginFailedResponse_S2C)

	mc.Connect()

	// login request
	mc.WriteMsg(&user.LoginRequest{
		Uid:proto.Int32(10000001),
		U1:   proto.Int32(50000000),
		Token:   proto.String("db6b6843e2f4000000"),
		WebserviceUuid: proto.String("260c74fd-4781-405f-9410-08beaa6264be"),
		ChannelId:   proto.Int32(80000000),
	})

	mc.Run()
}

func handleLoginSuccessResponse_S2C(c interface{}, m interface{}) {
	log.Debug("handleLoginSuccessResponse_S2C")
	//a := c.(*network.MockClient)
	//res := m.(*user.LoginSuccessResponse)
}

func handleLoginFailedResponse_S2C(c interface{}, m interface{}) {
	a := c.(*network.MockClient)
	res := m.(*user.LoginFailedResponse)

	if res.GetErrorcode() !="" {
		// do something
		// a.Close()
		_ = a
	}
}
