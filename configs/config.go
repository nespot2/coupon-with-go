package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DriverName string `json:"driverName"`
	DataSource string `json:"dataSource"`
}

func New(profile string) (*Config, error) {

	cfg := new(Config)

	if len(profile) == 0 {
		profile = "local"
	}

	file, err := os.Open(fmt.Sprintf("./configs/config-%s.json", profile))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
