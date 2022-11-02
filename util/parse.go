package util

import (
	yaml "gopkg.in/yaml.v2"
)

var Parse *parse

type parse struct{}

func init() {
	Parse = &parse{}
}

func (p *parse) ParseYml(path string, result interface{}) error {
	var (
		fileData []byte
		err      error
	)
	if fileData, err = File.Read(path); err != nil {
		return err
	}

	return yaml.Unmarshal(fileData, result)
}
