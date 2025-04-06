package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

var c Config

type Option struct {
	Prefix       string
	Delimiter    string
	Separator    string
	YamlFilePath string
	CallbackEnv  func(string) string
}

func init() {

	k := koanf.New(defaultDelimiter)

	// load configuration from yaml file
	if err := k.Load(file.Provider(defaultYamlFilePath), yaml.Parser()); err != nil {
		log.Printf("error loading config from `config.yml` file: %s", err)
	}

	if err := k.Unmarshal("", &c); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}
}

func C() Config {
	return c
}

func New(opt Option) Config {
	k := koanf.New(opt.Separator)

	if err := k.Load(file.Provider(opt.YamlFilePath), yaml.Parser()); err != nil {
		log.Printf("error loading config from `config.yml` file: %s", err)
	}

	if err := k.Unmarshal("", &c); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	return c
}
