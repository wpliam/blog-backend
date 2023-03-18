package mdb

import "blog-backend/model"

func (cli *client) GetCommentInfo(articleID int64, parentID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	where := make(map[string]interface{})
	where["status"] = 1
	where["article_id"] = articleID
	where["parent_id"] = parentID
	if err := cli.Where(where).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
func (cli *client) GetCommentUserIDs(articleID int64) ([]int64, error) {
	var userIDs []int64
	where := make(map[string]interface{})
	where["status"] = 1
	where["article_id"] = articleID
	if err := cli.Model(&model.Comment{}).
		Where(where).Group("user_id").Pluck("user_id", &userIDs).Error; err != nil {
		return nil, err
	}
	return userIDs, nil
}

// GetArticleCommentCount 获取文章评论数
func (cli *client) GetArticleCommentCount(articleID int64) (int64, error) {
	var count int64
	where := make(map[string]interface{})
	where["status"] = 1
	where["article_id"] = articleID
	if err := cli.Model(&model.Comment{}).Where(where).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// SetCommentInfo 设置评论信息
func (cli *client) SetCommentInfo(comment *model.Comment) error {
	return cli.Create(&comment).Error
}
