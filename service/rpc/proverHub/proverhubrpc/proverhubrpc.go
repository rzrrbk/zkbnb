// Code generated by goctl. DO NOT EDIT!
// Source: proverHub.proto

package proverhubrpc

import (
	"context"

	"github.com/zecrey-labs/zecrey-legend/service/rpc/proverHub/proverHubProto"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ReqGetUnprovedBlock    = proverHubProto.ReqGetUnprovedBlock
	ReqSubmitProof         = proverHubProto.ReqSubmitProof
	RespGetUnprovedBlock   = proverHubProto.RespGetUnprovedBlock
	RespSubmitProof        = proverHubProto.RespSubmitProof
	ResultGetUnprovedBlock = proverHubProto.ResultGetUnprovedBlock
	ResultSubmitProof      = proverHubProto.ResultSubmitProof

	ProverHubRPC interface {
		GetUnprovedBlock(ctx context.Context, in *ReqGetUnprovedBlock, opts ...grpc.CallOption) (*RespGetUnprovedBlock, error)
		SubmitProof(ctx context.Context, in *ReqSubmitProof, opts ...grpc.CallOption) (*RespSubmitProof, error)
	}

	defaultProverHubRPC struct {
		cli zrpc.Client
	}
)

func NewProverHubRPC(cli zrpc.Client) ProverHubRPC {
	return &defaultProverHubRPC{
		cli: cli,
	}
}

func (m *defaultProverHubRPC) GetUnprovedBlock(ctx context.Context, in *ReqGetUnprovedBlock, opts ...grpc.CallOption) (*RespGetUnprovedBlock, error) {
	client := proverHubProto.NewProverHubRPCClient(m.cli.Conn())
	return client.GetUnprovedBlock(ctx, in, opts...)
}

func (m *defaultProverHubRPC) SubmitProof(ctx context.Context, in *ReqSubmitProof, opts ...grpc.CallOption) (*RespSubmitProof, error) {
	client := proverHubProto.NewProverHubRPCClient(m.cli.Conn())
	return client.SubmitProof(ctx, in, opts...)
}
