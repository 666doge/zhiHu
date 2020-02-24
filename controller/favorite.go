package controller

import(
	"zhiHu/model"
	"zhiHu/logger"
	"zhiHu/util"
	"zhiHu/id_gen"
	"zhiHu/db"
	"zhiHu/middlewares/account"
	"github.com/gin-gonic/gin"
)

func CreateFavoriteDir(c *gin.Context) {
	var favoriteDir model.FavoriteDir
	err := c.BindJSON(&favoriteDir)
	if err != nil {
		logger.Error("bind json failed, err: %v", err)
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	cid, _ := id_gen.GetId()
	favoriteDir.DirId = int64(cid)

	userId, err := account.GetUserId(c)
	favoriteDir.UserId = userId
	err = db.CreateFavoriteDir(&favoriteDir)
	if err != nil {
		logger.Error("insert favorite dir failed, err: %v", err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, nil)
	return
}

func CreateFavorite(c *gin.Context) {
	var favorite model.Favorite
	err := c.BindJSON(&favorite)
	if err != nil {
		logger.Error("bind json failed, err: %v", err)
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	if favorite.DirId == 0 {
		logger.Error("favorite dirId is 0, favorite: %v", favorite)
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0{
		logger.Error("get userId failed, userId: %v, err: %v", userId, err)
		util.RespError(c, util.ErrCodeNotLogin)
		return
	}
	favorite.UserId = userId
	err = db.CreateFavorite(&favorite)
	if err != nil {
		logger.Error("insert favorite failed, favorite: %v, err: %v", favorite, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, nil)
	return
}