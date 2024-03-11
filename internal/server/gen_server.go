package bar

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./entity.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./enum.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./oneof.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./collection.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./num_fields.proto
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./well_known.proto
