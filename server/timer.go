package server

import (
	"blog-backend/internal/task"
	"github.com/robfig/cron/v3"
)

func (s *Server) timerStart() {
	c := cron.New()
	c.AddFunc("@every 60s", task.NewArticleCountTask(s.proxyService).Invoke)
	c.Start()
}
