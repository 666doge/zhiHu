package account

import (
	"zhiHu/session"
	"zhiHu/util"
	"strconv"
	"errors"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func (c *gin.Context) {
		userId, err := GetUserId(c)
		if err != nil || userId <= 0 {
			util.RespError(c, util.ErrCodeNotLogin)
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetUserId(c *gin.Context) (userId int64, err error) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		return
	}
	
	sessionId := cookie.Value
	if len(sessionId) == 0 {
		return
	}
	uid, err := session.Get(sessionId, "user_id")
	uidStr, ok := uid.(string)
	if !ok {
		err = errors.New("get user_id failed")
		return
	}
	userId, err = strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		return
	}
	return
}