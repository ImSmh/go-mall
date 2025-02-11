package logic

import (
	"context"
	"database/sql"
	"errors"

	"jijizhazha1024/go-mall/services/users/internal/svc"
	"jijizhazha1024/go-mall/services/users/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAddressesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAddressesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAddressesLogic {
	return &ListAddressesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有收货地址
func (l *ListAddressesLogic) ListAddresses(in *users.AllAddressLitstRequest) (*users.AddressListResponse, error) {
	// todo: add your logic here and delete this line

	allusers, err := l.svcCtx.AddressModel.FindAllByUserId(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("用户地址不存在")
		}
		return nil, err
	}
	addresslist := make([]*users.AddressData, 0)
	for _, user := range allusers {
		addresslist = append(addresslist, &users.AddressData{
			AddressId:       int32(user.AddressId),
			DetailedAddress: user.DetailedAddress,
			City:            user.City,
			Province:        user.Province.String,
			IsDefault:       user.IsDefault,
			RecipientName:   user.RecipientName,
			PhoneNumber:     user.PhoneNumber.String,
			CreatedAt:       user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &users.AddressListResponse{
		Data: addresslist,
	}, nil
}
