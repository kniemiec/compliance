module mymodules/compliance

go 1.18

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.28.0
)

require (
	golang.org/x/net v0.0.0-20190311183353-d8887717615a // indirect
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a // indirect
	golang.org/x/text v0.3.0 // indirect
)

// replace google.golang.org/grpc => ../