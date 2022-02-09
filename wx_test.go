package wework

import (
	"net/http"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	wx, err := New("xxx", "xxx", 44)
	if err != nil {
		t.Error(err)
	}

	t.Log(wx.Token)
	chat, err := wx.GetChat("xxx")
	if err != nil {
		t.Error(err)
	}
	t.Log(chat)
}

func TestStruct(t *testing.T) {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 10 * time.Second,
	}
	wx := Client{
		host:       DefaultHost,
		corpID:     "xxx",
		agentID:    44,
		corpSecret: "xxx",
		handlers:   nil,
		Token: tokenInfo{
			AccessToken: "123123",
		},
		client:   &http.Client{Transport: tr},
		maxRetry: 2,
	}

	wx.Use(TokenHandler, SystemBusyHandler)

	chat, err := wx.GetChat("xxx")
	if err != nil {
		t.Error(err)
	}
	t.Log(chat)
}
