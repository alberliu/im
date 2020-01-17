package api

import (
	"fmt"
	"im/pkg/pb"
	"testing"

	"google.golang.org/grpc"
)

func getUserIntClient() pb.UserIntClient {
	conn, err := grpc.Dial("localhost:50300", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewUserIntClient(conn)
}

func TestUserIntServer_Auth(t *testing.T) {
	_, err := getUserIntClient().Auth(getCtx(), &pb.AuthReq{
		UserId:   9,
		DeviceId: 19,
		Token:    "0",
	})
	fmt.Println(err)
}

func TestUserIntServer_GetUsers(t *testing.T) {
	resp, err := getUserIntClient().GetUsers(getCtx(), &pb.GetUsersReq{
		UserIds: []int64{1},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
