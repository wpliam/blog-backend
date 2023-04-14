package mdb

import (
	"blog-backend/model"
	"gorm.io/gorm"
)

// GetArticleInfo 获取文章信息
func (cli *MysqlClient) GetArticleInfo(articleID int64) (*model.Article, error) {
	var article *model.Article
	if err := cli.cli.
		Preload("Category").
		Preload("User").
		Where("id = ?", articleID).
		First(&article).
		Error; err != nil {
		return nil, err
	}
	return article, nil
}

// GetNextArticle 获取下一篇文章
func (cli *MysqlClient) GetNextArticle(articleID int64) (*model.Article, error) {
	var article *model.Article
	err := cli.cli.Select("id", "title", "user_id").First(&article, "id > ? and status = ?", articleID, 1).Error
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
	err := cli.cli.Select("id", "title", "user_id").Last(&article, "id < ? and status = ?", articleID, 1).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}

// GetArticleContentInfo 获取文章内容信息
func (cli *MysqlClient) GetArticleContentInfo(articleID int64) (*model.ArticleContentInfo, error) {
	var content *model.ArticleContentInfo
	if err := cli.cli.Scopes(filterStatus()).First(&content, "id = ?", articleID).Error; err != nil {
		return nil, err
	}
	return content, nil
}

// GetCategoryCard 获取分类卡片
func (cli *MysqlClient) GetCategoryCard() ([]*model.CategoryCard, error) {
	var cards []*model.CategoryCard
	if err := cli.cli.Model(&model.Article{}).
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
	if err := cli.cli.Preload("User").
		Scopes(filterStatus()).
		Order("view_count desc").
		Limit(5).
		Find(&articles).
		Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// AddArticle 添加文章
func (cli *MysqlClient) AddArticle(article *model.Article) error {
	return cli.cli.Create(&article).Error
}

// UpdateArticleStatus 更新文章状态
func (cli *MysqlClient) UpdateArticleStatus(article *model.Article) error {
	return cli.cli.Select("status").Updates(&article).Error
}
