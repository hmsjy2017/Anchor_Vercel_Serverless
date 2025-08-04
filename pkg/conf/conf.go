package conf

import (
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/spf13/viper"
	"os"
)

// Config 全局配置
var Config = new(config)

type config struct {
	System  system  `yaml:"system"`
	Proxy   proxy   `yaml:"proxy"`
	CORS    cors    `yaml:"cors"`
}

type system struct {
	Name    string `yaml:"name"`
	Listen  string `yaml:"listen"`
	Debug   bool   `yaml:"debug"`
}

type proxy struct {
	Github  github  `yaml:"github"`
	Npm     npm     `yaml:"npm"`
	Wp      wp      `yaml:"wordpress"`
}

type github struct {
	Open        bool     `yaml:"open"`
	Minify      string   `yaml:"minify"`
	Endpoint    []string `yaml:"endpoint"`
	WhiteList   []string `yaml:"white_list"`
}

type npm struct {
	Open        bool     `yaml:"open"`
	Minify      string   `yaml:"minify"`
	Endpoint    []string `yaml:"endpoint"`
	WhiteList   []string `yaml:"white_list"`
}

type wp struct {
	PluginOpen  bool     `yaml:"plugin_open"`
	ThemeOpen   bool     `yaml:"theme_open"`	
	Minify      string   `yaml:"minify"`
	PluginWhiteList []string `yaml:"plugin_white_list"`
	ThemeWhiteList  []string `yaml:"theme_white_list"`
}

type cors struct {
	AllowOrigins    []string `yaml:"allow_origins"`
	AllowMethods    []string `yaml:"allow_methods"`	
	AllowHeaders    []string `yaml:"allow_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	ExposeHeaders   []string `yaml:"expose_headers"`
}

// Init 初始化配置
func Init(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// Read from environment variables if file not found or for specific overrides
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		utils.Log().Warn("Failed to read config file %s, using defaults or environment variables: %s", path, err)
	}

	if err := v.Unmarshal(Config); err != nil {
		utils.Log().Panic("Failed to unmarshal config: %s", err)
	}

	// Set default values if not provided by file or environment variables
	if Config.System.Name == "" {
		Config.System.Name = "Anchor"
	}
	if Config.System.Listen == "" {
		Config.System.Listen = ":8080"
	}

	// Example: Override proxy settings from environment variables
	if os.Getenv("PROXY_GITHUB_OPEN") == "true" {
		Config.Proxy.Github.Open = true
	}
	// Add more environment variable overrides as needed
}


