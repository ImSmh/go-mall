package rpc

import (
	"context"
	"fmt"
	"jijizhazha1024/go-mall/common/consts/biz"
	"jijizhazha1024/go-mall/services/users/users"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var users_client users.UsersClient

func initusers() {
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", biz.UsersRpcPort), grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	users_client = users.NewUsersClient(conn)
}

func TestUsersRpc(t *testing.T) {
	initusers()
	resp, err := users_client.Register(context.Background(), &users.RegisterRequest{
		Email:           "test2@test.com",
		Password:        "234567",
		ConfirmPassword: "234567",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("register success", resp)
	t.Log("register success", resp)
}
