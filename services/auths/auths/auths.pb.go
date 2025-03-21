// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: auths.proto

package auths

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ClientIp string `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auths_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_auths_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_auths_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthReq) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type AuthGenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	ClientIp string `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
}

func (x *AuthGenReq) Reset() {
	*x = AuthGenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auths_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthGenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthGenReq) ProtoMessage() {}

func (x *AuthGenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auths_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthGenReq.ProtoReflect.Descriptor instead.
func (*AuthGenReq) Descriptor() ([]byte, []int) {
	return file_auths_proto_rawDescGZIP(), []int{1}
}

func (x *AuthGenReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthGenReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthGenReq) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type AuthRenewalReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // 用户刷新token
	ClientIp     string `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
}

func (x *AuthRenewalReq) Reset() {
	*x = AuthRenewalReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auths_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRenewalReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRenewalReq) ProtoMessage() {}

func (x *AuthRenewalReq) ProtoReflect() protoreflect.Message {
	mi := &file_auths_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRenewalReq.ProtoReflect.Descriptor instead.
func (*AuthRenewalReq) Descriptor() ([]byte, []int) {
	return file_auths_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRenewalReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthRenewalReq) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type AuthsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserId     uint32 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
}

func (x *AuthsRes) Reset() {
	*x = AuthsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auths_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthsRes) ProtoMessage() {}

func (x *AuthsRes) ProtoReflect() protoreflect.Message {
	mi := &file_auths_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthsRes.ProtoReflect.Descriptor instead.
func (*AuthsRes) Descriptor() ([]byte, []int) {
	return file_auths_proto_rawDescGZIP(), []int{3}
}

func (x *AuthsRes) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *AuthsRes) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *AuthsRes) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AuthGenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg    string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	AccessToken  string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`    // 访问令牌
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // 刷新令牌
	ExpiresIn    int64  `protobuf:"varint,5,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`         // 访问令牌有效期（秒）
}

func (x *AuthGenRes) Reset() {
	*x = AuthGenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auths_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthGenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthGenRes) ProtoMessage() {}

func (x *AuthGenRes) ProtoReflect() protoreflect.Message {
	mi := &file_auths_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthGenRes.ProtoReflect.Descriptor instead.
func (*AuthGenRes) Descriptor() ([]byte, []int) {
	return file_auths_proto_rawDescGZIP(), []int{4}
}

func (x *AuthGenRes) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *AuthGenRes) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *AuthGenRes) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthGenRes) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthGenRes) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

type AuthRenewalRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`      // 状态码，0-成功，其他值-失败
	StatusMsg    string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`          // 返回状态描述
	AccessToken  string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`    // 新的访问令牌
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // 新的刷新令牌
	ExpiresIn    int64  `protobuf:"varint,5,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`         // 新的访问令牌有效期（秒）
}

func (x *AuthRenewalRes) Reset() {
	*x = AuthRenewalRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auths_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRenewalRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRenewalRes) ProtoMessage() {}

func (x *AuthRenewalRes) ProtoReflect() protoreflect.Message {
	mi := &file_auths_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRenewalRes.ProtoReflect.Descriptor instead.
func (*AuthRenewalRes) Descriptor() ([]byte, []int) {
	return file_auths_proto_rawDescGZIP(), []int{5}
}

func (x *AuthRenewalRes) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *AuthRenewalRes) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *AuthRenewalRes) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthRenewalRes) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthRenewalRes) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_auths_proto protoreflect.FileDescriptor

var file_auths_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x75, 0x74, 0x68, 0x73, 0x22, 0x3c, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x22, 0x5e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x22, 0x52, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x22, 0x63, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52,
	0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x0a,
	0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x32, 0xad, 0x01, 0x0a, 0x05,
	0x41, 0x75, 0x74, 0x68, 0x73, 0x12, 0x31, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12,
	0x3a, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auths_proto_rawDescOnce sync.Once
	file_auths_proto_rawDescData = file_auths_proto_rawDesc
)

func file_auths_proto_rawDescGZIP() []byte {
	file_auths_proto_rawDescOnce.Do(func() {
		file_auths_proto_rawDescData = protoimpl.X.CompressGZIP(file_auths_proto_rawDescData)
	})
	return file_auths_proto_rawDescData
}

var file_auths_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auths_proto_goTypes = []any{
	(*AuthReq)(nil),        // 0: auths.AuthReq
	(*AuthGenReq)(nil),     // 1: auths.AuthGenReq
	(*AuthRenewalReq)(nil), // 2: auths.AuthRenewalReq
	(*AuthsRes)(nil),       // 3: auths.AuthsRes
	(*AuthGenRes)(nil),     // 4: auths.AuthGenRes
	(*AuthRenewalRes)(nil), // 5: auths.AuthRenewalRes
}
var file_auths_proto_depIdxs = []int32{
	0, // 0: auths.Auths.Authentication:input_type -> auths.AuthReq
	1, // 1: auths.Auths.GenerateToken:input_type -> auths.AuthGenReq
	2, // 2: auths.Auths.RenewToken:input_type -> auths.AuthRenewalReq
	3, // 3: auths.Auths.Authentication:output_type -> auths.AuthsRes
	4, // 4: auths.Auths.GenerateToken:output_type -> auths.AuthGenRes
	5, // 5: auths.Auths.RenewToken:output_type -> auths.AuthRenewalRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auths_proto_init() }
func file_auths_proto_init() {
	if File_auths_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auths_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AuthReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auths_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AuthGenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auths_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AuthRenewalReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auths_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AuthsRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auths_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AuthGenRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auths_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AuthRenewalRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auths_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auths_proto_goTypes,
		DependencyIndexes: file_auths_proto_depIdxs,
		MessageInfos:      file_auths_proto_msgTypes,
	}.Build()
	File_auths_proto = out.File
	file_auths_proto_rawDesc = nil
	file_auths_proto_goTypes = nil
	file_auths_proto_depIdxs = nil
}
