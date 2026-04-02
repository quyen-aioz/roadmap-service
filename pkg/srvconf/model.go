package srvconf

type Configuration interface {
	GetServer() Server
	GetDir() string
}

type Server struct {
	Name         string
	Host         string
	Port         uint16
	Env          string
	AllowOrigins []string
}

type SQLite struct {
	Directory    string
	DatabaseName string
}

type JWT struct {
	SigningKey string
}

type SeedAdmin struct {
	Username string
	Password string
}

type BaseConfig struct {
	Server    Server
	SQLite    SQLite
	JWT       JWT
	SeedAdmin SeedAdmin
}

func (c BaseConfig) GetServer() Server {
	return c.Server
}
