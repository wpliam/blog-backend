package server

import (
	"blog-backend/internal/service"
	"blog-backend/middleware"
	"blog-backend/repo/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	router *gin.Engine            // 路由
	port   uint16                 // 端口
	proxy  service.ProxyService   // 代理
	middle *middleware.Middleware // 中间件

	DisableServerRouter bool // 禁用服务路由
	MaxShutDownTimeout  int  // Shutdown 的超时时间 ms
}

func NewServer(opts ...Option) *Server {
	if err := config.LoadConfig(); err != nil {
		panic("config load err " + err.Error())
	}
	server := defaultServerOption()
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// Run 服务启动
// https://www.bookstack.cn/read/golang_development_notes/zh-9.10.md
// SIGHUP	1	Term	终端控制进程结束(终端连接断开)
// SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
// SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
// SIGILL	4	Core	非法指令(程序错误、试图执行数据段、栈溢出等)
// SIGABRT	6	Core	调用abort函数触发
// SIGFPE	8	Core	算术运行错误(浮点运算错误、除数为零等)
// SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)
// SIGSEGV	11	Core	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
// SIGPIPE	13	Term	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
// SIGALRM	14	Term	时钟定时信号
// SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
func (s *Server) Run() {
	if !s.DisableServerRouter {
		s.initRouter()
	}
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.router,
	}
	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server ListenAndServe err:%v", err)
		}
	}()
	log.Infof("server start succ process %d", os.Getpid())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-c:
	}
	ctx := context.Background()
	if s.MaxShutDownTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(s.MaxShutDownTimeout)*time.Millisecond)
		defer cancel()
	}
	if err := svr.Shutdown(ctx); err != nil {
		log.Errorf("Server Shutdown err:%v", err)
	}
}
