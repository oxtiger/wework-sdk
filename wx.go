package wework

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	json "github.com/json-iterator/go"
)

// Client 企业微信client
type Client struct {
	host       string       // api host
	corpID     string       // 企业corpID
	agentID    int64        // 应用agentID
	corpSecret string       // 应用corpSecret
	handlers   HandlersChan // api response hook handlers
	Token      tokenInfo    // 应用token
	client     *http.Client // http client
	maxRetry   int          // 重试次数
	lock       sync.Mutex
}

// New TODO 实现一个infof() errorf()标准logger interface
// New 初始化一个企业微信client
func New(corpID, corpSecret string, agentID int64) (*Client, error) {
	wx := &Client{
		host:       DefaultHost,
		corpID:     corpID,
		agentID:    agentID,
		corpSecret: corpSecret,
		maxRetry:   2,
		Token:      tokenInfo{},
	}

	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 10 * time.Second,
	}
	wx.client = &http.Client{Transport: tr}
	err := wx.updateToken()
	wx.Use(TokenHandler, SystemBusyHandler)
	return wx, err
}

// SetHost 设置企业微信api host
func (wx *Client) SetHost(host string) {
	wx.host = host
}

// SetHttpClient 设置自定义的http client以支持更高级的姿势或兼容企业内部框架
func (wx *Client) SetHttpClient(client *http.Client) {
	wx.client = client
}

// SetRetry 设置重试次数, handFunc返回重试的情况下最大可重试次数, 建议三次以内
func (wx *Client) SetRetry(reTry int) *Client {
	wx.maxRetry = reTry
	return wx
}

// joinUri get请求参数拼接
func (wx *Client) joinUri(api string, params ...string) string {
	var uri strings.Builder
	uri.Grow(len(wx.host))
	uri.Grow(len(api))
	uri.Grow(len(wx.Token.AccessToken) + 200)
	uri.WriteString(wx.host)
	uri.WriteString(api)
	uri.WriteString("?access_token=")
	uri.WriteString(wx.Token.AccessToken)
	for n, p := range params {
		if n%2 == 0 {
			uri.WriteString("&")
			uri.WriteString(p)
			uri.WriteString("=")
		} else {
			uri.WriteString(p)
		}
	}
	return uri.String()
}

// parseResult 对所有请求的response预处理
func (wx *Client) parseResult(api string, response, data []byte) (bool, error) {
	errCode := json.Get(response, "errcode").ToInt()
	if errCode != 0 {
		var isRetry bool
		for _, handFunc := range wx.handlers {
			if b := handFunc(wx, api, errCode, &data); b {
				isRetry = true
			}
		}
		errMsg := json.Get(response, "errmsg").ToString()
		err := NewWeworkError("errCode = %d, errmsg = %s", errCode, errMsg)
		return isRetry, &err
	}
	return false, nil
}

// removeToken 移除失效token
func (wx *Client) removeToken() {
	wx.Token = tokenInfo{}
}

// updateToken Token更新
func (wx *Client) updateToken() error {
	// 防止并发获取token的情况出现
	wx.lock.Lock()
	defer wx.lock.Unlock()
	if wx.Token.AccessToken != "" && !wx.Token.UpdateTime.IsZero() {
		if err := wx.verifyToken(); err == nil {
			return nil
		}
	}

	res, err := wx.Get(ApiGetToken, "corpid", wx.corpID, "corpsecret", wx.corpSecret)
	if err != nil {
		return fmt.Errorf("get token failed: %s", err.Error())
	}
	var tk tokenInfo
	err = json.Unmarshal(res, &tk)
	if err != nil {
		return fmt.Errorf("unmarshal tokenResponse failed: %s", err.Error())
	}
	// 更新token设置更新时间
	wx.Token.AccessToken = tk.AccessToken
	wx.Token.UpdateTime = time.Now()
	return nil
}

// updateToken Token更新
func (wx *Client) verifyToken() error {
	resp, err := wx.client.Get(wx.joinUri(ApiVerifyToken, "agentid", strconv.FormatInt(wx.agentID, 10)))
	if err != nil {
		return err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	errCode := json.Get(res, "errcode").ToInt()
	if errCode != 0 {
		return fmt.Errorf("token error, errCode: %d", errCode)
	}
	return nil
}

// LoopUpdateToken 异步循环更新token
func (wx *Client) LoopUpdateToken() {
	for {
		// 默认过期时间是7200秒, 这里一小时更新一次
		time.Sleep(time.Hour)
		err := wx.updateToken()
		if err != nil {
			log.Println("go-wework UpdateToken error: ", err)
		}
	}
}

// Use 增加返回错误码处理中间件
func (wx *Client) Use(handlerFunc ...HandlerFunc) *Client {
	wx.handlers = append(wx.handlers, handlerFunc...)
	return wx
}

// Get get请求企微api
func (wx *Client) Get(api string, params ...string) ([]byte, error) {
	var (
		reTry   int
		isRetry bool
		resp    *http.Response
		res     []byte
		err     error
	)
	for reTry < wx.maxRetry {
		resp, err = wx.client.Get(wx.joinUri(api, params...))
		if err != nil {
			reTry += 1
			continue
		}
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			reTry += 1
			continue
		}
		_ = resp.Body.Close()
		isRetry, err = wx.parseResult(api, res, nil)
		if isRetry && reTry < wx.maxRetry {
			reTry += 1
			continue
		} else {
			break
		}
	}
	return res, err
}

// Post post请求企微api
func (wx *Client) Post(api string, data []byte) ([]byte, error) {
	var (
		reTry   int
		isRetry bool
		resp    *http.Response
		res     []byte
		err     error
	)
	for reTry < wx.maxRetry {
		reTry += 1
		resp, err = wx.client.Post(wx.joinUri(api), "application/json", bytes.NewReader(data))
		if err != nil {
			continue
		}
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		_ = resp.Body.Close()
		isRetry, err = wx.parseResult(api, res, nil)
		if isRetry && reTry < wx.maxRetry {
			continue
		} else {
			break
		}
	}
	return res, err
}
