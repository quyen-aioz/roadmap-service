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

type W3Storage struct {
	AccessKey      string
	SecretKey      string
	Endpoint       string
	PublicEndpoint string
	Bucket         string
	PathFolder     string
}

type BaseConfig struct {
	Server    Server
	SQLite    SQLite
	JWT       JWT
	SeedAdmin SeedAdmin
	W3Storage W3Storage
}

func (c BaseConfig) GetServer() Server {
	return c.Server
}
