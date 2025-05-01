package envs

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	DriverName string
	Host       string
	User       string
	Password   string
	Database   string
	SourceName string
}

type KafkaConfig struct {
	Servers string
	GroupId string
}

type Config struct {
	Port  int
	DB    DBConfig
	Kafka KafkaConfig
}

func SetupEnvs() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.BindEnv("PORT")

	viper.BindEnv("DB_DRIVER_NAME")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_DATABASE")
	viper.BindEnv("DB_SOURCE_NAME")

	viper.BindEnv("KAFKA_BOOTSTRAP_SERVERS")
	viper.BindEnv("KAFKA_CONSUMER_GROUP_ID")
}

func GetEnvs() *Config {
	return &Config{
		Port: viper.GetInt("PORT"),
		DB: DBConfig{
			DriverName: viper.GetString("DB_DRIVER_NAME"),
			Host:       viper.GetString("DB_HOST"),
			User:       viper.GetString("DB_USER"),
			Password:   viper.GetString("DB_PASSWORD"),
			Database:   viper.GetString("DB_DATABASE"),
			SourceName: viper.GetString("DB_SOURCE_NAME"),
		},
		Kafka: KafkaConfig{
			Servers: viper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
			GroupId: viper.GetString("KAFKA_CONSUMER_GROUP_ID"),
		},
	}
}
