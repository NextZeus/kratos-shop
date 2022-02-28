package main

import (
	"context"
	"fmt"
	v1 "github.com/nextzeus/kratos-shop/api/user/service/v1"
	"google.golang.org/grpc"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err " + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

func main() {
	TestCreateUser()
	conn.Close()
}

func TestCreateUser() {
	resp, err := userClient.CreateUser(context.Background(), &v1.CreateUserRequest{
		Username: "nextzeus",
		Password: "123456",
	})
	if err != nil {
		panic("grpc create user err " + err.Error())
	}
	fmt.Println(resp)
}
