package StackArray

import (
	"errors"
	"io/ioutil"
)

func GetAll(path string, files []string) ([]string, error) {

	read, e := ioutil.ReadDir(path)

	if e != nil {
		return files, errors.New("文件不可读取")
	}

	for _, file := range read {
		fullPath := path + "/" + file.Name()
		files = append(files, fullPath)
		if file.IsDir() {
			files, _ = GetAll(fullPath, files)
		}
	}

	return files, nil
}

func GetAllStack(path string) []string {

	stack := NewStack()
	stack.Push(path)

	files := []string{}
	for !stack.IsEmtpy() {

		v := stack.Pop().(string)
		//files = append(files, v) // 加入列表
		all, _ := ioutil.ReadDir(v)

		for _, file := range all {
			fullPath := v + "/" + file.Name()
			files = append(files, fullPath)
			if file.IsDir() {
				stack.Push(fullPath)
			}
		}
	}

	return files
}
