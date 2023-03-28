package mdb

import "blog-backend/model"

func (cli *MysqlClient) GetCategoryInfo(categoryID int64) (*model.Category, error) {
	var categoryInfo *model.Category
	if err := cli.First(&categoryInfo, categoryID).Error; err != nil {
		return nil, err
	}
	return categoryInfo, nil
}

func (cli *MysqlClient) GetCategoryList() ([]*model.Category, error) {
	var categoryList []*model.Category
	if err := cli.Find(&categoryList, "status =  ?", 1).Error; err != nil {
		return nil, err
	}
	return categoryList, nil
}