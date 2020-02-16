package db

import(
	"zhiHu/model"
)

func GetCategoryList () (cateList []*model.Category, err error){
	sqlStr := `select category_id, category_name from category`
	err = DB.Select(&cateList, sqlStr)
	return
}

func GetCategoryName(cateId int64) (cateName string, err error) {
	sqlStr := `select category_name from category where category_id = ?`
	err = DB.Get(&cateName, sqlStr, cateId)
	return
}