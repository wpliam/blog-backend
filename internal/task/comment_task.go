package task

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/util"
	"context"
	"github.com/wpliap/common-wrap/log"
)

func NewCommentCountTask(proxyService service.ProxyService) *CommentCountTask {
	return &CommentCountTask{
		proxyService: proxyService,
	}
}

type CommentCountTask struct {
	proxyService service.ProxyService
}

func (c *CommentCountTask) Invoke() {
	ctx := context.Background()
	redisCli := c.proxyService.GetRedisProxy()
	dbCli := c.proxyService.GetGormProxy()
	cursor := uint64(0)
	for {
		keys, index, err := redisCli.HScan(ctx, constant.CommentLikeCountKey, cursor, "")
		if err != nil {
			log.Errorf("CommentCountTask HScan err:%v", err)
			return
		}
		for i := 0; i < len(keys); i += 2 {
			key := keys[i]
			val := keys[i+1]
			if err = dbCli.UpdateCommentInfo(util.ParseInt64(key),
				map[string]interface{}{"like_count": val}); err != nil {
				log.Errorf("CommentCountTask UpdateCommentInfo err:%v", err)
				continue
			}
			if err = redisCli.HDel(ctx, constant.CommentLikeCountKey, key); err != nil {
				log.Errorf("CommentCountTask HDel err:%v", err)
			}
		}
		if index == 0 {
			break
		}
		cursor = index
	}
}
