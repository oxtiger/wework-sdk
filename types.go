package wework

import "time"

// 其他配置  TODO json做bind检查

// tokenInfo 企业微信应用token
type tokenInfo struct {
	AccessToken string    `json:"access_token"`
	UpdateTime  time.Time `json:"-"`
}

// errResponse 企业微信错误返回结构
type errResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Result 函数错误返回结构
type Result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Error   error  `json:"-"`
}

// HandlerFunc api response 钩子函数
type HandlerFunc func(wx *Client, api string, code int, data *[]byte) bool

// HandlersChan api response 钩子函数数组
type HandlersChan []HandlerFunc

// 用户相关api

// UserDetail 用户详情
type UserDetail struct {
	ErrCode        int    `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
	Userid         string `json:"userid" `
	Name           string `json:"name"`
	Department     []int  `json:"department"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	Mobile         string `json:"mobile"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	Avatar         string `json:"avatar"`
	ThumbAvatar    string `json:"thumb_avatar"`
	Telephone      string `json:"telephone"`
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	OpenUserid     string `json:"open_userid"`
	MainDepartment int    `json:"main_department"`
	ExtAttr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status           int    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int    `json:"status"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			MiniProgram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

// UserListResponse 用户详情列表
type UserListResponse struct {
	ErrCode  int          `json:"errcode"`
	ErrMsg   string       `json:"errmsg"`
	UserList []UserDetail `json:"userlist"`
}

// 群聊相关api

// createChatBody 创建群聊结构体
type createChatBody struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
	ChatId   string   `json:"chatid"`
}

// updateChatBody 更新群聊结构体
type updateChatBody struct {
	ChatId      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}

// ChatInfo 群聊详情
type ChatInfo struct {
	ChatId   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
}

// ChatResponse chat api response
type ChatResponse struct {
	Errcode   int      `json:"errcode"`
	Errmsg    string   `json:"errmsg"`
	ChatInfos ChatInfo `json:"chat_info"`
}

// 消息推送结构

// // 用户推送

// userBaseMsg 用户私信消息推送公共参数
type userBaseMsg struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentID int64  `json:"agentid"`
}

// UserTextMsg 用户私信文本消息参数
type UserTextMsg struct {
	userBaseMsg
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// UserMarkdownMsg 用户私信Markdown消息参数
type UserMarkdownMsg struct {
	userBaseMsg
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type UserPushResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

// // 群推送

// chatBaseMsg 群聊消息推送公共参数
type chatBaseMsg struct {
	ChatId  string `json:"chatid"`
	MsgType string `json:"msgtype"`
}

// ChatTextMsg 群聊文本消息推送参数
type ChatTextMsg struct {
	chatBaseMsg
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// ChatMarkdownMsg 群聊Markdown消息推送参数
type ChatMarkdownMsg struct {
	chatBaseMsg
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

// // 企业微信机器人

// RobotTextMsg 推送文本消息
type RobotTextMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

// RobotMarkdownMsg 推送markdown消息
type RobotMarkdownMsg struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

// RobotImageMsg 推送图片消息
type RobotImageMsg struct {
	MsgType string `json:"msgtype"`
	Image   struct {
		Base64 string `json:"base64"` // 图片base64 不超过2M
		Md5    string `json:"md5"`
	} `json:"image"`
}

// RobotUploadResponse 文件上传
type RobotUploadResponse struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// RobotFileMsg 推送文件
type RobotFileMsg struct {
	MsgType string `json:"msgtype"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}
