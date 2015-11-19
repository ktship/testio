package testio

import "strconv"

func New() *testio {
	return &testio{
		Ddbio: NewJsonFileIO("db"),
		cio: NewJsonFileIO("c"),
	}
}

type testio struct {
	Ddbio 	*JsonFileIO
	cio   	*JsonFileIO
}

const TEST_KEY_USER = "testuser"
const TEST_KEY_TASK = "testtask"

// -------------------------------------------------
// 디비와 캐쉬를 동시에 사용
// -------------------------------------------------
// 1. 캐쉬에서 읽고 키가 없으면,
// 2. 디비에서 읽음.
// 3. 디비에서 읽었으면, 캐쉬에 저장.
func (io *testio)Read2Way(hkey string, hid string, hkey2 string, hid2 string) (map[string]interface{}, error) {
	// 1. 캐쉬에서 읽고 키가 없으면,
	resp, err := io.cio.Read(hkey, hid, hkey2, hid2)
	if err != nil {
		return resp, err
	}
	// 값이 캐쉬에 이미 존재하면 바로 리턴함.
	if resp != nil {
		return resp, err
	}

	// 2. 캐쉬에 없으므로, 디비에서 읽음.
	resp, err = io.Ddbio.Read(hkey, hid, hkey2, hid2)
	if err != nil {
		return resp, err
	}

	// 3. 디비에서 읽었으면, 캐쉬에 저장.
	err = io.cio.Write(hkey, hid, hkey2, hid2, resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// 1. 디비 / 캐쉬에 쓰기
func (io *testio)Write2Way(hkey string, hid string, hkey2 string, hid2 string, updateAttrs map[string]interface{}) (error) {
	err := io.Ddbio.Write(hkey, hid, hkey2, hid2, updateAttrs)
	if err != nil {
		return err
	}
	err = io.cio.Write(hkey, hid, hkey2, hid2, updateAttrs)
	if err != nil {
		return err
	}

	return nil
}

func (io *testio)Del2Way(hkey string, hid string, hkey2 string, hid2 string) (error) {
	err := io.Ddbio.Delete(hkey, hid, hkey2, hid2)
	if err != nil {
		return err
	}
	err = io.cio.Delete(hkey, hid, hkey2, hid2)
	if err != nil {
		return err
	}

	return nil
}

// -------------------------------------------------
// user : taskbytime interface
// -------------------------------------------------
func (io *testio)ReadUserTask(uid int, tid int) (map[string]interface{}, error) {
	resp, err := io.Read2Way(TEST_KEY_USER, strconv.Itoa(uid), TEST_KEY_TASK, strconv.Itoa(tid))
	return resp, err
}

func (io *testio)WriteUserTask(uid int, tid int, updateAttrs map[string]interface{}) (error) {
	err := io.Write2Way(TEST_KEY_USER, strconv.Itoa(uid), TEST_KEY_TASK, strconv.Itoa(tid), updateAttrs)
	return err
}

func (io *testio)DelUserTask(uid int, tid int) (error) {
	err := io.Del2Way(TEST_KEY_USER, strconv.Itoa(uid), TEST_KEY_TASK, strconv.Itoa(tid))
	return err
}
