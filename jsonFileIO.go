package testio

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const filename string = "jsondat"

type JsonFileIO struct {}

func New() *JsonFileIO{
	return &JsonFileIO{}
}

func (io *JsonFileIO)Read(id uint32, id2 uint32) map[string]interface{} {
	name := fmt.Sprintf("%s_%d_%d", filename, id, id2)
	r, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(r, &dat); err != nil {
		panic(err)
	}
	return dat;
}

func (io *JsonFileIO)Write(id uint32, id2 uint32, data map[string]interface{}) {
	name := fmt.Sprintf("%s_%d_%d", filename, id, id2)
	b, _ := json.Marshal(data)
	err := ioutil.WriteFile(name, b, 0644)
	if err != nil {
		panic(err)
	}
}
