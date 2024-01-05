package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func getAuthKey(secret string, expire int, category int, appName, pushDomain string) string {
	// 计算未来时间戳
	timestamp := time.Now().Add(time.Duration(expire) * time.Minute).Unix()

	// 拼接字符串
	sourceStr := fmt.Sprintf("%s:%s:%d:%s", pushDomain, appName, timestamp, secret)

	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(sourceStr))
	md5Hash := hex.EncodeToString(hash[:])

	// 组合 token
	token := fmt.Sprintf("%d-%d-%s", timestamp, category, md5Hash)

	return token
}

func Push2(sourceFile string) {
	// 设置推流密钥和流路径
	key := "meeoy83X1UutUComcFv1a5PThMOBs"
	streamPath := "testStream4"
	domain := "push2.abcjzfyh8.com"
	token := getAuthKey(key, 60, 13, "livetest2", domain)

	// 构建完整的RTMP URL
	rtmpServerURL := fmt.Sprintf("rtmp://push2.abcjzfyh8.com/livetest2/%s?auth_key=%s", streamPath, token)

	cmd := exec.Command("ffmpeg", "-re", "-i", sourceFile, "-c", "copy", "-f", "flv", rtmpServerURL)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error obtaining stderr: %v\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(stderr)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error during command execution: %v\n", err)
		return
	}
}

func Pull2(filePath string) error {
	key := "fofDQUHRnCNmghWSnC36ZimFbBhkh7"
	streamPath := "testStream4"
	domain := "pull2.eo6rgej.com"
	token := getAuthKey(key, 60, 13, "livetest2", domain)
	url := fmt.Sprintf("http://%s/livetest2/%s.flv?auth_key=%s", domain, streamPath, token)
	// 发起请求
	for i := 0; i < 3; i++ {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// 创建文件
		//file, err := os.Create(filePath)
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer file.Close()

		counter := &CountingReader{
			reader:     resp.Body,
			lastReport: time.Now(),
		}

		_, err = io.Copy(file, counter)
		if err != nil {
			panic(err)
			return err
		}
	}
	return nil
}

type CountingReader struct {
	reader     io.Reader
	bytesRead  int64
	lastReport time.Time
}

func (cr *CountingReader) Read(p []byte) (int, error) {
	n, err := cr.reader.Read(p)
	cr.bytesRead += int64(n)

	// 每秒报告一次速率
	if time.Since(cr.lastReport) >= time.Second {
		fmt.Printf("Downloaded %d bytes (%.2f KB/sec)\n", cr.bytesRead, float64(n)/1024.0)
		cr.lastReport = time.Now()
	}
	return n, err
}
