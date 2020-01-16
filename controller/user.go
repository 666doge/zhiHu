package controller

import(
	"github.com/gin-gonic/gin"
	// "zhihu/modal"
)

func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "success",
	})
}