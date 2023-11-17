package config

type AppConfig struct {
	AppName            string `envconfig:"APPNAME" default:"subscription"`
	AppEnv             string `split_words:"true" required:"true"`
	LogLevel           string `split_words:"true" required:"true"`
	LogChannel         string `split_words:"true" default:"daily"`
	LogFilePath        string `split_words:"true" default:""`
	LogFormat          string `split_words:"true" default:""`
	IsProduction       bool   `split_words:"true" default:"false"`
	RestPort           string `split_words:"true" default:"80"`
	TimeoutSecond      int    `split_words:"true" default:"5"`
	HashCost           int    `split_words:"true" default:"10"`
	JWTSecret          string `split_words:"true" default:"noh-secret"`
	RefreshTokenLength int    `split_words:"true" default:"10"`
}
