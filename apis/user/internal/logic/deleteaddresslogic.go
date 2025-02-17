package logic

import (
	"context"

	"jijizhazha1024/go-mall/apis/user/internal/svc"
	"jijizhazha1024/go-mall/apis/user/internal/types"
	"jijizhazha1024/go-mall/common/consts/biz"
	"jijizhazha1024/go-mall/common/consts/code"
	"jijizhazha1024/go-mall/services/users/users"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type DeleteAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAddressLogic {
	return &DeleteAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAddressLogic) DeleteAddress(req *types.DeleteAddressRequest) (resp *types.DeleteAddressResponse, err error) {
	user_id := l.ctx.Value(biz.UserIDKey).(int32)
	DeleteAddResp, err := l.svcCtx.UserRpc.DeleteAddress(l.ctx, &users.DeleteAddressRequest{
		UserId:    user_id,
		AddressId: req.AddressID,
	})

	if err != nil {
		l.Logger.Errorf("调用 rpc 删除地址失败", logx.Field("err", err))
		return nil, errors.New(code.ServerError, code.ServerErrorMsg)
	} else {
		if DeleteAddResp.StatusCode != code.DeleteUserAddressSuccess {
			l.Logger.Errorf("调用 rpc 删除地址失败", logx.Field("status_code", DeleteAddResp.StatusCode), logx.Field("status_msg", DeleteAddResp.StatusMsg))
			return nil, errors.New(int(DeleteAddResp.StatusCode), DeleteAddResp.StatusMsg)
		}
	}

	return
}
