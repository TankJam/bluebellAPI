package controller

type ResCode int64 // 初始化ResCode的int64结构体

// 批量定义常量
const (
	CodeSuccess         ResCode = 1000 + iota // 成功状态码
	CodeInvalidParam                          // 1001 请求参数错误
	CodeUserExist                             // 1002 用户名存在
	CodeUserNotExist                          // 1003 用户不存在
	CodeInvalidPassword                       // 1004 用户名或密码错误
	CodeServerBusy                            // 1005 服务繁忙
	CodeNeedLogin                             // 1006 需要登录
	CodeInvalidToken                          // 1007 token无效
)

// 定义状态码对应返回响应结果的 map
var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "token无效",
}

// Msg 给ResCode结构体定义方法，用于返回状态码对应的消息
func (c ResCode) Msg() string{
	msg, ok := codeMsgMap[c]
	if !ok {  // 若服务出问题，则返回服务繁忙
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

