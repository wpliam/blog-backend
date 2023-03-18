package mdb

import (
	"blog-backend/model"
	"gorm.io/gorm"
)

func (cli *client) GetAllArticle() ([]*model.Article, error) {
	var articles []*model.Article
	if err := cli.Preload("Category").
		Preload("User").
		Find(&articles, "status = ?", 1).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticleInfo 获取文章信息
func (cli *client) GetArticleInfo(articleID int64) (*model.Article, error) {
	var article *model.Article
	if err := cli.
		Preload("Category").
		Preload("User").
		Where("id = ? and status = ?", articleID, 1).
		First(&article).
		Error; err != nil {
		return nil, err
	}
	return article, nil
}

// GetNextArticle 获取下一篇文章
func (cli *client) GetNextArticle(articleID int64) (*model.Article, error) {
	var article *model.Article
	err := cli.First(&article, "id > ? and status = ?", articleID, 1).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}

// GetPrevArticle 获取上一篇文章
func (cli *client) GetPrevArticle(articleID int64) (*model.Article, error) {
	var article *model.Article
	err := cli.Last(&article, "id < ? and status = ?", articleID, 1).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}
