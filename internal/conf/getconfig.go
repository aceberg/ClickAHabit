package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/ClickAHabit/internal/check"
	"github.com/aceberg/ClickAHabit/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8852")
	viper.SetDefault("THEME", "united")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("BTNWIDTH", "195px")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)
	config.BtnWidth, _ = viper.Get("BTNWIDTH").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("btnwidth", config.BtnWidth)

	err := viper.WriteConfig()
	check.IfError(err)
}
