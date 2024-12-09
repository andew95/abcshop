package configs

import (
	"strings"

	"github.com/spf13/viper"
)

type conf struct {
	DbConfig dbConfig
}

type dbConfig struct {
	User     string
	Password string
	DbName   string
}

func GetConfig() *conf {
	viper.SetConfigName("config")                          // ชื่อ config file
	viper.AddConfigPath(".")                               // ระบุ path ของ config file
	viper.AutomaticEnv()                                   // อ่าน value จาก ENV variable  // แปลง _ underscore ใน env เป็น . dot notation ใน viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // อ่าน config
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &conf{
		DbConfig: dbConfig{
			User:     viper.GetString("postgresql.user"),
			Password: viper.GetString("postgresql.password"),
			DbName:   viper.GetString("postgresql.dbname"),
		},
	}
}

func GetConfigTest() *conf {
	return &conf{
		DbConfig: dbConfig{
			User:     "admin",
			Password: "S3cret",
			DbName:   "abcShop",
		},
	}
}
