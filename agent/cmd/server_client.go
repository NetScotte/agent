package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ServerClient struct {
	Schema string
	Host string
	Port int
	Interval int
	logger *log.Helper
	client *http.Client
}

func NewServerClient(conf *ServerConf, logger log.Logger) *ServerClient {
	return &ServerClient{
		Schema: conf.Schema,
		Host: conf.Host,
		Port: conf.Port,
		Interval: conf.Interval,
		logger: log.NewHelper(log.With(logger, "module", "serverClient")),
		client: http.DefaultClient,
	}

}


func (s *ServerClient) Poll(ctx context.Context) {
	var (
		duration = time.Duration(10) * time.Second
	)
	if s.Interval > 10 {
		duration = time.Duration(s.Interval) * time.Second
	}
	for {
		select {
		case <- time.Tick(duration):
			go s.SendHeartAndGetTask()
		case <- ctx.Done():
			s.logger.Infof("exit")
			return
		}
	}
}

func (s *ServerClient) PostToSever(data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%v://%v:%v/agent/heart", s.Schema, s.Host, s.Port),
		bytes.NewReader(data),
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// SendHeartAndGetTask 发送心跳并接收服务端的返回
func (s *ServerClient) SendHeartAndGetTask() {
	systemInfo := GetSystemInfo()
	s.logger.Info(systemInfo)
	data, err  := json.Marshal(systemInfo)
	if err != nil {
		s.logger.Error(err)
		return
	}
	resp, err := s.PostToSever(data)
	if err != nil {
		s.logger.Error(err)
		return
	}
	tasksChannel <- resp
}