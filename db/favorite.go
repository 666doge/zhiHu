package db

import (
	"zhiHu/model"
)

func CreateFavoriteDir(dir *model.FavoriteDir) (err error) {
	var count int64
	sqlStr := `select count(dir_id) from favorite_dir where user_id=? and dir_name=?`
	err = DB.Get(&count, sqlStr, dir.UserId, dir.DirName)

	if count > 0 {
		err = ErrRecordExists
		return
	}

	sqlStr = `insert into favorite_dir (dir_id, dir_name, user_id) values (?, ?, ?)`
	_, err = DB.Exec(sqlStr, dir.DirId, dir.DirName, dir.UserId)
	return
}

func CreateFavorite(favorite *model.Favorite) (err error) {
	var count int64
	sqlStr := `select count(answer_id) from favorite where dir_id = ? and user_id = ?`
	err = DB.Get(&count, sqlStr, favorite.DirId, favorite.UserId)
	if err != nil {
		return
	}
	if count > 0 {
		err = ErrRecordExists
		return
	}

	sqlStr = `insert into favorite (answer_id, dir_id, user_id) values(?,?,?)`
	_, err = DB.Exec(sqlStr, favorite.AnswerId, favorite.DirId, favorite.UserId)
	return
}

func GetFavoriteDirList(userId int64) (dirList []*model.FavoriteDir, err error) {
	sqlStr := `select user_id, dir_id, dir_name from favorite_dir where user_id = ?`
	err = DB.Select(&dirList, sqlStr, userId)
	return
}