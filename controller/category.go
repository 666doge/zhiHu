package controller

import (
	"zhiHu/db"
	"zhiHu/util"

	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	categoryList, err := db.GetCategoryList()

	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, categoryList)
	return
}