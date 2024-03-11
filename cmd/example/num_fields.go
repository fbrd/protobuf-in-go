package main

import (
	"fmt"
	server "github.com/fbrd/protobuf-in-go/server"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

func numFields() {
	var codec = encoding.GetCodec(proto.Name)

	nums := &server.Nums{
		Id_1:    100,
		Id_15:   101,
		Id_16:   102,
		Id_2047: 103,
		Id_2048: 104,
	}

	b, err := codec.Marshal(nums)
	noErr(err)
	fmt.Printf("%d %d\n%s\n",
		len(b), b, nums)
}
