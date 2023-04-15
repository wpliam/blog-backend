package task

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/model"
	"blog-backend/util"
	"blog-backend/util/thread"
	"context"
	"fmt"
	"github.com/wpliap/common-wrap/log"
	"sync"
)

// NewArticleCountTask ...
func NewArticleCountTask(proxyService service.ProxyService) *ArticleCountTask {
	return &ArticleCountTask{
		rw:           sync.RWMutex{},
		article:      make(map[int64]*model.ArticleContentInfo),
		proxyService: proxyService,
	}
}

type ArticleCountTask struct {
	rw           sync.RWMutex
	article      map[int64]*model.ArticleContentInfo
	proxyService service.ProxyService
}

func (a *ArticleCountTask) addView(ctx context.Context) {
	redisCli := a.proxyService.GetRedisProxy()
	cursor := uint64(0)
	for {
		keys, index, err := redisCli.ZScan(ctx, constant.ArticleViewCountKey, cursor, "")
		if err != nil {
			log.Errorf("AddView ZScan err:%v", err)
			return
		}
		for i := 0; i < len(keys); i += 2 {
			key := keys[i]
			val := keys[i+1]
			a.add(util.ParseInt64(key)).ViewCount = util.ParseInt64(val)
		}
		if index == 0 {
			break
		}
		cursor = index
	}
}

func (a *ArticleCountTask) addLike(ctx context.Context) {
	redisCli := a.proxyService.GetRedisProxy()
	cursor := uint64(0)
	for {
		keys, index, err := redisCli.ZScan(ctx, constant.ArticleLikeCountKey, cursor, "")
		if err != nil {
			log.Errorf("AddLike ZScan err:%v", err)
			return
		}
		for i := 0; i < len(keys); i += 2 {
			key := keys[i]
			val := keys[i+1]
			a.add(util.ParseInt64(key)).LikeCount = util.ParseInt64(val)
		}
		if index == 0 {
			break
		}
		cursor = index
	}
}

func (a *ArticleCountTask) addCollect(ctx context.Context) {
	redisCli := a.proxyService.GetRedisProxy()
	cursor := uint64(0)
	for {
		keys, index, err := redisCli.HScan(ctx, constant.ArticleCollectCountKey, cursor, "")
		if err != nil {
			log.Errorf("AddCollect HScan err:%v", err)
			return
		}
		for i := 0; i < len(keys); i += 2 {
			key := keys[i]
			val := keys[i+1]
			a.add(util.ParseInt64(key)).CollectCount = util.ParseInt64(val)
		}
		if index == 0 {
			break
		}
		cursor = index
	}
}

func (a *ArticleCountTask) add(id int64) *model.ArticleContentInfo {
	a.rw.Lock()
	defer a.rw.Unlock()
	content, ok := a.article[id]
	if !ok {
		content = &model.ArticleContentInfo{}
		a.article[id] = content
	}
	return content
}

func (a *ArticleCountTask) Invoke() {
	dbCli := a.proxyService.GetGormProxy()
	redisCli := a.proxyService.GetRedisProxy()
	ctx := context.Background()
	_ = thread.GoAndWait(
		func() error {
			a.addView(ctx)
			return nil
		},
		func() error {
			a.addLike(ctx)
			return nil
		},
		func() error {
			a.addCollect(ctx)
			return nil
		},
	)
	for id, article := range a.article {
		updateMap := make(map[string]interface{})
		if article.LikeCount > 0 {
			updateMap["like_count"] = article.LikeCount
		}
		if article.CollectCount > 0 {
			updateMap["collect_count"] = article.CollectCount
		}
		if article.ViewCount > 0 {
			updateMap["view_count"] = article.ViewCount
		}
		if err := dbCli.UpdateArticleCount(id, updateMap); err != nil {
			log.Errorf("UpdateArticleCount err:%v", err)
			continue
		}
		for key := range updateMap {
			var err error
			switch key {
			case "view_count":
				err = redisCli.ZRem(ctx, constant.ArticleViewCountKey, fmt.Sprintf("%d", id))
			case "like_count":
				err = redisCli.ZRem(ctx, constant.ArticleLikeCountKey, fmt.Sprintf("%d", id))
			case "collect_count":
				err = redisCli.HDel(ctx, constant.ArticleCollectCountKey, fmt.Sprintf("%d", id))
			}
			if err != nil {
				log.Errorf("articleTask run redis del err:%v", err)
			}
		}
	}
}
