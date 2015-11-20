package testio

import (
	"os"
	"fmt"
	"path/filepath"
	"log"
)

func IsExists(name string) (bool) {
	_, err := os.Stat(name)
	return err == nil
}

func FindByExt(dirname string, ext string) []string {
	dirname = dirname + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ret := make([]string, 0)
	dotExt := fmt.Sprintf(".%s", ext)
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == dotExt {
				ret = append(ret, file.Name())
			}
		}
	}
	return ret
}

func DeleteTestioFiles() {
	flist := FindByExt(".", "testio")
	for _, v := range flist {
		err := os.Remove(v)
		if err != nil {
			log.Printf("Can NOT Delete file : %s", v)
		}

	}
}