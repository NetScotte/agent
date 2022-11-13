package main


type Conf struct {
	Server ServerConf
	Worker WorkerConf
}

type ServerConf struct {
	Schema string
	Host string
	Port int
	Interval int
	AccessKey string
}

type WorkerConf struct {

}


