package model

//cookie 和session
const (
	CookieSession = "session_id"
)

//错误码
const (
	SUCCESS_NO         = 0
	ERR_NO_EMAIL       = 1001
	ERR_NO_SESSION     = 1002
	ERR_PWD_ERROR      = 1003
	ERR_SAVE_SESSION   = 1004
	ERR_NO_LOGIN       = 1005
	ERR_LOGIN_LIMIT    = 1006
	ERR_SORT_MENU      = 1007
	ERR_JSON_UNPACK    = 1008
	ERR_SAVE_MENUS     = 1009
	ERR_SAVE_ARTICLE   = 1010
	ERR_DEL_ARTICLE    = 1011
	ERR_UPDATE_ARTICLE = 1012
	ERR_INVALID_PARAM  = 1013
)

//错误描述
const (
	MSG_SUCCESS        = "res-success"
	MSG_NO_EMAIL       = "email not found"
	MSG_LOGIN_PWD_ERR  = "login pwd error"
	MSG_SAVE_SESSION   = "save session failed"
	MSG_NO_LOGIN       = "no login"
	MSG_LOGIN_LIMIT    = "login limit"
	MSG_SORT_MENU      = "sort menu"
	MSG_JSON_UNPACK    = "json unpack failed"
	MSG_SAVE_MENUS     = "save menus failed"
	MSG_SAVE_ARTICLE   = "save article failed"
	MSG_DEL_ARTICLE    = "delete article failed"
	MSG_UPDATE_ARTICLE = "update article failed"
	MSG_INVALID_PARAM  = "invalid param"
)

//模板渲染状态
const (
	RENDER_MSG_SUCCESS = "res-success"
)
