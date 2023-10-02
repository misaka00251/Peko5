package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	TenantID     string `json:"TenantID"`
	ClientID     string `json:"ClientID"`
	ClientSecret string `json:"ClientSecret"`
	AdminID      string `json:"AdminID"`
}

var config Config

func ReadConfigFile(fileName string) Config {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Please create a config file first.")
	}

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
	fmt.Println("Your config file seems ready.")
	defer jsonFile.Close()

	return config
}
