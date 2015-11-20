package testio

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
	if (hkey2 == "") {
		fname = fmt.Sprintf("%s_%s_%s.testio", io.fname, hkey, hid)
	} else {
		fname = fmt.Sprintf("%s_%s_%s_%s_%s.testio", io.fname, hkey, hid, hkey2, hid2)
	}

	r, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(r, &dat); err != nil {
		return nil, err
	}

	for k, v := range dat {
		switch v.(type) {
		case float64:
			dat[k] = strconv.Itoa(int(v.(float64)))
		case string:
		default:
			return nil, fmt.Errorf("NOT SUPPORT TYPE")
		}
	}
	return dat, nil;
}

func (io *JsonFileIO)Write(hkey string, hid string, hkey2 string, hid2 string, data map[string]interface{}) error {
	var fname string
	if (hkey2 == "") {
		fname = fmt.Sprintf("%s_%s_%s.testio", io.fname, hkey, hid)
	} else {
		fname = fmt.Sprintf("%s_%s_%s_%s_%s.testio", io.fname, hkey, hid, hkey2, hid2)
	}
	modifyData := make(map[string]interface{})
	if IsExists(fname) {
		resp, err := io.Read(hkey, hid, hkey2, hid2)
		if err != nil {
			return err
		}
		// 기존 데이터 먼저 읽고,
		for k, v := range resp {
			modifyData[k] = v
		}
	}
	// 새로운 데이터를 덮어 씌움
	for k,v := range data {
		modifyData[k] = v
	}

	b, _ := json.Marshal(modifyData)
	err := ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (io *JsonFileIO)Delete(hkey string, hid string, hkey2 string, hid2 string) error {
	var fname string
	if (hkey2 == "") {
		fname = fmt.Sprintf("%s_%s_%s.testio", io.fname, hkey, hid)
	} else {
		fname = fmt.Sprintf("%s_%s_%s_%s_%s.testio", io.fname, hkey, hid, hkey2, hid2)
	}
	err := os.Remove(fname)
	if err != nil {
		return err
	}

	return nil
}
