package main

/*
	作业: 参考 Hystrix 实现一个滑动窗口计数器
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SlidingWindow 秒级为单位的滑动窗口计数器
type SlidingWindow struct {
	buckets  map[int64]*bucket // 秒级为单位的桶
	interval int64             // 时间周期
	mu       *sync.RWMutex
}

type bucket struct {
	Value float64
}

// NewSlidingWindow 创建一个滑动窗口计数器
func NewSlidingWindow(interval int64) *SlidingWindow {
	return &SlidingWindow{
		buckets:  make(map[int64]*bucket),
		interval: interval,
		mu:       &sync.RWMutex{},
	}
}

func (w *SlidingWindow) currentBucket() *bucket {
	now := time.Now().Unix()

	// 当前时间有桶存在 直接返回
	if b, ok := w.buckets[now]; ok {
		return b
	}

	// 否则创建新的桶
	b := &bucket{}
	w.buckets[now] = b
	return b
}

func (w *SlidingWindow) removeOldBuckets() {
	t := time.Now().Unix() - w.interval
	for timestamp := range w.buckets {
		if timestamp <= t {
			delete(w.buckets, timestamp)
		}
	}
}

// Incr 累加
func (w *SlidingWindow) Incr(i float64) {
	if i == 0 {
		return
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	b := w.currentBucket()
	b.Value += i
	w.removeOldBuckets()
}

// Sum 累计
func (w *SlidingWindow) Sum() float64 {
	t := time.Now().Unix() - w.interval

	sum := float64(0)

	w.mu.RLock()
	defer w.mu.RUnlock()

	for timestamp, bucket := range w.buckets {
		if timestamp >= t {
			sum += bucket.Value
		}
	}

	return sum
}

// Max 最大值
func (w *SlidingWindow) Max() float64 {
	t := time.Now().Unix() - w.interval

	var max float64

	w.mu.RLock()
	defer w.mu.RUnlock()

	for timestamp, bucket := range w.buckets {
		if timestamp >= t {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

// Min 最小值
func (w *SlidingWindow) Min() float64 {
	t := time.Now().Unix() - w.interval

	var min float64

	w.mu.RLock()
	defer w.mu.RUnlock()

	for timestamp, bucket := range w.buckets {
		if timestamp >= t {
			if min == 0 {
				min = bucket.Value
				continue
			}
			if bucket.Value < min {
				min = bucket.Value
			}
		}
	}

	return min
}

// Avg 平均值
func (w *SlidingWindow) Avg() float64 {
	return w.Sum() / float64(w.interval)
}

func main() {
	// 窗口周期为10秒
	window := NewSlidingWindow(10)

	// 统计
	go func() {
		tick := time.Tick(1 * time.Second)
		for range tick {
			m := make(map[int64]float64)
			for t, v := range window.buckets {
				m[t] = v.Value
			}
			fmt.Println("buckets:", m)
			fmt.Println("max:", window.Max())
			fmt.Println("min:", window.Min())
			fmt.Println("sum:", window.Sum())
			fmt.Println("avg:", window.Avg())

		}
	}()

	// 每500ms累加一次数据
	for {
		n := rand.Intn(100)
		window.Incr(float64(n))
		time.Sleep(500 * time.Millisecond)
	}
}
