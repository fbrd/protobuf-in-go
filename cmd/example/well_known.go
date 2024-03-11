package main

import (
	server "github.com/fbrd/protobuf-in-go/server"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/durationpb"
)

var personOne = &server.PersonName{
	FirstName: "Иванов",
	LastName:  "Иван",
	SecondName: &wrappers.StringValue{
		Value: "Иванович",
	},
}

var personTwo = &server.PersonName{
	FirstName: "Бьёрн",
	LastName:  "Страуструп",
}

var raceResults = &server.RaceResults{
	Person: personOne,
	Number: 2,
	Result: &durationpb.Duration{
		Seconds: 26,
		Nanos:   400_000_000,
	},
}

var raceDuration = raceResults.Result.AsDuration()

var _, _ = raceDuration, personTwo
