package account

import (
	"fmt"
	"zhiHu/session"

	"github.com/gin-gonic/gin"
)

const (
	SessionName = "session"
	CookieSessionId = "session_id"
	UserId = "user_id"
	UserLoginStatus = "user_status"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context){
		fmt.Println("授权成功")
		c.Next()
	}
}

func processRequest(c *gin.Context) {
	var userSession session.Session

	defer func() {
		if userSession == nil {
			userSession, _ = session.CreateSession()
		}
		c.Set(SessionName, userSession)
	}()

	cookie, err := c.Request.Cookie("CookieSessionId")
	if err != nil {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, int64(0))
		return
	}

	sessionId := cookie.Value
	if len(sessionId) == 0 {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, int64(0))
		return
	}

	userSession, err = session.Get(sessionId)
	if err != nil {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, int64(0))
		return
	}

	tmpUserId, err := userSession.Get(UserId)
	if err != nil {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, int64(0))
		return
	}

	userId, ok := tmpUserId.(int64)
	if !ok || userId == 0 {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, int64(0))
		return
	}

	c.Set(UserId, int64(userId))
	c.Set(UserLoginStatus, int64(1))
	return
}

func IsLogin(c *gin.Context) (isLogin bool) {
	loginStatus, exists := c.Get(UserLoginStatus)
	if exists {
		return
	}
	loginStat, ok := loginStatus.(int64)
	if !ok {
		return
	}
	if loginStat == 0 {
		return
	}
	return true
}
