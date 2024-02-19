package main

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    logger "github.com/sirupsen/logrus"
)

type Config struct {
	Port int `yaml:"port"`
	LogLevel string `yaml:"logLevel"`
}

var config Config


func getLogLevel(level string) logger.Level {
	switch level {
		case "DEBUG":
			return logger.DebugLevel
		case "INFO":
			return logger.InfoLevel
		case "ERROR":
			return logger.ErrorLevel
		default:
			return logger.ErrorLevel
	}
}


func loadConfig() {
    file, err := ioutil.ReadFile("config.yaml")
    
    if err != nil {
	logger.Error(err)

 	return  
    }

    err = yaml.Unmarshal(file, &config)
    
    if err != nil {
	logger.Error(err)	 

	return
    } 	 

    logger.WithFields(logger.Fields{
	"port": config.Port,
	"logLevel": config.LogLevel,
    }).Info("Config loaded")

}
