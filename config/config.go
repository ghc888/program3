package config
import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

//用于通过api保存配置
var configFileName string

//整个config文件对应的结构
type Config struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`

	WebAddr     string `yaml:"web_addr"`
	WebUser     string `yaml:"web_user"`
	WebPassword string `yaml:"web_password"`

	LogPath     string       `yaml:"log_path"`
	LogLevel    string       `yaml:"log_level"`

}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	configFileName = fileName

	return ParseConfigData(data)
}