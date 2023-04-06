package mdb

import (
	"blog-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AddSearchFlow 添加搜索的流水
func (cli *MysqlClient) AddSearchFlow(keyword string) error {
	flow := &model.SearchFlow{
		Keyword: keyword,
		Version: 1,
		Flag:    1,
	}
	if err := cli.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "keyword"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"version": gorm.Expr("version +1"),
			"flag":    1,
		}),
	}).Create(flow).Error; err != nil {
		return err
	}
	return nil
}

func (cli *MysqlClient) GetSearchFlow(keyword string) ([]*model.SearchFlow, error) {
	var flows []*model.SearchFlow
	if err := cli.Scopes(func(db *gorm.DB) *gorm.DB {
		if (keyword) == "" {
			return db
		}
		return db.Where("keyword like ?", "%"+keyword+"%")
	}).Order("version desc").
		Limit(10).
		Find(&flows).
		Error; err != nil {
		return nil, err
	}
	return flows, nil
}
