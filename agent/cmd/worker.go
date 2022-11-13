package main

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
)

/*
worker 接收任务并处理
 */

type Worker struct {
	logger *log.Helper
}

func NewWorker(conf *WorkerConf, logger log.Logger) *Worker{
	return &Worker{
		logger: log.NewHelper(log.With(logger, "module", "Worker")),
	}
}

func (w *Worker) Start(ctx context.Context) {
	w.logger.Info("start")
	for {
		select {
		case task := <- tasksChannel:
			go w.Do(task)
		case <- ctx.Done():
			w.logger.Info("exit")
			return
		}
	}
}

func (w *Worker) Do(task []byte) {
	var resp map[string]string
	w.logger.Infof("receive task: %s", task)
	err := json.Unmarshal(task, &resp)
	if err != nil {
		w.logger.Error(err)
	}else{
		w.logger.Info(resp)
	}
}
