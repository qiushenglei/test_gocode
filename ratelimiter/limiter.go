package ratelimiter

import (
	"sync"
	"time"
)

var limit = NewSlidingWindow()

type SlidingWindow struct {
	maxRequest   int         // 窗口下的最大请求数
	window       int         // 窗口大小
	timeSlice    int         // 时间间隔
	LastTimeSeek int         // 最后一次更新的时间余数
	countData    map[int]int // 每个时间间隔的统计数量
	mux          sync.Mutex  // 线程安全锁
}

// SlidingWindow 滑动窗口
func NewSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		window:       60,
		timeSlice:    1,
		LastTimeSeek: 0,
		countData:    make(map[int]int),
	}
}

// Entry 入口
func (s *SlidingWindow) Entry() {
	if s.count() {
		// todo::success
	} else {
		// todo::reject request
	}
}

func (s *SlidingWindow) count() bool {
	// 统计需要加锁，并发做到线程安全
	s.mux.Lock()
	defer s.mux.Unlock()

	//当前时间戳
	now := time.Now().Unix()

	// 获取余数，当前请求保存到哪个slice节点
	seek := int(now) % s.window

	// 判断跟上次余数是否一样，如果不一样说明要把上一轮的统计清除
	if seek != s.LastTimeSeek {
		s.countData[seek] = 1
	} else {
		s.countData[seek]++
	}

	var res int
	for _, v := range s.countData {
		res = res + v
	}

	if res > s.maxRequest {
		return false
	}
	return true
}
