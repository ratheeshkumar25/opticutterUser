package config

import "github.com/spf13/viper"

type Config struct {
	SECERETKEY   string `mapstructure:"JWTKEY"`
	Host         string `mapstructure:"HOST"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Database     string `mapstructure:"DBNAME"`
	Port         string `mapstructure:"PORT"`
	Sslmode      string `mapstructure:"SSL"`
	GrpcPort     string `mapstructure:"GRPCUSERPORT"`
	AdminPort    string `mapstructure:"GRPCADMINPORT"`
	MateialPort  string `mapstructure:"GRPCMATERIALPORT"`
	SID          string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TOKEN        string `mapstructure:"TWILIO_AUTH_TOKEN"`
	SERVICETOKEN string `mapstructure:"SERVICE_TOKEN"`
	PHONE        string `mapstructure:"TWILIO_PHONE_NUMBER"`
	REDISHOST    string `mapstructure:"REDISHOST"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
