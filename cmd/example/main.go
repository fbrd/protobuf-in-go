package main

import (
	"fmt"
	"os"

	client "github.com/fbrd/protobuf-in-go/internal/client"
	server "github.com/fbrd/protobuf-in-go/server"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

func main() {
	var codec = encoding.GetCodec(proto.Name)
	clientEntity := &client.Entity{
		Id:           1,
		Alias:        "abc",
		OBSOLETEName: "name",
		IsSef:        true,
	}

	b, err := codec.Marshal(clientEntity)
	noErr(err)
	/*
	 если будете повторять, установите
	 GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
	 не делайте так в настоящих проектах
	*/

	serverEntity := &server.Entity{}
	err = codec.Unmarshal(b, serverEntity)
	noErr(err)

	fmt.Printf("%d %d\n", len(b), b)
	fmt.Println(serverEntity.String())
}

func noErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
