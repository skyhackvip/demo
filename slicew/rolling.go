package rolling

import (
	"sync"
	"time"
)

const (
	expire int64 = 5
)

type SliceWindow struct {
	Buckets map[int64]float64
	Mutex   *sync.RWMutex
}

func NewSliceWindow() *SliceWindow {
	return &SliceWindow{
		Buckets: make(map[int64]float64),
		Mutex:   &sync.RWMutex{},
	}
}

func (sw *SliceWindow) Incr(i float64) {
	if i == 0 {
		return
	}
	sw.Mutex.Lock()
	defer sw.Mutex.Unlock()

	now := time.Now().Unix()
	if _, ok := sw.Buckets[now]; ok {
		sw.Buckets[now] += i
	} else {
		sw.Buckets[now] = i
	}
	sw.removeOldBuckets()
}

func (sw *SliceWindow) removeOldBuckets() {
	now := time.Now().Unix() - expire
	for timestamp := range sw.Buckets {
		if timestamp <= now {
			delete(sw.Buckets, timestamp)
		}
	}
}

func (sw *SliceWindow) Max(now time.Time) float64 {
	var max float64
	sw.Mutex.RLock()
	defer sw.Mutex.RUnlock()
	for timestamp, value := range sw.Buckets {
		if timestamp >= now.Unix()-expire {
			max = value
		}
	}
	return max
}

func (sw *SliceWindow) Sum(now time.Time) float64 {
	var sum float64
	sw.Mutex.RLock()
	defer sw.Mutex.RUnlock()
	for timestamp, value := range sw.Buckets {
		if timestamp >= now.Unix()-expire {
			sum += value
		}
	}
	return sum
}

func (sw *SliceWindow) Avg(now time.Time) float64 {
	return sw.Sum(now) / float64(expire)
}
