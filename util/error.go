package util

const (
	ErrCodeSuccess = 0
	ErrCodeParameter = 1001
	ErrCodeUserExist = 1002
	ErrCodeServerBusy = 1003
	ErrCodeUserNotExist = 1004
	ErrCodeUserPasswordWrong = 1005
	ErrCodeNotLogin = 1006
	ErrCodeNoRecord = 1007
	ErrCodeRecordExists = 1008
)

func GetMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParameter:
		message = "参数错误"
	case ErrCodeUserExist:
		message = "用户已存在"
	case ErrCodeServerBusy:
		message = "服务器繁忙"
	case ErrCodeUserNotExist:
		message = "用户不存在"
	case ErrCodeUserPasswordWrong:
		message = "密码错误"
	case ErrCodeNotLogin:
		message = "用户未登录"
	case ErrCodeNoRecord:
		message = "查无记录"
	case ErrCodeRecordExists:
		message = "记录已存在"
	default:
		message = "未知错误"
	}
	return
}