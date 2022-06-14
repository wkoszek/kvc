package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ServerInfo struct {
	Ip       string
	Password string
	Db       string
}

//Save server ip and password in a config file
func WriteConfig(ipHost string, pass string, db string) {
	data := ServerInfo{ipHost, pass, db}
	dataJson, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := homeDir + "/.configrcli.json"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	err = ioutil.WriteFile(fileName, dataJson, 0644)
	if err != nil {
		log.Println(err)
	}
}

//Read configs from the config file
func ReadConfig() ServerInfo {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := homeDir + "/.configrcli.json"

	jsonData, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Println(err)
	}

	var data ServerInfo

	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		log.Println(err)
	}

	return data
}

//Check if config file exists
func CheckConfig() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := homeDir + "/.configrcli.json"

	if _, err := os.Stat(fileName); err == nil {
		return true
	} else {
		return false
	}
}
