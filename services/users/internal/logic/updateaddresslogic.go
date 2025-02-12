package logic

import (
	"context"
	"database/sql"

	"jijizhazha1024/go-mall/common/consts/code"
	"jijizhazha1024/go-mall/dal/model/user_address"
	"jijizhazha1024/go-mall/services/users/internal/svc"
	"jijizhazha1024/go-mall/services/users/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户地址
func (l *UpdateAddressLogic) UpdateAddress(in *users.UpdateAddressRequest) (*users.UpdateAddressResponse, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.AddressModel.Update(l.ctx, &user_address.UserAddresses{

		AddressId:     int64(in.AddressId),
		RecipientName: in.RecipientName,
		PhoneNumber: sql.NullString{
			String: string(in.PhoneNumber),
			Valid:  in.PhoneNumber != ""},
		Province: sql.NullString{
			String: string(in.Province),
			Valid:  in.Province != ""},
		City:            in.City,
		DetailedAddress: in.DetailedAddress,
		IsDefault:       in.IsDefault,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return &users.UpdateAddressResponse{
				StatusMsg:  code.UserAddressNotFoundMsg,
				StatusCode: code.UserAddressNotFound,
			}, nil
		}
		return &users.UpdateAddressResponse{
			StatusMsg:  code.ServerErrorMsg,
			StatusCode: code.ServerError,
		}, nil
	}

	data := &users.AddressData{
		AddressId:       int32(in.AddressId),
		RecipientName:   in.RecipientName,
		PhoneNumber:     in.PhoneNumber,
		Province:        in.Province,
		City:            in.City,
		DetailedAddress: in.DetailedAddress,
		IsDefault:       in.IsDefault,
		CreatedAt:       "",
		UpdatedAt:       "",
	}

	return &users.UpdateAddressResponse{
		StatusMsg:  code.UpdateUserAddressSuccessMsg,
		StatusCode: code.UpdateUserAddressSuccess,
		Data:       data,
	}, nil
}
