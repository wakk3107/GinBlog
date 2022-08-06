package errmsg

const (
	SUCCESS       = 200
	ERROR         = 500
	ERROR_REQUEST = 501
	//code=1000...  用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_EXPIRE   = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007
	ERROR_USER_TYPE      = 1008
	ERROR_USER_NOT_RIGHT = 1009
	//code=2000...  文章模块的错误
	ERROR_ART_NOT_FOUND = 2001
	//code=3000...  分类模块的错调
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_TYPE      = 3002
	ERROR_CATE_NOT_EXIST = 3003
)

var CodeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户名不存在",
	ERROR_TOKEN_EXIST:    "Token不存在",
	ERROR_TOKEN_EXPIRE:   "Token已过期",
	ERROR_TOKEN_WRONG:    "Token错误",
	ERROR_TOKEN_TYPE:     "Token格式错误",
	ERROR_REQUEST:        "Bad Request",
	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_TYPE:      "分类格式错误",
	ERROR_ART_NOT_FOUND:  "文章不存在",
	ERROR_CATE_NOT_EXIST: "分类不存在",
	ERROR_USER_NOT_RIGHT: "该用户无权限",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
