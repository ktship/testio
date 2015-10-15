package testio

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

const cioFileName string = "cacheio"

type testCacheIO struct {}

func NewTestCacheIO() *testCacheIO{
	return &testCacheIO{}
}

func (io *testCacheIO)GetCacheTask(id uint32, id2 uint32) (taskVar map[string]interface{}, err error) {
	name := fmt.Sprintf("%s_%d_%d", cioFileName, id, id2)
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

func (io *testCacheIO)PutCacheTask(id uint32, id2 uint32, taskVar map[string]interface{}) error {
	name := fmt.Sprintf("%s_%d_%d", cioFileName, id, id2)
	b, _ := json.Marshal(taskVar)
	err := ioutil.WriteFile(name, b, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func (io *testCacheIO)DelCacheTask(id uint32, id2 uint32) error {
	name := fmt.Sprintf("%s_%d_%d", cioFileName, id, id2)
	err := os.Remove(name)
	if err != nil {
		return err
	}

	return nil
}
