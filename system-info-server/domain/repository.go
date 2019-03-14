package domain

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func parse_file(filename string) []SystemDef {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var ret []SystemDef
	err = json.Unmarshal(byteValue, &ret)
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

type Repository struct {
	Environments map[string][]SystemDef
}

func Load_Environment(env string, filename string, repo *Repository) bool {
	repo.Environments[env] = parse_file(filename)
	return true
}