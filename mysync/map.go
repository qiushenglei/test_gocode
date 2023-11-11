package mysync

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func readlog(smap map[int]string, wg *sync.WaitGroup) {
	file, _ := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND, 0600)
	buf := bufio.NewReader(file)
	i := 0
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(isPrefix)
		smap[i] = string(line)
		i++
	}

	//fmt.Println(smap)
	fmt.Printf("%v", smap)
}

func readfile(smap map[int]string, wg *sync.WaitGroup, filename string) {
	defer wg.Done()
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0600)
	buf := bufio.NewReader(file)
	i := 0
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(isPrefix)
		smap[i] = string(line)
		i++
	}
	//fmt.Println(smap)
	fmt.Printf("%v", smap)
}

func UnsafeMap() {
	smap := make(map[int]string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go readfile(smap, &wg, "log.log")
	go readfile(smap, &wg, "god.mod")

	wg.Wait()
}

func safeReadfile(smap *sync.Map, wg *sync.WaitGroup, filename string) {
	defer wg.Done()
	if filename == "go.mod" {
		fmt.Println("is mod")
	}
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0600)
	buf := bufio.NewReader(file)
	i := 0
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		smap.Store(strconv.Itoa(i), string(line))
		i++
	}
}

func SafeMap() {
	smap := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go safeReadfile(&smap, &wg, "log.log")
	go safeReadfile(&smap, &wg, "go.mod")
	wg.Wait()
	smap.Range(func(key any, value any) bool {
		fmt.Printf("%v %v\"", key, value)
		return true
	})
}
