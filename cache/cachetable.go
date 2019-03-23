package cache

import (
	"container/list"
	"sync"
	"time"
)

type CacheTable struct {
	MaxItems  int //max items to store,if 0, means no limit
	OnEvicted func(key interface{})

	items    *list.List
	itemsMap map[interface{}]*list.Element

	cleanTimer    *time.Timer
	cleanInterval time.Duration
	sync.RWMutex
}

func NewCacheTable(maxitems int) *CacheTable {
	ct := CacheTable{MaxItems: maxitems, items: list.New(), itemsMap: make(map[interface{}]*list.Element), cleanInterval: 5 * time.Second}
	ct.clean()
	return &ct
}

func (ctable *CacheTable) Add(citem *CacheItem) {
	if ctable == nil {
		ctable = NewCacheTable(0)
	}
	key := citem.Key
	if k, ok := ctable.itemsMap[key]; ok { //if key already exist
		val := ctable.itemsMap[k]
		ctable.items.MoveToFront(val)
		val.Value.(*CacheItem).Value = citem.Value
		return
	}

	ele := ctable.items.PushFront(citem)
	ctable.itemsMap[key] = ele

	if ctable.MaxItems > 0 && ctable.items.Len() > ctable.MaxItems {
		ctable.removeTheOldest()
	}

	return
}

func (ctable *CacheTable) Get(key interface{}) *CacheItem {
	if v, ok := ctable.itemsMap[key]; ok { //found
		ctable.items.MoveToFront(v)
		ci := v.Value.(*CacheItem)
		ci.KeepAlive()
		return ci
	}
	return nil
}

func (ctable *CacheTable) RemoveCacheItem(key interface{}) {
	//println("removing item :" + key.(string))
	if ele, ok := ctable.itemsMap[key]; ok { //命中缓存
		citem := ele.Value.(*CacheItem)
		if citem.aboutToExpired != nil {
			citem.aboutToExpired()
		}
		delete(ctable.itemsMap, key)
		ctable.items.Remove(ele)
		//	item
	}
}

func (ctable *CacheTable) removeTheOldest() {
	//ct.items.Remove()
	//println("cleaning the oldest")
	ele := ctable.items.Back()
	ci := ele.Value.(*CacheItem)
	ctable.RemoveCacheItem(ci.Key)

}

//定时启动缓存清理线程
//由于缓存清除线程为定期5秒启动，所以缓存生存周期可能有至多+5秒的误差

func (ctable *CacheTable) clean() {
	//println("cache-clean thread started!")
	if ctable.cleanTimer != nil {
		ctable.cleanTimer.Stop()
	}

	//smallestCleanupInterval := ctable.cleanupInterval
	//now :=time.Now()
	for k, v := range ctable.itemsMap {
		citem := v.Value.(*CacheItem)
		timeToExpired := citem.TimeToExpired()
		if timeToExpired <= 0*time.Second { //已经过期可以清除
			//println("remove item:", citem.Key.(string))
			ctable.RemoveCacheItem(k)

		}

	}

	if ctable.cleanInterval > 0 {
		ctable.cleanTimer = time.AfterFunc(ctable.cleanInterval, func() {
			//println("the time will restart at:", ctable.cleanupInterval)
			go ctable.clean()
		})
	}
}
