package api

import (
	"context"
	"fmt"
	"im/pkg/logger"
	"im/pkg/pb"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getLogicExtClient() pb.LogicExtClient {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewLogicExtClient(conn)
}

func getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", "2",
		"device_id", "3",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func TestLogicExtServer_RegisterDevice(t *testing.T) {
	resp, err := getLogicExtClient().RegisterDevice(context.TODO(),
		&pb.RegisterDeviceReq{
			Type:          1,
			Brand:         "huawei",
			Model:         "huawei P30",
			SystemVersion: "1.0.0",
			SdkVersion:    "1.0.0",
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_SendMessage(t *testing.T) {
	resp, err := getLogicExtClient().SendMessage(getCtx(),
		&pb.SendMessageReq{
			MessageId:    "11111",
			ReceiverType: pb.ReceiverType_RT_USER,
			ReceiverId:   1,
			ToUserIds:    nil,
			MessageBody: &pb.MessageBody{
				MessageType: pb.MessageType_MT_TEXT,
				MessageContent: &pb.MessageContent{
					Content: &pb.MessageContent_Text{
						Text: &pb.Text{
							Text: "test",
						},
					},
				},
			},
			IsPersist: true,
			SendTime:  0,
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_CreateGroup(t *testing.T) {
	resp, err := getLogicExtClient().CreateGroup(getCtx(),
		&pb.CreateGroupReq{
			Group: &pb.Group{
				GroupId:      10,
				Name:         "10",
				Introduction: "10",
				Type:         1,
				Extra:        "10",
			},
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_UpdateGroup(t *testing.T) {
	resp, err := getLogicExtClient().UpdateGroup(getCtx(),
		&pb.UpdateGroupReq{
			Group: &pb.Group{
				GroupId:      10,
				Name:         "11",
				Introduction: "11",
				Extra:        "11",
			},
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_GetGroup(t *testing.T) {
	resp, err := getLogicExtClient().GetGroup(getCtx(),
		&pb.GetGroupReq{
			GroupId: 10,
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Error(resp)
}

func TestLogicExtServer_GetUserGroups(t *testing.T) {
	resp, err := getLogicExtClient().GetUserGroups(getCtx(), &pb.GetUserGroupsReq{})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	// todo 不能获取用户所在的超大群组
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_AddGroupMember(t *testing.T) {
	resp, err := getLogicExtClient().AddGroupMember(getCtx(),
		&pb.AddGroupMemberReq{
			GroupUser: &pb.GroupUser{
				GroupId: 2,
				UserId:  3,
				Label:   "3",
				Extra:   "3",
			},
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_UpdateGroupMember(t *testing.T) {
	resp, err := getLogicExtClient().UpdateGroupMember(getCtx(),
		&pb.UpdateGroupMemberReq{
			GroupUser: &pb.GroupUser{
				GroupId: 10,
				UserId:  1,
				Label:   "2",
				Extra:   "2",
			},
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_DeleteGroupMember(t *testing.T) {
	resp, err := getLogicExtClient().DeleteGroupMember(getCtx(),
		&pb.DeleteGroupMemberReq{
			GroupId: 10,
			UserId:  1,
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}
