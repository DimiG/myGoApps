/**
 * ========================================================================
 *  File: main.go
 *  Creator: Dmitri G.
 *  Date: 2018-11-19
 *  Description: Test the Go lang yaml configuration
 * ========================================================================
 */

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	User1 string `yaml:"user_1"`
	User2 string `yaml:"user_2"`
}

func loadConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}

func main() {
	c, err := loadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%+v\n\n", c)
	fmt.Println(c.User1)
	fmt.Println(c.User2 + "\n")
}
