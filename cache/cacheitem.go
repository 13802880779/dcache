package cache

import (
	"sync"
	"time"
)

type ttltype int

const (
	RESET_TTL_ON_ACCESS ttltype = iota
	DO_NOT_RESET_TTL_ON_ACCESS
)

// define a cache item where the data is actually stored
type CacheItem struct {
	Key            interface{}
	Value          interface{}
	TTL            time.Duration //默认缓存5秒
	ttlType        ttltype       //0,访问命中时不重置存储周期，1，访问命中时重置缓存周期
	aboutToExpired func()
	CreatedAt      time.Time
	AccessAt       time.Time
	AccessCount    int64
	sync.RWMutex
}

func NewCacheItem(args ...interface{}) *CacheItem {
	now := time.Now()
	switch len(args) {
	case 2:
		return &CacheItem{Key: args[0], Value: args[1], TTL: 5 * time.Second, ttlType: 0, aboutToExpired: nil, CreatedAt: now, AccessAt: now, AccessCount: 0}
	case 3:
		tt := args[2].(time.Duration)
		return &CacheItem{Key: args[0], Value: args[1], TTL: tt, ttlType: 0, aboutToExpired: nil, CreatedAt: now, AccessAt: now, AccessCount: 0}

	case 4:
		tt := args[2].(time.Duration)
		ttype := args[3].(ttltype)
		return &CacheItem{Key: args[0], Value: args[1], TTL: tt, ttlType: ttype, aboutToExpired: nil, CreatedAt: now, AccessAt: now, AccessCount: 0}
	case 5:
		tt := args[2].(time.Duration)
		ttype := args[3].(ttltype)
		fn := args[4].(func())
		return &CacheItem{Key: args[0], Value: args[1], TTL: tt, ttlType: ttype, aboutToExpired: fn, CreatedAt: now, AccessAt: now, AccessCount: 0}
	default:
		return nil
	}
}

func (ci *CacheItem) SetAboutToExpiredCallback(fn func()) {
	ci.aboutToExpired = fn
}

func (ci *CacheItem) IsItemExpired() bool {
	//ci.RWMutex.Lock()
	//defer ci.RWMutex.Unlock()
	now := time.Now()
	switch ci.ttlType {
	case DO_NOT_RESET_TTL_ON_ACCESS:
		return now.Sub(ci.CreatedAt) >= ci.TTL
	case RESET_TTL_ON_ACCESS:
		return now.Sub(ci.AccessAt) >= ci.TTL
	default:
		return false
	}
}

func (ci *CacheItem) TimeToExpired() time.Duration {
	//ci.RWMutex.Lock()
	//defer ci.RWMutex.Unlock()
	now := time.Now()
	switch ci.ttlType {
	case DO_NOT_RESET_TTL_ON_ACCESS:
		return ci.TTL - now.Sub(ci.CreatedAt)
	case RESET_TTL_ON_ACCESS:
		return ci.TTL - now.Sub(ci.AccessAt)
	default:
		return 0 * time.Second
	}
}

func (ci *CacheItem) KeepAlive() {
	//ci.RWMutex.Lock()
	//defer ci.RWMutex.Unlock()
	ci.AccessAt = time.Now()
	ci.AccessCount++
}
