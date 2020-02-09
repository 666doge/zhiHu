package account

import (
	"zhiHu/session"
	"zhiHu/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SessionName = "session"
	UserId = "user_id"
	UserLoginStatus = "user_status"
	CookieSessionId = "session_id"
	CookieMaxAge = 30 * 24 * 3600
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context){
		
		ProcessRequest(c)
		if IsLogin(c) == false {
			util.RespError(c, util.ErrCodeNotLogin)
			c.Abort()
			return
		}
		c.Next()
	}
}


func ProcessRequest(c *gin.Context) {
	var userSession session.Session

	defer func() {
		if userSession == nil {
			userSession, _ = session.CreateSession()
		}
		c.Set(SessionName, userSession)
	}()

	cookie, err := c.Request.Cookie(CookieSessionId)
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

func SetUserId(userId int64, ctx *gin.Context) {
	var userSession session.Session
	tempSession, exists := ctx.Get(SessionName)
	if !exists {
		return
	}

	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}

	if userSession == nil {
		return
	}

	userSession.Set(UserId, userId)
}

func ProcessResponse(c *gin.Context) {
	var userSession session.Session
	tempSession, exists := c.Get(SessionName)
	if !exists {
		return
	}

	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}

	if userSession == nil {
		return
	}

	if userSession.IsModify() == false {
		return
	}

	err := userSession.Save()
	if err != nil {
		return
	}

	sessionId := userSession.GetId()
	cookie := &http.Cookie{
		Name: CookieSessionId,
		Value: sessionId,
		HttpOnly: true,
		Path: "/",
		MaxAge: CookieMaxAge,
	}

	http.SetCookie(c.Writer, cookie)
	return	
}
