package main

import server "github.com/fbrd/protobuf-in-go/server"

type list []int

type lists struct {
	list []int
}

var collectionMap = &server.EntityMap{
	Entities: map[int64]*server.Entity{
		1: {
			Id:    1,
			Alias: "first",
			Name:  "Первый",
		},
		2: {
			Id:    2,
			Alias: "second",
			Name:  "Второй",
			IsSef: true,
		},
	},
}

var collectionList = &server.EntityList{
	Entities: []*server.Entity{
		{
			Id:    1,
			Alias: "first",
			Name:  "Первый",
		},
		{
			Id:    2,
			Alias: "second",
			Name:  "Второй",
			IsSef: true,
		},
	},
}

var _ = collectionList

var _ = collectionMap
