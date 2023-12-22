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
	use   = flag.String("u", "live", "live or video")
	room  = flag.String("r", "12306", "room num")
	file  = flag.String("f", "flvfile/333.flv", "push stream file")
	p     = flag.Int("p", 1, "push or pul, push 1 and pull 2")
	proto = flag.Int("pr", 1, "rtmp 1 and http 2")
)

const (
	PullDomain = "pull1.eo6rgej.com"
	PushDomain = "push1.abcjzfyh8.com"
	PushKey    = ""
	PUllKey    = ""
)

func main() {
	flag.Parse()

	//dns1()

	//
	Push2(*file)

}

func dns1() {
	if *p == 1 {
		pushStream()
	} else {
		pullStream()
	}

}

func pushStream() {
	url := getUrl(PushDomain, PushKey)
	fmt.Println("ffmpeg", "-re", "-i", *file, "-c", "copy", "-f", "flv", url)
	//推流
	err := exec.Command("ffmpeg", "-re", "-i", *file, "-c", "copy", "-f", "flv", url).Run()
	if err != nil {
		panic(err)
	}
}

func pullStream() {
	url := getUrl(PullDomain, PUllKey)
	file := fmt.Sprintf("flvfile/%v.%s", time.Now().Unix(), "flv")
	fmt.Println("ffmpeg", "-re", "-i", url, "-c", "copy", file)
	//推流
	err := exec.Command("ffmpeg", "-re", "-i", url, "-c", "copy", file).Run()
	if err != nil {
		panic(err)
	}
}

func getUrl(domain, key string) string {

	now := time.Now().Unix()

	sign := generateSign(key, now)
	var url string
	if *proto == 1 {
		url = fmt.Sprintf("rtmp://%s/%s/%s?sign=%s&t=%d", domain, *use, *room, sign, now)
	} else {
		url = fmt.Sprintf("http://%s/%s/%s?sign=%s&t=%d", domain, *use, *room, sign, now)
	}

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

//func getAuthKey1() {
//	secret := "your_secret"
//	expire := 30
//	category := 1
//	appName := "your_app_name"
//	pushDomain := "your_push_domain"
//
//	timestamp := time.Now().Add(time.Minute*time.Duration(expire)).Unix()
//	sourceStr := fmt.Sprintf("%s:%s:%d:%s", pushDomain, appName, timestamp, secret)
//
//	md5Hash := fmt.Sprintf("%x", hashMD5([]byte(sourceStr)))
//	token := fmt.Sprintf("%d-%d-%s", timestamp, category, md5Hash)
//
//	fmt.Println(token)
//}
