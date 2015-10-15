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

func (io *testCacheIO)GetCacheTask(id uint32, id2 uint32) (chTime int64, num int32, err error) {
	name := fmt.Sprintf("%s_%d_%d", cioFileName, id, id2)
	r, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	var dat map[string]int64
	if err := json.Unmarshal(r, &dat); err != nil {
		panic(err)
	}
	chTime = dat["ct"]
	num = int32(dat["num"])

	return chTime, num, nil;
}

func (io *testCacheIO)PutCacheTask(id uint32, id2 uint32, chTime int64, num int32) error {
	name := fmt.Sprintf("%s_%d_%d", cioFileName, id, id2)

	dat := make(map[string]int64)
	dat["ct"] = chTime
	dat["num"] = int64(num)
	b, _ := json.Marshal(dat)
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
