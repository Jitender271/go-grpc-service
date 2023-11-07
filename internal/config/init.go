package config

import "github.com/spf13/viper"

const(
	defaultConfigPath = "./config"
	configName = "app"
	configType = "env"

)

var Configurations AppConfig

func InitConfig(testMode bool) *AppConfig{
	viper.AddConfigPath(defaultConfigPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AutomaticEnv()
	if testMode{
		viper.AddConfigPath("../../config")
	}
	err := viper.ReadInConfig()
	if err != nil{
		panic("error reding config file"+ err.Error())
	}
	err = viper.Unmarshal(&Configurations)
	if err != nil{
		panic("error reding config"+ err.Error())
	}

	err = viper.Unmarshal(&Configurations.DbConfigs)
	if err != nil{
		panic("error reding db config"+ err.Error())
	}
	return &Configurations
}