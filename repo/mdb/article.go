package mdb

import (
	"blog-backend/model"
	"gorm.io/gorm"
)

func (cli *MysqlClient) GetAllArticle() ([]*model.Article, error) {
	var articles []*model.Article
	if err := cli.Preload("Category").
		Preload("User").
		Find(&articles, "status = ?", 1).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticleInfo 获取文章信息
func (cli *MysqlClient) GetArticleInfo(articleID int64) (*model.Article, error) {
	var article *model.Article
	if err := cli.
		Where("id = ? and status = ?", articleID, 1).
		First(&article).
		Error; err != nil {
		return nil, err
	}
	return article, nil
}

// GetNextArticle 获取下一篇文章
func (cli *MysqlClient) GetNextArticle(articleID int64) (*model.Article, error) {
	var article *model.Article
	err := cli.Select("id", "title").First(&article, "id > ? and status = ?", articleID, 1).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}

// GetPrevArticle 获取上一篇文章
func (cli *MysqlClient) GetPrevArticle(articleID int64) (*model.Article, error) {
	var article *model.Article
	err := cli.Select("id", "title").Last(&article, "id < ? and status = ?", articleID, 1).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}

// GetCategoryCard 获取分类卡片
func (cli *MysqlClient) GetCategoryCard() ([]*model.CategoryCard, error) {
	var cards []*model.CategoryCard
	if err := cli.Model(&model.Article{}).
		Select("category_id", "sum(view_count) as view_count").
		Scopes(filterStatus()).
		Group("category_id").
		Order("view_count desc").
		Limit(2).
		Find(&cards).
		Error; err != nil {
		return nil, err
	}
	return cards, nil
}

// GetHotArticle 获取热门文章
func (cli *MysqlClient) GetHotArticle() ([]*model.Article, error) {
	var articles []*model.Article
	if err := cli.Preload("User").
		Scopes(filterStatus()).
		Order("view_count desc").
		Limit(5).
		Find(&articles).
		Error; err != nil {
		return nil, err
	}
	return articles, nil
}
