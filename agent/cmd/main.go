package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	tasksChannel = make(chan []byte, 1000)
	conf Conf
)

func init() {
	var flagConf string
	flag.StringVar(&flagConf, "conf", "./conf/", "配置文件路径， -conf ./conf")
	// 加载配置文件
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf)),
	)
	defer c.Close()

	err := c.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err = c.Scan(&conf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		logger = log.NewStdLogger(os.Stdout)
		l = log.NewHelper(log.With(logger, "module", "main"))
	)
	ctx, cancel := context.WithCancel(context.Background())

	// 监听退出信号
	go func() {
		l := log.NewHelper(log.With(logger, "module", "sign control"))
		signChan := make(chan os.Signal, 1)
		defer close(signChan)

		signal.Notify(signChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)
		for {
			select {
			case s, _ := <-signChan:
				l.Info("WARNING: Receiving signal %v, exit", s)
				cancel()
			case <- ctx.Done():
				l.Info("exit")
				return
			}
		}

	}()

	// 启动serverClient，向服务端发送心跳，接收任务
	serverClient := NewServerClient(&conf.Server, logger)
	go serverClient.Poll(ctx)

	// 启动worker，处理任务
	worker := NewWorker(&conf.Worker, logger)
	worker.Start(ctx)

	<- ctx.Done()
	l.Info("wait 5s to end")
	time.Sleep(5)
	l.Info("end server")
}

