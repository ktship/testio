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

func (io *testFileIO)Read(id uint32, id2 uint32) (chTime int64, num int32, err error) {
	name := fmt.Sprintf("%s_%d_%d", tioFileName, id, id2)
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

func (io *testFileIO)Write(id uint32, id2 uint32, chTime int64, num int32) error {
	name := fmt.Sprintf("%s_%d_%d", tioFileName, id, id2)

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

func (io *testFileIO)Del(id uint32, id2 uint32) error {
	name := fmt.Sprintf("%s_%d_%d", tioFileName, id, id2)
	err := os.Remove(name)
	if err != nil {
		return err
	}

	return nil
}
