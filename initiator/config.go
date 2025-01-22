package initiator

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/spf13/viper"
)

func InitViper(currentDir string) (error, *dto.Config) {

	var cfg dto.Config

	// currentDir, err := os.Getwd() // Get the current working directory
	// if err != nil {
	// 	log.Fatal("Failed to get current directory:", err)
	// }
	viper.AddConfigPath(currentDir + "/config")
	// viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Print("failed to read config", err)
		return err, &dto.Config{}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Print(err)
		return err, &dto.Config{}
	}

	return nil, &cfg

}
