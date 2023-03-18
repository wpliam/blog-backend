package mdb

import (
	"blog-backend/model"
)

// GetTagList 获取标签
func (cli *client) GetTagList(ids ...int64) ([]*model.Tag, error) {
	var tags []*model.Tag
	where := make(map[string]interface{})
	where["status"] = 1
	if len(ids) > 0 {
		where["id"] = ids
	}
	if err := cli.Where(where).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// FirstOrCreateTag tag_name存在查询记录,不存在则创建记录
func (cli *client) FirstOrCreateTag(tag *model.Tag) error {
	return cli.Where(&model.Tag{TagName: tag.TagName}).FirstOrCreate(&tag).Error
}
