package mdb

import "blog-backend/model"

// GetArticleTagID 获取文章标签id
func (cli *MysqlClient) GetArticleTagID(articleID int64) ([]int64, error) {
	var ids []int64
	if err := cli.
		Model(&model.ArticleTag{}).
		Where("article_id = ?", articleID).
		Pluck("tag_id", &ids).
		Error; err != nil {
		return nil, err
	}
	return ids, nil
}

func (cli *MysqlClient) BatchInsertArticleTagID(articleTags []*model.ArticleTag) error {
	return cli.CreateInBatches(&articleTags, 10).Error
}
