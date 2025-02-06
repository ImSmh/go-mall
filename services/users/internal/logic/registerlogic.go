package logic

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"math/big"

	"jijizhazha1024/go-mall/common/consts/code"
	"jijizhazha1024/go-mall/dal/model/user"
	"jijizhazha1024/go-mall/services/users/internal/svc"
	"jijizhazha1024/go-mall/services/users/internal/users_biz"
	"jijizhazha1024/go-mall/services/users/users"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var avatarList = []string{
	"http://example.com/avatar1.jpg",
	"http://example.com/avatar2.jpg",
	"http://example.com/avatar3.jpg",
	// 添加更多的头像URL
}

func getRandomAvatar() (string, error) {
	max := big.NewInt(int64(len(avatarList)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return avatarList[n.Int64()], nil
}

// 注册方法
func (l *RegisterLogic) Register(in *users.RegisterRequest) (*users.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	//判断密码是否一致
	if in.Password != in.ConfirmPassword {
		l.Logger.Error("密码不一致")
		return users_biz.HandleRegisterResp("密码不一致", 0, 0, "token")

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Error("密码哈希生成失败", err)
		return users_biz.HandleRegisterResp("密码哈希生成失败", 0, 0, "token")

	}
	email := sql.NullString{
		String: in.Email,
		Valid:  true,
	}
	passwordhash := sql.NullString{
		String: string(hashedPassword),
		Valid:  true,
	}
	//判断邮箱是否已注册，如果已注册，是否处于删除状态
	existUser, err := l.svcCtx.UsersModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			l.Logger.Info("用户不存在", email)
			avatar, err := getRandomAvatar()
			if err != nil {
				l.Logger.Error("获取头像失败", err)
				return users_biz.HandleRegistererror("获取头像失败", 1, nil)
			}

			// 用户不存在，直接注册
			result, insertErr := l.svcCtx.UsersModel.Insert(l.ctx, &user.Users{
				Email:        email,
				PasswordHash: passwordhash,
				AvatarUrl:    sql.NullString{String: avatar, Valid: true},
			})
			l.svcCtx.Bf.Add(in.Email)
			if insertErr != nil {
				l.Logger.Error(code.UserCreationFailedMsg, insertErr)
				return users_biz.HandleRegistererror(code.UserCreationFailedMsg, code.UserCreationFailed, nil)
			}

			userId, lastInsertErr := result.LastInsertId()
			if lastInsertErr != nil {
				l.Logger.Error(code.UserInfoRetrievalFailedMsg, lastInsertErr)
				return users_biz.HandleRegistererror(code.UserInfoRetrievalFailedMsg, code.UserInfoRetrievalFailed, nil)
			}
			return users_biz.HandleRegisterResp(code.CartCreatedMsg, code.CartCreated, uint32(userId), "token")
		}
		l.Logger.Error(code.UserInfoRetrievalFailedMsg, err)
		return users_biz.HandleRegistererror(code.UserInfoRetrievalFailedMsg, code.UserInfoRetrieved, nil)
	}

	if existUser != nil {
		l.Logger.Info(code.UserAlreadyExistsMsg, existUser)
		// 用户已存在，判断是否处于删除状态
		userDeleted := existUser.UserDeleted
		if userDeleted { // 已删除
			// 将删除状态改为false
			updateErr := l.svcCtx.UsersModel.UpdateDeletebyEmail(l.ctx, in.Email, false)
			if updateErr != nil {
				l.Logger.Error(code.UserInfoRetrievalFailedMsg, updateErr)
				return users_biz.HandleRegistererror(code.UserInfoRetrievalFailedMsg, code.UserInfoRetrievalFailed, nil)
			}

			return users_biz.HandleRegisterResp(code.UserCreatedMsg, code.UserCreated, uint32(existUser.UserId), "token")
		} else { // 未删除
			l.Logger.Error(code.UserAlreadyExistsMsg)
			return users_biz.HandleRegistererror(code.UserAlreadyExistsMsg, code.UserAlreadyExists, nil)
		}

	}

	return users_biz.HandleRegisterResp(code.ServerErrorMsg, code.ServerError, 0, "token")
}
