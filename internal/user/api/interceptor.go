package api

import (
	"context"
	"im/internal/user/service"
	"im/pkg/gerrors"
	"im/pkg/grpclib"
	"im/pkg/logger"
	"im/pkg/util"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	IntServerName = "user_int"
	ExtServerName = "user_ext"
)

func logPanic(serverName string, ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, err *error) {
	p := recover()
	if p != nil {
		logger.Logger.Error(serverName+" panic", zap.Any("info", info), zap.Any("ctx", ctx), zap.Any("req", req),
			zap.Any("panic", p), zap.String("stack", util.GetStackInfo()))
		*err = gerrors.ErrUnknown
	}
}

// 服务器端的单向调用的拦截器
func UserIntInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		logPanic(IntServerName, ctx, req, info, &err)
	}()

	resp, err = handler(ctx, req)
	logger.Logger.Debug(IntServerName, zap.Any("info", info), zap.Any("req", req), zap.Any("resp", resp), zap.Error(err))

	s, _ := status.FromError(err)
	if s.Code() != 0 && s.Code() < 1000 {
		md, _ := metadata.FromIncomingContext(ctx)
		logger.Logger.Error(IntServerName, zap.String("method", info.FullMethod), zap.Any("md", md), zap.Any("req", req),
			zap.Any("resp", resp), zap.Error(err), zap.String("stack", gerrors.GetErrorStack(s)))
	}
	return resp, err
}

// 服务器端的单向调用的拦截器
func UserExtInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		logPanic(ExtServerName, ctx, req, info, &err)
	}()

	resp, err = doLogicClientExt(ctx, req, info, handler)
	logger.Logger.Debug(ExtServerName, zap.Any("info", info), zap.Any("ctx", ctx), zap.Any("req", req),
		zap.Any("resp", resp), zap.Error(err))

	s, _ := status.FromError(err)
	if s.Code() != 0 && s.Code() < 1000 {
		md, _ := metadata.FromIncomingContext(ctx)
		logger.Logger.Error(ExtServerName, zap.String("method", info.FullMethod), zap.Any("md", md), zap.Any("req", req),
			zap.Any("resp", resp), zap.Error(err), zap.String("stack", gerrors.GetErrorStack(s)))
	}
	return
}

func doLogicClientExt(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod != "/pb.UserExt/SignIn" {
		userId, deviceId, err := grpclib.GetCtxData(ctx)
		if err != nil {
			return nil, err
		}
		token, err := grpclib.GetCtxToken(ctx)
		if err != nil {
			return nil, err
		}

		err = service.AuthService.Auth(ctx, userId, deviceId, token)
		if err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
