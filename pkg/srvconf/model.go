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

type SQLite struct {
	Directory    string
	DatabaseName string
}

type JWT struct {
	SigningKey string
}

type BaseConfig struct {
	Server Server
	SQLite SQLite
	JWT    JWT
}

func (c BaseConfig) GetServer() Server {
	return c.Server
}
