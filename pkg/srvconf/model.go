package srvconf

type Configuration interface {
	GetServer() Server
	GetDir() string
}

type Server struct {
	Name string
	Host string
	Port uint16
	Env  string
}

type BaseConfig struct {
	Server Server
}

func (c BaseConfig) GetServer() Server {
	return c.Server
}
