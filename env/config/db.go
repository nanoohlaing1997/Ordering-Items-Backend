package config

import "time"

type DbConfig struct {
	DatabaseEngine       string        `split_words:"true" required:"true"`
	DatabaseConnLifetime time.Duration `split_words:"true" default:"3600s"`
	TestDbSuffix         string        `split_words:"true" default:"_testing"`
	OrderDB              string        `split_words:"true" required:"true"`
}
