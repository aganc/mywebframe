/**************************************************************
 * Copyright (c) 2021 anxin.com, Inc. All Rights Reserved
 * User: zhangdongsheng<zhangdongsheng@anxin.com>
 * Date: 2021/9/5
 * Desc:
 **************************************************************/

package defines

import "airport/zlog"

const UploadMaxSize = 500 * 1024 * 1024       // 500MB
const BackUpMaxSize = 10 * 1024 * 1024 * 1024 // 10GB

// RequestMethod 接收参数方式
type RequestMethod string

// 以秒为单位的时间常量, 用于定义时间段
// const customKeyExpire = 2 * hour + 30 * minute
const (
	Second int64 = 1
	Minute       = 60 * Second
	Hour         = 60 * Minute
	Day          = 24 * Hour
	Week         = 7 * Day
)

const (
	ReqMethodJson    = "json"
	ReqMethodDefault = "default"
	CtxKeyUserID     = zlog.ContextKeyUserID

	CtxKeyOutsidePlatformCode = "_outsidePlatformCode"
	CtxKeyOpLog               = "oplog"
	CtxKeyLocale              = "locale"

	JWTTokenHeader        = "Authorization"
	JWTTokenRefreshHeader = "new-token"

	// 随机生成的uuid串，用作JWT密钥
	JWTSecretKey = "f9fb6adb-cb42-4b27-a940-e23dc637b23d"
	JWTExpire    = 1 * Hour    // token默认有效期1小时
	JWTRefreshAt = 15 * Minute // 最后15分钟时刷新token
)

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"

const (
	LdapEnable = 1
)

const TokenType = "Bearer"

const LicenseValid = "valid"

const DefaultAvatarPath = "avatar/default.png" // MinIO path of the default user avatar
const DefaultLogoPath = "logo/default.png"     // MinIO path of the default system logo

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

const (
	// 管理员类型

	UTypeAdmin uint32 = 1 + iota // 系统管理员
	UTypeAudit                   // 审计管理员
	UTypeSec                     // 安全管理员
)

var AdminNameKeyMapping = map[uint32]string{
	UTypeAdmin: "admin_system",
	UTypeAudit: "admin_audit",
	UTypeSec:   "admin_secure",
}

const (
	// 用户来源

	USrcGeneral uint32 = 1 + iota // 普通用户
	USrcLDAP                      // LDAP用户
)

const (
	UserEnabled uint32 = 1 + iota
	UserDisabled
)

const (
	UserAttrSimpleLogin uint32 = 1 // 免密登录用户
)

const (
	JWTCheckTypeMultiple int64 = 1 // 可以多端登录
)

// syslog 资产报告选项
const (
	Daily uint32 = 1 + iota
	Weekly
	Manually
)

// cron 任务轮询间隔
const CronPollInterval = 10 * Minute

const DefaultIllaoiPort = 2222
const DefaultHttpsPort = 443
const DefaultHttpPort = 80

const PathPlaceHolder = "/"

const DefaultLanguage = "en-US,en;q=0.9"

const (
	LdapLoginType = "ldap"
)

const (
	MsgYes = "yes"
	MsgNo  = "no"
)

const (
	MsgSuccess = "success"
	MsgFailed  = "failed"
)

const (
	MsgRight = "√"
	MsgError = "×"
)

const (
	MsgOnline  = "online"
	MsgOffline = "offline"
)

const (
	MsgLoaded    = "loaded"
	MsgNotLoaded = "not_loaded"
)

const (
	MsgEnabled  = "enabled"
	MsgDisabled = "disabled"
	MsgStoped   = "stoped"
)

const (
	MSGDaysAgo    = "days_ago"
	MSGHoursAgo   = "hours_ago"
	MSGMinutesAgo = "minutes_ago"
)

const (
	LoginIPHistSearchLimit  int   = 100        // 查询最多多少条最近登录IP, 设为零时不限制
	LoginIPHistSearchDuring int64 = 0 * Minute // 查询最近多长时间内的历史登录IP, 设为零时不限制
)

const HttpsPre = "https://"
const HttpPre = "http://"
