package mdb

import "blog-backend/model"

// GetRecommendIDList 获取文章推荐id列表
func (cli *client) GetRecommendIDList(articleID int64) ([]int64, error) {
	var ids []int64
	if err := cli.Model(&model.ArticleRecommend{}).
		Where("article_id = ?", articleID).
		Pluck("recommend_article_id", &ids).
		Error; err != nil {
		return nil, err
	}
	return ids, nil
}

func (cli *client) BatchCreateRecommendArticle(list []*model.ArticleRecommend) error {
	return cli.CreateInBatches(&list, 20).Error
}
