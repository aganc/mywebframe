package utils

import (
	"airport/defines"
	"airport/zlog"
	"context"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

var (
	key = []byte(defines.JWTSecretKey) // 转换为[]byte的签名密钥
)

// JWTClaims 载荷信息
type JWTClaims struct {
	TokenID   int64 `json:"tokenid"`
	UserID    int64 `json:"userid"`
	TenantID  int64 `json:"tenantid"`
	IssuedAt  int64 `json:"issued_at"`
	ExpiresAt int64 `json:"expires_at"`
	JType     int64 `json:"jtype"` // 1:免密登录，其他暂时为正常登录
}

// Valid 检查是否有效
func (claims *JWTClaims) Valid() error {
	now := time.Now().Unix()

	// 过期
	if now > claims.ExpiresAt {
		return jwt.ErrInvalidKey
	}

	// 早于签发
	if claims.IssuedAt == 0 || now < claims.IssuedAt {
		return jwt.ErrInvalidKey
	}

	return nil
}

// NeedRefresh 检查载荷是否需要刷新
// 调用前应确认claims有效
func (claims *JWTClaims) NeedRefresh() bool {
	now := time.Now().Unix()
	remaining := claims.ExpiresAt - now

	return remaining < defines.JWTRefreshAt
}

type JWT struct {
	Key    []byte
	Expire int64
}

// NewJWT 创建一个jwt工具对象
func NewJWT() *JWT {
	return &JWT{
		Key:    key,
		Expire: defines.JWTExpire,
	}
}

// NewClaims 创建jwt载荷
// 生成随机tokenID
func (j *JWT) NewClaims(userid int64, jwtCheckType ...int64) *JWTClaims {
	rand.Seed(time.Now().UnixNano())

	var jt int64
	if len(jwtCheckType) == 1 && jwtCheckType[0] > 0 {
		jt = jwtCheckType[0]
	}
	return &JWTClaims{
		UserID:  userid,
		TokenID: rand.Int63(),
		JType:   jt,
	}
}

// GenToken 根据载荷生成签名token串
// 传入claims的[IssuedAt, ExpiresAt]将被更新
func (j *JWT) GenToken(ctx context.Context, claims *JWTClaims) (string, error) {
	now := time.Now().Unix()

	claims.IssuedAt = now
	claims.ExpiresAt = now + j.Expire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.Key)
}

// keyfunc 根据未经验证的token信息选择合适的密钥
func (j *JWT) keyfunc(token *jwt.Token) (interface{}, error) {
	return j.Key, nil
}

// ParseToken 校验token串并返回载荷信息
func (j *JWT) ParseToken(ctx context.Context, tk string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tk, &JWTClaims{}, j.keyfunc)

	// 其他错误记录日志
	if err != nil {
		zlog.Warnf(ctx, "jwt parse error: %s", err)
	}

	// token 有效
	if token != nil && token.Valid {
		return token.Claims.(*JWTClaims), nil
	}

	// 隐藏内部错误
	return nil, jwt.ErrInvalidKey
}
