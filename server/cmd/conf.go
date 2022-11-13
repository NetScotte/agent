package main

type Conf struct {
	Server ServerConf
}

type ServerConf struct {
	Host string
	Port int
}