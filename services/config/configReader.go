package config

import (
	"io/ioutil"
	"log"
)
import "gopkg.in/yaml.v2"

func GetLedStripConfig(filename string) *LedStripConfig {
	var config LedStripConfig
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}

func GetEffectsConfig(filename string) *EffectsConfig {
	var config EffectsConfig
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
