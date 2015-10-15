package testio

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

const tioFileName string = "testio"

type testFileIO struct {}

func NewTestFileIO() *testFileIO{
	return &testFileIO{}
}

func (io *testFileIO)Read(id uint32, id2 uint32) (taskVar map[string]interface{}, err error) {
	name := fmt.Sprintf("%s_%d_%d", tioFileName, id, id2)
	r, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(r, &dat); err != nil {
		panic(err)
	}

	return dat, nil;
}

func (io *testFileIO)Write(id uint32, id2 uint32, taskVar map[string]interface{}) error {
	name := fmt.Sprintf("%s_%d_%d", tioFileName, id, id2)
	b, _ := json.Marshal(taskVar)
	err := ioutil.WriteFile(name, b, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func (io *testFileIO)Del(id uint32, id2 uint32) error {
	name := fmt.Sprintf("%s_%d_%d", tioFileName, id, id2)
	err := os.Remove(name)
	if err != nil {
		return err
	}

	return nil
}
