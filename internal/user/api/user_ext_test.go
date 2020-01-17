package api

import (
	"context"
	"fmt"
	"im/pkg/pb"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getUserExtClient() pb.UserExtClient {
	conn, err := grpc.Dial("localhost:50301", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewUserExtClient(conn)
}

func getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", "9",
		"device_id", "19",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func TestUserExtServer_SignIn(t *testing.T) {
	resp, err := getUserExtClient().SignIn(getCtx(), &pb.SignInReq{
		PhoneNumber: "2",
		Code:        "1",
		DeviceId:    3,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_GetUser(t *testing.T) {
	resp, err := getUserExtClient().GetUser(getCtx(), &pb.GetUserReq{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
