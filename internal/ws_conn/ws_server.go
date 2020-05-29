package ws_conn

import (
	"context"
	"encoding/json"
	"im/config"
	"im/pkg/gerrors"
	"im/pkg/grpclib"
	"im/pkg/logger"
	"im/pkg/pb"
	"im/pkg/rpc_cli"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseInt(r.Header.Get(grpclib.CtxUserId), 10, 64)
	deviceId, _ := strconv.ParseInt(r.Header.Get(grpclib.CtxDeviceId), 10, 64)
	token := r.Header.Get(grpclib.CtxToken)
	requestId, _ := strconv.ParseInt(r.Header.Get(grpclib.CtxRequestId), 10, 64)

	if userId == 0 || deviceId == 0 || token == "" {
		s, _ := status.FromError(gerrors.ErrUnauthorized)
		bytes, err := json.Marshal(s.Proto())
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		w.Write(bytes)
		return
	}

	_, err := rpc_cli.LogicIntClient.ConnSignIn(grpclib.ContextWithRequstId(context.TODO(), requestId), &pb.ConnSignInReq{
		UserId:   userId,
		DeviceId: deviceId,
		Token:    token,
		ConnAddr: config.WSConn.LocalAddr,
	})

	s, _ := status.FromError(err)
	if s.Code() != codes.OK {
		bytes, err := json.Marshal(s.Proto())
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		w.Write(bytes)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	// 断开这个设备之前的连接
	preCtx := load(deviceId)
	if preCtx != nil {
		preCtx.DeviceId = PreConn
	}

	ctx := NewWSConnContext(conn, userId, deviceId)
	store(deviceId, ctx)
	ctx.DoConn()
}

func StartWSServer(address string) {
	http.HandleFunc("/ws", wsHandler)
	logger.Logger.Info("websocket server start")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
