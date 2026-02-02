package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    DbURL       string `json:"db_url"`
    CurrentUser string `json:"current_user_name"`
}



func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Home Directory was not found")
	}
	my_string := homeDir + "/" + configFileName
	return my_string, nil
}

func write(cfg Config) error {
    data , err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	file_path , err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(file_path , data , 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config , error){
	file_path , err := getConfigFilePath()
	if err != nil {
		return Config{} , err
	}
	myfile , err := os.ReadFile(file_path)
	if err != nil {
		return Config{} , err
	}
	var cfg Config
	err = json.Unmarshal(myfile,&cfg)
	if err != nil {
		return Config{} , err
	}
	return cfg , nil
}

func (c *Config) SetUser(name string) error {
    c.CurrentUser = name
    return write(*c)
}



