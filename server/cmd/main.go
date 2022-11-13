package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"os"
)

var (
	logger = log.NewHelper(log.With(log.NewStdLogger(os.Stdout), "caller", log.DefaultCaller))
	conf Conf
)

func init() {
	var flagConf string
	flag.StringVar(&flagConf, "conf", "", "配置文件路径, eg: -conf ../conf")
	flag.Parse()
	if flagConf == "" {
		flag.Usage()
		return
	}
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
			),
		)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(any(err))
	}

	if err := c.Scan(&conf); err != nil {
		panic(any(err))
	}

}

func AgentHeartHandler(c *gin.Context) {
	var postData map[string]interface{}
	err := c.ShouldBind(&postData)
	if err != nil {
		logger.Error(err)
	}else{
		logger.Info(postData)
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func main() {
	logger.Info("start")
	router := gin.Default()
	router.POST("/agent/heart", AgentHeartHandler)
	if err := router.Run(fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)); err != nil {
		panic(any(err))
	}
	logger.Info("end")
}