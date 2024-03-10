package config

import (
	"fmt"
	"log"
	"os"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadConf(path string) error {

	filePath := fmt.Sprintf("./%s", path)

	if _, err := os.Stat(filePath); err != nil {
		log.Panic("app configuration file is not exist")
		return err
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic("can not read app config file, err : ", err.Error())
		return err
	}
	if err = yaml.Unmarshal(file, &Conf); err != nil {
		log.Panic("can not unmarshal app configuration yaml file, err : ", err.Error())
		return err
	}
	return nil

}
