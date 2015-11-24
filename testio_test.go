package testio

import (
	"testing"
	. "gopkg.in/check.v1"
	"log"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type TableSuite struct {
	io 	*testio
}

func (s *TableSuite) SetUpSuite(c *C) {
}
func (s *TableSuite) SetUpTest(c *C) {
}
func (s *TableSuite) TearDownTest(c *C) {
}
func (s *TableSuite) TearDownSuite(c *C) {
	DeleteTestioFiles()
}

var _ = Suite(&TableSuite {
	io : New(),
})

func (s *TableSuite) Test000_BASIC_HASH_IO(c *C) {
	log.Println("# Test000_BASIC_HASH_IO")

	// Test Data
	tt := time.Now().Unix()
	data1 := map[string]interface{} {
		"createTime":int(time.Now().Unix()),
		"greeting": "hello",
		"ac": "test",
		"b": 1234,
	}
	data2 := map[string]interface{} {
		"greeting": "hello 3",
		"ac": "product",
		"c": 1111,
	}

	// 일단 데이터를 씀.
	var err error
	err = s.io.Write2Way(TEST_KEY_USER, "111", "", "", data1)
	if (err != nil) {
		c.Fatal(err)
	}

	// 1차적으로 쓴 내용 확인.
	resp, errRead := s.io.Read2Way(TEST_KEY_USER, "111", "", "")
	if (errRead != nil) {
		c.Fatal(err)
	}
	if (resp["createTime"] != int(tt)) {
		c.Fatalf(" createTime(%d) is not %d... type: %T", resp["createTime"], tt, resp["createTime"])
	}
	if (resp["greeting"] != "hello") {
		c.Fatalf(" greeting(%s) is not tt... type: %T", resp["greeting"], resp["greeting"])
	}
	if (resp["ac"] != "test") {
		c.Fatal("error")
	}
	if (resp["b"] != 1234) {
		c.Fatal("error")
	}

	// 2차적으로 데이터 갱신
	err = s.io.Write2Way(TEST_KEY_USER, "111", "", "", data2)
	if (err != nil) {
		c.Fatal(err)
	}
	// 2차적으로 갱신한 데이터 확인
	resp, errRead = s.io.Read2Way(TEST_KEY_USER, "111", "", "")
	if (errRead != nil) {
		c.Fatal(err)
	}
	if (resp["greeting"] != "hello 3") {
		c.Fatalf(" (%v)... type: %T", resp["greeting"], resp["greeting"])
	}
	if (resp["ac"] != "product") {
		c.Fatalf(" (%v) ... type: %T", resp["ac"], resp["ac"])
	}
	if (resp["b"] != 1234) {
		c.Fatalf(" (%v) is not tt... type: %T", resp["b"], resp["b"])
	}
	if (resp["c"] != 1111) {
		c.Fatalf(" (%v) is not tt... type: %T", resp["c"], resp["c"])
	}
}

func (s *TableSuite) Test001_BASIC_RANGE_IO(c *C) {
	log.Println("# Test001_BASIC_RANGE_IO")

	// Test Data
	tt := time.Now().Unix()
	data1 := map[string]interface{} {
		"createTime":int(time.Now().Unix()),
		"greeting": "hello",
		"ac": "test",
		"b": 1234,
	}
	data2 := map[string]interface{} {
		"greeting": "hello 3",
		"ac": "product",
		"c": 1111,
	}

	// 일단 데이터를 씀.
	var err error
	err = s.io.Write2Way(TEST_KEY_USER, "111", TEST_KEY_TASK, "0", data1)
	if (err != nil) {
		c.Fatal(err)
	}

	// 1차적으로 쓴 내용 확인.
	resp, errRead := s.io.Read2Way(TEST_KEY_USER, "111", TEST_KEY_TASK, "0")
	if (errRead != nil) {
		c.Fatal(err)
	}
	if (resp["createTime"] != int(tt)) {
		c.Fatalf(" createTime(%d) is not %d... type: %T", resp["createTime"], tt, resp["createTime"])
	}
	if (resp["greeting"] != "hello") {
		c.Fatalf(" greeting(%s) is not tt... type: %T", resp["greeting"], resp["greeting"])
	}
	if (resp["ac"] != "test") {
		c.Fatal("error")
	}
	if (resp["b"] != 1234) {
		c.Fatal("error")
	}

	// 2차적으로 데이터 갱신
	err = s.io.Write2Way(TEST_KEY_USER, "111", TEST_KEY_TASK, "0", data2)
	if (err != nil) {
		c.Fatal(err)
	}

	// 2차적으로 갱신한 데이터 확인
	resp, errRead = s.io.Read2Way(TEST_KEY_USER, "111", TEST_KEY_TASK, "0")
	if (errRead != nil) {
		c.Fatal(err)
	}
	if (resp["greeting"] != "hello 3") {
		c.Fatalf(" (%v)... type: %T", resp["greeting"], resp["greeting"])
	}
	if (resp["ac"] != "product") {
		c.Fatalf(" (%v) ... type: %T", resp["ac"], resp["ac"])
	}
	if (resp["b"] != 1234) {
		c.Fatalf(" (%v) is not tt... type: %T", resp["b"], resp["b"])
	}
	if (resp["c"] != 1111) {
		c.Fatalf(" (%v) is not tt... type: %T", resp["c"], resp["c"])
	}
}

/*
func (s *TableSuite) Test003_CacheIO_TTL(c *C) {
	log.Println("# Tests to TTL Cache Redis read/write item")

	// Test Data
	tt := time.Now().Unix()
	data1 := map[string]interface{} {
		"createTime":time.Now().Unix(),
		"s_greeting": "hello",
	}

	// 일단 데이터를 씀.
	var err error
	err = s.io.cio.writeHashItem(TEST_CACHE_NAME_USERS, "111", "", "", data1)
	if (err != nil) {
		c.Fatal(err)
	}

	// cache 내용 읽기 --------------------
	resp, errRead := s.io.cio.readHashItem(TEST_CACHE_NAME_USERS, "111", "", "")
	if (errRead != nil) {
		c.Fatal(errRead)
	}
	if (resp["createTime"] != strconv.Itoa(int(tt))) {
		c.Fatalf(" createTime(%d) is not %s...", tt, resp["createTime"])
	}
	if (resp["s_greeting"] != "hello") {
		c.Fatalf(" greeting(%s) is not tt... type: %T", resp["s_greeting"], resp["s_greeting"])
	}

	time.Sleep(time.Second * (time.Duration)(s.io.cio.GetTTL() + 1))

	// Expire 된 키는 소멸되어야함. resp nil체크
	resp, errRead = s.io.cio.readHashItem(TEST_CACHE_NAME_USERS, "111", "", "")
	if (errRead != nil) {
		log.Printf(" expired: %v", errRead)
	}
	if resp != nil {
		c.Fatalf(" Does NOT expired!!! WHY??? ")
	}
}
*/