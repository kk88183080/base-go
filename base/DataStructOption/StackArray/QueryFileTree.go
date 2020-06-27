package StackArray

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func GetAllX(path string, files []string, level int) ([]string, error) {
	levelStr := ""
	if level == 1 {
		levelStr += "|"
	} else {
		for i := level; i > 1; i-- {
			levelStr += "|--"
		}
	}

	read, e := ioutil.ReadDir(path)

	if e != nil {
		return files, errors.New("文件不可读取")
	}

	for _, file := range read {
		fullPath := path + "/" + file.Name()
		files = append(files, levelStr+fullPath)
		if file.IsDir() {
			files, _ = GetAllX(fullPath, files, level+1)
		}
	}

	return files, nil
}

/**
求文件所在目录的深度
*/
func QueryFileDeep(path string, level int) {
	fmt.Println("level:", level)

	levelStr := ""
	if level == 1 {
		levelStr += "|"
	} else {
		for i := level; i > 1; i-- {
			levelStr += "|--"
		}
	}

	read, e := ioutil.ReadDir(path)

	if e != nil {
		return
	}

	for _, file := range read {
		fullPath := path + "/" + file.Name()
		fmt.Println(levelStr + fullPath)
		if file.IsDir() {
			QueryFileDeep(fullPath, level+1)
		}
	}
}
