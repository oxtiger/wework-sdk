package wework

// TokenHandler token异常则重新获取token并重试
func TokenHandler(wx *Client, api string, code int, data *[]byte) bool {
	var isRetry bool
	if code == ErrorToken || code == ErrorTokenNotFound || code == ErrorTokenExpired {
		wx.removeToken()
		_ = wx.updateToken()
		isRetry = true
	}
	return isRetry
}

// SystemBusyHandler 企业微信服务器繁忙重试
func SystemBusyHandler(wx *Client, api string, code int, data *[]byte) bool {
	var isRetry bool
	if code == ErrorSystemBusy {
		isRetry = true
	}
	return isRetry
}
