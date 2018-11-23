/**
 * ========================================================================
 *  File: main.go
 *  Creator: Dmitri G.
 *  Date: 2018-11-23
 *  Description: Test the Go lang yaml configuration (little bit improved)
 * ========================================================================
 */

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Titles struct {
		Title1 string `yaml:"Title1"`
		Title2 string `yaml:"Title2"`
	} `yaml:"Titles"`

	Message1 string `yaml:"Message1"`
	Message2 string `yaml:"Message2"`
}

func loadConfig(filename string) (Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func main() {
	c, err := loadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n"+"%+v\n\n", c)
	fmt.Printf("Title1: %#v\n", c.Titles.Title1)
	fmt.Printf("Title2: %#v\n", c.Titles.Title2)
	fmt.Println()
}
