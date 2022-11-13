# 简介
agent一般有两种模式: pull和push
pull: 定时轮询服务端，发送心跳，发送信息，检查是否有任务
push: 服务端主动请求agent，推送任务

pull模式
不够及时

push模式：
agent需要开启端口，以便服务端访问
如果server和agent存在网络问题，不能直连，不好解决


# pull模式
agent有一个goroutine，定时请求server，发送心跳，

