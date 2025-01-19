// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5

package types

type DeleteRequest struct {
	UserId int64 `json:"userId"`
}

type DeleteResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type GetInfoRequest struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

type GetInfoResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	UserId  int64  `json:"userId"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	UserId  int64  `json:"userId"`
	Token   string `json:"token"`
}

type LogoutRequest struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

type LogoutResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type RegisterResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	UserId  int64  `json:"userId"`
	Token   string `json:"token"`
}

type UpdateRequest struct {
	UserId   int64  `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateResponse struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	UserId   int64  `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
