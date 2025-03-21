// Code generated by goctl. DO NOT EDIT.
// Source: audit.proto

package server

import (
	"context"

	"jijizhazha1024/go-mall/services/audit/audit"
	"jijizhazha1024/go-mall/services/audit/internal/logic"
	"jijizhazha1024/go-mall/services/audit/internal/svc"
)

type AuditServer struct {
	svcCtx *svc.ServiceContext
	audit.UnimplementedAuditServer
}

func NewAuditServer(svcCtx *svc.ServiceContext) *AuditServer {
	return &AuditServer{
		svcCtx: svcCtx,
	}
}

// CreateAuditLog 创建审计日志
func (s *AuditServer) CreateAuditLog(ctx context.Context, in *audit.CreateAuditLogReq) (*audit.CreateAuditLogRes, error) {
	l := logic.NewCreateAuditLogLogic(ctx, s.svcCtx)
	return l.CreateAuditLog(in)
}
