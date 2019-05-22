/**
 * @Author: figaro
 * @Description:
 * @File:  msg
 * @Version: 1.0.0
 * @Date: 2019-05-22 13:35
 */

package err

var MsgFalags = map[int]string{

	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_AUTH : "token错误",
	ERROR_AUTH_TOKEN = "token生成失败",
	ErrorAuthCheckTokenTimeout = "token超时失效",
	ErrorAuthCheckTokenFail = "token鉴权失败",

}

func GetMsg(code int) string {
	msg,ok := MsgFalags[code]
	if ok {
		return msg
	}

	return MsgFalags[ERROR]
}