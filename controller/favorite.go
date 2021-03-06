package controller

import(
	"zhiHu/model"
	"zhiHu/logger"
	"zhiHu/util"
	"zhiHu/id_gen"
	"zhiHu/db"
	"strconv"
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
	if err == db.ErrRecordExists {
		logger.Error("dir name exsits, dir_name: %v, err: %v", favoriteDir.DirName, err)
		util.RespError(c, util.ErrCodeRecordExists)
		return
	}
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
	if err == db.ErrRecordExists {
		logger.Error("the answer is already in favorite dir, favorite: %#v", favorite)
		util.RespError(c, util.ErrCodeRecordExists)
		return
	}
	if err != nil {
		logger.Error("insert favorite failed, favorite: %v, err: %v", favorite, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, nil)
	return
}

func GetFavoriteDirList(c *gin.Context) {
	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		logger.Error("get user id failed, userId: %v, err: %v", userId, err)
		util.RespError(c, util.ErrCodeNotLogin)
		return
	}
	dirList, err := db.GetFavoriteDirList(userId)
	if err != nil {
		logger.Error("get dir list failed, userId: %v, err: %v", userId, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, dirList)
	return
}

func GetFavoriteList(c *gin.Context) {
	// 获取dirId
	dirIdStr := c.Query("dirId")
	dirId, _ := strconv.ParseInt(dirIdStr, 10, 64)
	
	// 获取userId
	userId, err := account.GetUserId(c)
	if err != nil || userId == 0{
		logger.Error("get user id failed, userId: %v", userId)
		util.RespError(c, util.ErrCodeNotLogin)
		return
	}
	
	// 取 answer id
	answerIdList, err := db.GetFavoriteList(userId, dirId)
	if err != nil {
		logger.Error("get favorite list failed, userId: %v, dirId: %v", userId, dirId)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	// 取 answer list
	answerList, err := db.GetAnswerList(answerIdList)
	if err != nil {
		logger.Error("get answer list failed, answerIds: %v, err: %v", answerIdList, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	util.RespSuccess(c, answerList)
}