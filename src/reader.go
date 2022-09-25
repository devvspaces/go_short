package src

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func ReadFile(name string) []byte {
	b, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

func ParseYaml(data []byte) (map[string]string, error) {
	m := []map[string]string{}
	ret := map[string]string{}

	err := yaml.Unmarshal(data, &m)
	if err == nil {
		for _, value := range m {
			ret[value["path"]] = value["url"]
		}
	}

	return ret, err
}
