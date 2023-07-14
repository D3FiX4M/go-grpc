package main

import (
	"github.com/D3FiX4M/go-grpc/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Name:  "Max",
		Id:    1,
		Email: "test@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "123",
				Type:   pb.Person_MOBILE,
			},
		},
	}

	body, _ := proto.Marshal(p)

	p1 := &pb.Person{}

	_ = proto.Unmarshal(body, p1)

}
