package wework

import (
	"fmt"

	json "github.com/json-iterator/go"
)

// UserByID 通过user_id获取用户详情
func (wx *Client) UserByID(userID string) (UserDetail, error) {
	var userDetail UserDetail
	res, err := wx.Get(ApiGetUserByID, "userid", userID)
	if err != nil {
		if IsWeworkError(err) {
			return userDetail, fmt.Errorf("wework error: %s", err.Error())
		}
		return userDetail, fmt.Errorf("http error: %s", err.Error())
	}
	err = json.Unmarshal(res, &userDetail)
	if err != nil {
		return userDetail, fmt.Errorf("json.Marshal error: %s", err.Error())
	}
	return userDetail, nil
}

// UsersByDeptID 通过部门id获取用户详情
func (wx *Client) UsersByDeptID(departmentID string) ([]UserDetail, error) {
	var userList UserListResponse
	res, err := wx.Get(ApiGetUsersByDeptID, "department_id", departmentID, "fetch_child", "1")
	if err != nil {
		if IsWeworkError(err) {
			return nil, fmt.Errorf("wework error: %s", err.Error())
		}
		return nil, fmt.Errorf("http error: %s", err.Error())
	}
	err = json.Unmarshal(res, &userList)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error: %s", err.Error())
	}
	return userList.UserList, nil
}

// UserTextPush 私信文本消息推送
func (wx *Client) UserTextPush(users, content string) (string, error) {
	var msg UserTextMsg
	msg.MsgType = "text"
	msg.ToUser = users
	msg.AgentID = wx.agentID
	msg.Text.Content = content

	req, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("json.Marshal error: %s", err.Error())
	}
	res, err := wx.Post(ApiUserSend, req)
	if err != nil {
		return "", fmt.Errorf("http error: %s", err.Error())
	}
	var userPushRes UserPushResponse
	_ = json.Unmarshal(res, &userPushRes)
	if len(userPushRes.InvalidUser) > 0 {
		return userPushRes.InvalidUser, fmt.Errorf("wework error: invalid user_id: %v", userPushRes.InvalidUser)
	}
	return "", nil
}

// UserMarkdownPush 私信Markdown消息推送
func (wx *Client) UserMarkdownPush(users, content string) (string, error) {
	var msg UserMarkdownMsg
	msg.MsgType = "markdown"
	msg.ToUser = users
	msg.AgentID = wx.agentID
	msg.Markdown.Content = content

	req, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("json.Marshal error: %s", err.Error())
	}
	res, err := wx.Post(ApiUserSend, req)
	if err != nil {
		if IsWeworkError(err) {
			var userPushRes UserPushResponse
			_ = json.Unmarshal(res, &userPushRes)
			return userPushRes.InvalidUser, fmt.Errorf("wework error: %s", err.Error())
		}
		return "", fmt.Errorf("http error: %s", err.Error())
	}
	return "", nil
}
