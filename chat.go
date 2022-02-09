package wework

import (
	"fmt"

	json "github.com/json-iterator/go"
)

// CreateChat 创建群聊
func (wx *Client) CreateChat(name, chatId string, users []string) error {
	if name == "" || chatId == "" || len(users) < 1 {
		return fmt.Errorf("params error")
	}
	var cc createChatBody
	cc.ChatId = chatId
	cc.Name = name
	cc.UserList = users
	cc.Owner = users[0]
	req, err := json.Marshal(cc)
	if err != nil {
		return fmt.Errorf("json error: %s", err.Error())
	}
	_, err = wx.Post(ApiCreateChat, req)
	if err != nil {
		if IsWeworkError(err) {
			return fmt.Errorf("wework error: %s", err.Error())
		}
		return fmt.Errorf("http error: %s", err.Error())
	}
	return nil
}

// UpdateChat 更新群聊
func (wx *Client) UpdateChat(name, chatId, owner string, addUsers, delUsers []string) error {
	if chatId == "" || (len(addUsers) < 1 && len(delUsers) < 1) {
		return fmt.Errorf("params error")
	}
	var cu updateChatBody
	cu.ChatId = chatId
	cu.AddUserList = addUsers
	cu.DelUserList = delUsers
	if name != "" {
		cu.Name = name
	}
	if owner != "" {
		cu.Owner = owner
	}
	req, err := json.Marshal(cu)
	if err != nil {
		return fmt.Errorf("json error: %s", err.Error())
	}
	_, err = wx.Post(ApiUpdateChat, req)
	if err != nil {
		if IsWeworkError(err) {
			return fmt.Errorf("wework error: %s", err.Error())
		}
		return fmt.Errorf("http error: %s", err.Error())
	}
	return nil
}

// GetChat 获取群聊信息
func (wx *Client) GetChat(chatId string) (ChatInfo, error) {
	var cr ChatResponse
	res, err := wx.Get(ApiGetChat, "chatid", chatId)
	if err != nil {
		return cr.ChatInfos, err
	}
	err = json.Unmarshal(res, &cr)
	if err != nil {
		if IsWeworkError(err) {
			return cr.ChatInfos, fmt.Errorf("wework error: %s", err.Error())
		}
		return cr.ChatInfos, fmt.Errorf("http error: %s", err.Error())
	}
	return cr.ChatInfos, nil
}

// ChatTextPush 群聊文本消息推送
func (wx *Client) ChatTextPush(ChatId, content string) error {
	var msg ChatTextMsg
	msg.MsgType = "text"
	msg.ChatId = ChatId
	msg.Text.Content = content

	req, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json error: %s", err.Error())
	}
	_, err = wx.Post(ApiChatSend, req)
	if err != nil {
		if IsWeworkError(err) {
			return fmt.Errorf("wework error: %s", err.Error())
		}
		return fmt.Errorf("http error: %s", err.Error())
	}
	return nil
}

// ChatMarkdownPush 群聊Markdown消息推送
func (wx *Client) ChatMarkdownPush(ChatId, content string) error {
	var msg ChatMarkdownMsg
	msg.MsgType = "markdown"
	msg.ChatId = ChatId
	msg.Markdown.Content = content

	req, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json error: %s", err.Error())
	}
	_, err = wx.Post(ApiChatSend, req)
	if err != nil {
		if IsWeworkError(err) {
			return fmt.Errorf("wework error: %s", err.Error())
		}
		return fmt.Errorf("http error: %s", err.Error())
	}
	return nil
}
