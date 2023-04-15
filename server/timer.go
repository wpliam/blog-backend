package server

import (
	"blog-backend/internal/task"
	"github.com/robfig/cron/v3"
)

func (s *Server) timerStart() {
	c := cron.New()
	_, _ = c.AddFunc("@every 120s", task.NewArticleCountTask(s.proxyService).Invoke)
	_, _ = c.AddFunc("@every 120s", task.NewCommentCountTask(s.proxyService).Invoke)
	c.Start()
}
