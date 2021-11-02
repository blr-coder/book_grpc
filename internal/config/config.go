package config

import "github.com/BurntSushi/toml"

type postgres struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type Config struct {
	BindAdr  string   `toml:"bind_addr"`
	LogLevel string   `toml:"log_lvl"`
	Postgres postgres `toml:"postgres"`
}

func ParseConfig(path string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(path, c); err != nil {
		return nil, err
	}
	return c, nil
}
