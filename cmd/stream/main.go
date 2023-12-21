package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os/exec"
	"time"
)

var (
	use  = flag.String("u", "live", "live or video")
	room = flag.String("r", "12306", "room num")
	file = flag.String("f", "222.flv", "push stream file")
	p    = flag.Int("p", 1, "push or pul, push 1 and pull 2")
)

const (
	PullDomain = "pull.eo6rgej.com"
	PushDomain = "push1.abcjzfyh8.com"
	PushKey    = ""
	PUllKey    = ""
)

func main() {
	flag.Parse()

	if *p == 1 {
		pushStream()
	} else {
		pullStream()
	}
}

func pushStream() {
	url := getUrl(PushDomain, PushKey)
	fmt.Println("ffmpeg", "-i", *file, "-f", "flv", url)
	//推流
	err := exec.Command("ffmpeg", "-re", "-i", *file, "-f", "flv", url).Run()
	if err != nil {
		panic(err)
	}
}

func pullStream() {
	url := getUrl(PullDomain, PUllKey)
	file := fmt.Sprintf("flvfile/%v.%s", time.Now().Unix(), "flv")
	fmt.Println("ffmpeg", "-i", url, "-c", "copy", file)
	//推流
	err := exec.Command("ffmpeg", "-i", url, "-c", "copy", file).Run()
	if err != nil {
		panic(err)
	}
}

func getUrl(domain, key string) string {

	now := time.Now().Unix()

	sign := generateSign(key, now)
	url := fmt.Sprintf("rtmp://%s/%s/%s?sign=%s&t=%d", domain, *use, *room, sign, now)
	return url
}

func generateSign(key string, now int64) string {
	str := fmt.Sprintf("%s/%s/%s%d", key, *use, *room, now)

	d := md5.New()
	d.Write([]byte(str))
	s := d.Sum(nil)
	//sign := hex.EncodeToString(s)

	sign := make([]byte, hex.EncodedLen(len(s)))
	hex.Encode(sign, s)
	return string(sign)
}
