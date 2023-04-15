package mdb

import "blog-backend/model"

func (cli *MysqlClient) GetCommentByID(commentID int64) (*model.Comment, error) {
	var comment *model.Comment
	if err := cli.cli.First(&comment, commentID).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (cli *MysqlClient) GetCommentInfo(articleID int64, parentID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	where := make(map[string]interface{})
	where["status"] = 1
	where["article_id"] = articleID
	where["parent_id"] = parentID
	if err := cli.cli.Where(where).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
func (cli *MysqlClient) GetCommentUserIDs(articleID int64) ([]int64, error) {
	var userIDs []int64
	where := make(map[string]interface{})
	where["status"] = 1
	where["article_id"] = articleID
	if err := cli.cli.Model(&model.Comment{}).
		Where(where).Group("user_id").Pluck("user_id", &userIDs).Error; err != nil {
		return nil, err
	}
	return userIDs, nil
}

// GetArticleCommentCount 获取文章评论数
func (cli *MysqlClient) GetArticleCommentCount(articleID int64) (int64, error) {
	var count int64
	where := make(map[string]interface{})
	where["status"] = 1
	where["article_id"] = articleID
	if err := cli.cli.Model(&model.Comment{}).Where(where).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetUserCommentCount 获取用户评论数
func (cli *MysqlClient) GetUserCommentCount(uid int64) (int64, error) {
	var count int64
	where := make(map[string]interface{})
	where["status"] = 1
	where["user_id"] = uid
	if err := cli.cli.Model(&model.Comment{}).Where(where).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// SetCommentInfo 设置评论信息
func (cli *MysqlClient) SetCommentInfo(comment *model.Comment) error {
	return cli.cli.Create(&comment).Error
}

// UpdateCommentInfo 更新评论信息
func (cli *MysqlClient) UpdateCommentInfo(id int64, m map[string]interface{}) error {
	return cli.cli.Model(&model.Comment{}).Where("id = ?", id).Updates(m).Error
}
