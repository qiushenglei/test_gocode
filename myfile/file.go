package myfile

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func ReadFile() {

	// 读取整个文件内容到内存
	path, _ := os.Getwd()
	pc, path1, line, _ := runtime.Caller(0)
	fmt.Println(pc, line, path1)
	content, err := os.ReadFile(path + "\\go.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// 逐行读取文件内容
	file, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, _ := os.Stat("go.mod")
	data := make([]byte, info.Size())
	//方案1
	n, err := file.Read(data)
	//方案2
	//n, err := file.Read(data[len(data):cap(data)])
	if err != nil {
		fmt.Println(err.Error(), n)
	}

	data = data[0:info.Size()]
	//data := make([]byte, 0, info.Size())
	//for {
	//	if len(data) >= cap(data) {
	//		d := append(data[:cap(data)], 0)
	//		data = d[:len(data)]
	//	}
	//	n, err := file.Read(data[len(data):cap(data)])
	//	data = data[:len(data)+n]
	//	if err != nil {
	//		if err == io.EOF {
	//			err = nil
	//		}
	//		//return data, err
	//	}
	//}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func Ta() {

	//strings.NewReader()
	f, err := os.OpenFile("test", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer([]byte{'a', 'b'})
	b1 := strings.NewReader("987")

	// 读b的数据，写入file文件
	f.ReadFrom(b)
	f.ReadFrom(b1)
	//
	c := make([]byte, 2)
	f1, _ := os.Open("test") //前面f1是追加，是从最后开始读，导致第一次进去就遇到了EOF，所以创建了f1
	for err == nil {
		_, err = f1.Read(c)
		fmt.Println(c)
	}

	f.Write([]byte{'3', '4'})

}

func Tstring() {
	b2 := strings.NewReplacer("sdsdf", "sdfddf", "nihao", "hello") //
	a := b2.Replace("nihao fuck you sdsdf")                        // s这个串需要被替换，替换规则是上面的old被new替换
	fmt.Println(a)

	buff := bytes.NewBuffer([]byte{'1', '2'})
	b2.WriteString(buff, a)
	fmt.Println(a)
	fmt.Println(buff.String())
}
