package Queue

import (
	"fmt"
	"io/ioutil"
)

func QueryFile(path string) {

	q := NewQueue()
	q.EnQueue(path)

	for !q.IsEmpty() {

		filePath := q.DeQueue().(string)
		files, e := ioutil.ReadDir(filePath)

		if e != nil {
			fmt.Println("path:", filePath, " read error")
			break
		}
		fmt.Println(filePath)
		fullName := ""
		for _, file := range files {

			fullName = filePath + "/" + file.Name()
			if file.IsDir() {
				q.EnQueue(fullName)
			} else {
				fmt.Println(fullName)
			}
		}
	}
}
