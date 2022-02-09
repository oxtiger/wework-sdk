package wework

// 非get即post
const (
	ApiGetToken         = "/cgi-bin/gettoken"        // 获取token
	ApiGetUserByID      = "/cgi-bin/user/get"        // 获取用户信息
	ApiGetUsersByDeptID = "/cgi-bin/user/list"       // 获取用户列表
	ApiGetDept          = "/cgi-bin/department/list" // 获取部门列表
	ApiUserSend         = "/cgi-bin/message/send"    // 私信推送
	ApiChatSend         = "/cgi-bin/appchat/send"    // 群消息推送
	ApiCreateChat       = "/cgi-bin/appchat/create"  // 创建群聊
	ApiUpdateChat       = "/cgi-bin/appchat/update"  // 更新群聊
	ApiGetChat          = "/cgi-bin/appchat/get"     // 获取群聊详情
	ApiVerifyToken      = "/cgi-bin/menu/get"        // 获取应用菜单, 用来验证token有效性
)
