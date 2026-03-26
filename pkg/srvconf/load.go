package srvconf

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var mu sync.Mutex

func Load[T Configuration](config T) T {
	mu.Lock()
	defer mu.Unlock()

	configPath, err := getConfigPath(config)
	if err != nil {
		return config
	}

	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("load config: read file: %w", err))
	}

	if err = viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("load config: unmarshal: %w", err))
	}

	return config
}

func getConfigPath[T Configuration](config T) (string, error) {
	path, err := filepath.Abs(fmt.Sprintf("./%s", config.GetDir()))
	if err != nil {
		return "", err
	}

	return path, nil
}
