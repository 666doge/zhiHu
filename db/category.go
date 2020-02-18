package db

import(
	"zhiHu/model"
	"github.com/jmoiron/sqlx"
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

func GetCategoryListById(cateIdList []int64) (cateList []*model.Category, err error) {
	sqlStr := `select category_id, category_name from category where category_id in (?)`
	query, args, err := sqlx.In(sqlStr, cateIdList)
	if err != nil {
		return
	}
	err = DB.Select(&cateList, query, args...)
	if err != nil {
		return
	}
	return
}
