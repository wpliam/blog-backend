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
		Preload("Category").
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

// GetArticleContentInfo 获取文章内容信息
func (cli *MysqlClient) GetArticleContentInfo(articleID int64) (*model.ArticleContentInfo, error) {
	var content *model.ArticleContentInfo
	if err := cli.Scopes(filterStatus()).First(&content, "id = ?", articleID).Error; err != nil {
		return nil, err
	}
	return content, nil
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

// GetUserArticleCount 获取用户文章数
func (cli *MysqlClient) GetUserArticleCount(uid int64) (int64, error) {
	var count int64
	if err := cli.
		Model(&model.Article{}).
		Where("user_id = ? and status = ?", uid, 1).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetUserViewCount 获取用户文章总浏览量
func (cli *MysqlClient) GetUserViewCount(uid int64) (int64, error) {
	viewInfo := struct {
		Count int64 `json:"count"`
	}{}
	if err := cli.
		Model(&model.Article{}).
		Select("sum(view_count) as count").
		Where("user_id = ? and status = ?", uid, 1).
		Scan(&viewInfo).
		Error; err != nil {
		return 0, err
	}
	return viewInfo.Count, nil
}
