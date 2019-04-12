package e

const SUCCESS = 200

const INVALID_PARAMS = 2001

const ERROR_AUTH_CHECK_TOKEN_FAIL = 2002

const ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 2003

const ERROR_AUTH_TOKEN = 2004

const ERROR_AUTH = 2005

func GetMsg(code int) string {
	var result string
	switch code {
	case 200:
		result = "成功"
		break
	case 2001:
		result = " 缺少 token "
		break
	case 2002:
		result = "token 验证失败"
		break
	case 2003:
		result = "token 超时！！"
		break
	case 2004:
		result = "错误的TOKEN"
		break
	case 2005:
		result = "不存在的用户来获取Token"
		break
	}
	return result
}
