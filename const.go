package wework

// 企业微信错误码
const (
	DefaultHost = "https://qyapi.weixin.qq.com"

	ErrorTokenExpired  = 42001 // 没有token
	ErrorTokenNotFound = 41001 // 没有token
	ErrorToken         = 40014 // token不合法
	ErrorSystemBusy    = -1    // 系统繁忙
)
