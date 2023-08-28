package user

type HTTPLoginReq struct {
	Username string `json:"username"` // email or ldap username
	Password string `json:"password"` // password
}

type HTTPLoginRes struct {
	Token JWTToken `json:"token"`
	User  User     `json:"user"`
}

type JWTToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	UserType     uint32 `json:"u_type"`
	LastLoginAt  string `json:"last_login_at"`
	LastLogoutAt string `json:"last_logout_at"`
	UpdateAt     string `json:"update_at"`
}
