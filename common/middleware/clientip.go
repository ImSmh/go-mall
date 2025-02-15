package middleware

import (
	"context"
	"jijizhazha1024/go-mall/common/consts/biz"
	"jijizhazha1024/go-mall/common/consts/code"
	"jijizhazha1024/go-mall/common/response"
	"net"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WithClientMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置客户端ip，到ctx
		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			// 处理错误，例如记录日志或使用默认值
			clientIP = r.RemoteAddr // 或者使用默认的 IP 地址
		}
		if clientIP == "" {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.IllegalProxyAddress, code.IllegalProxyAddressMsg))
			return
		}
		ctx := context.WithValue(r.Context(), biz.ClientIPKey, clientIP)
		*r = *r.WithContext(ctx)
		next(w, r)
	}
}
