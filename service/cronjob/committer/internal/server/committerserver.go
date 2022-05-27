// Code generated by goctl. DO NOT EDIT!
// Source: committer.proto

package server

import (
	"context"
	"github.com/zecrey-labs/zecrey-legend/service/cronjob/committer/committerProto"
	"github.com/zecrey-labs/zecrey-legend/service/cronjob/committer/internal/logic"
	"github.com/zecrey-labs/zecrey-legend/service/cronjob/committer/internal/svc"
)

type CommitterServer struct {
	svcCtx *svc.ServiceContext
	committerProto.UnimplementedCommitterServer
}

func NewCommitterServer(svcCtx *svc.ServiceContext) *CommitterServer {
	return &CommitterServer{
		svcCtx: svcCtx,
	}
}

func (s *CommitterServer) Ping(ctx context.Context, in *committerProto.Request) (*committerProto.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}