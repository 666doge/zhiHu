package controller

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"zhiHu/model"
	"zhiHu/util"
	"zhiHu/id_gen"
	"zhiHu/db"
	"zhiHu/session"
	"net/http"
	uuid "github.com/satori/go.uuid"
)

func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "success",
	})
}

func UserRegister(c *gin.Context) {
	var userInfo model.User
	err := c.BindJSON(&userInfo)

	if err != nil {
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	if (len(userInfo.Password) == 0 || (len(userInfo.Email) == 0 && len(userInfo.Phone) == 0)) {
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	if (userInfo.Sex != model.UserSexMan && userInfo.Sex != model.UserSexWoman ) {
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	userId, err := id_gen.GetId()
	userInfo.UserId = int64(userId)
	if err != nil {
		fmt.Println("id err:", err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	err = db.Register(&userInfo)
	if err == db.ErrUserExist{
		util.RespError(c, util.ErrCodeUserExist)
		return
	}
	if err != nil {
		fmt.Println("db err:", err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	util.RespSuccess(c, nil)
}

func UserLogin(c *gin.Context) {
	var err error
	var userInfo model.User

	defer func() {
		if err != nil {
			return
		}
		// 登录成功-设置cookie, 保存session
		uuid := uuid.NewV4()
		sessionId := uuid.String()
		err := session.Set(sessionId, "user_id", userInfo.UserId)
		if err != nil {
			return
		}
	
		cookie := &http.Cookie{
			Name: "session_id",
			Value: sessionId,
			HttpOnly: true,
			Path: "/",
			MaxAge: 30 * 24 * 3600,
		}
	
		http.SetCookie(c.Writer, cookie)
		util.RespSuccess(c, nil)
	}()

	err = c.BindJSON(&userInfo)
	if err != nil {
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	if len(userInfo.UserName) == 0 || len(userInfo.Password) == 0 {
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	err = db.UserLogin(&userInfo)
	if err == db.ErrUserNotExist {
		util.RespError(c, util.ErrCodeUserNotExist)
		return
	}

	if err == db.ErrUserPasswordWrong {
		util.RespError(c, util.ErrCodeUserPasswordWrong)
		return
	}

	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
}