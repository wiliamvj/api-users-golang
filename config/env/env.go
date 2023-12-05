package env

import (
  "github.com/spf13/viper"
)

var Env *config

type config struct {
  GoEnv       string `mapstructure:"GO_ENV"`
  GoPort      string `mapstructure:"GO_PORT"`
  DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadingConfig(path string) (*config, error) {
  viper.SetConfigFile("app_config")
  viper.SetConfigType("env")
  viper.AddConfigPath(path)
  viper.SetConfigFile(".env")
  viper.AutomaticEnv()

  err := viper.ReadInConfig()
  if err != nil {
    return nil, err
  }
  err = viper.Unmarshal(&Env)
  if err != nil {
    return nil, err
  }
  return Env, nil
}
