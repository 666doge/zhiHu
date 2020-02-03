package controller

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"zhiHu/model"
	"zhiHu/util"
	"zhiHu/id_gen"
	"zhiHu/db"
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
	if err == db.ErrCodeUserExist{
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