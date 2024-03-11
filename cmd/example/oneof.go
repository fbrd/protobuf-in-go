package main

import server "github.com/fbrd/protobuf-in-go/server"

var metadataString = &server.Metadata{
	Key: "size",
	Value: &server.Metadata_StringValue{
		StringValue: "L",
	},
}

var metadataInt = &server.Metadata{
	Key: "size",
	Value: &server.Metadata_IntValue{
		IntValue: 48,
	},
}

var _, _ = metadataInt, metadataString
