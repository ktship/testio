package testio

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type JsonFileIO struct {
	fname 	string
}

func NewJsonFileIO(fname string) *JsonFileIO {
	return &JsonFileIO{
		fname	:fname,
	}
}

func (io *JsonFileIO)Read(hkey string, hid string, hkey2 string, hid2 string) (map[string]interface{}, error) {
	var fname string
	if (hkey2 != "") {
		fname = fmt.Sprintf("%s_%s_%s", io.fname, hkey, hid)
	} else {
		fname = fmt.Sprintf("%s_%s:%s_%s:%s", io.fname, hkey, hid, hkey2, hid2)
	}

	r, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(r, &dat); err != nil {
		return nil, err
	}
	return dat, nil;
}

func (io *JsonFileIO)Write(hkey string, hid string, hkey2 string, hid2 string, data map[string]interface{}) error {
	var fname string
	if (hkey2 != "") {
		fname = fmt.Sprintf("%s_%s_%s", io.fname, hkey, hid)
	} else {
		fname = fmt.Sprintf("%s_%s:%s_%s:%s", io.fname, hkey, hid, hkey2, hid2)
	}
	b, _ := json.Marshal(data)
	err := ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (io *JsonFileIO)Delete(hkey string, hid string, hkey2 string, hid2 string) error {
	var fname string
	if (hkey2 != "") {
		fname = fmt.Sprintf("%s_%s_%s", io.fname, hkey, hid)
	} else {
		fname = fmt.Sprintf("%s_%s:%s_%s:%s", io.fname, hkey, hid, hkey2, hid2)
	}
	err := os.Remove(fname)
	if err != nil {
		return err
	}

	return nil
}
