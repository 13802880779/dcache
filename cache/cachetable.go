package cache

import (
	"container/list"
	"sync"
	"time"
)

type CacheTable struct {
	Capacity  int //max items to store,if 0, means no limit
	OnEvicted func(key interface{})

	CacheList *list.List
	CacheMap  map[interface{}]*list.Element

	cleanTimer    *time.Timer
	cleanInterval time.Duration

	addCount int64
	getCount int64
	hitCount int64

	sync.RWMutex
}

func NewCacheTable(maxitems int) *CacheTable {
	ct := CacheTable{Capacity: maxitems, CacheList: list.New(), CacheMap: make(map[interface{}]*list.Element), cleanInterval: 5 * time.Second}
	ct.clean()
	return &ct
}
func (ctable *CacheTable) Size() int {
	ctable.Lock()
	defer ctable.Unlock()
	return ctable.CacheList.Len()
}

func (ctable *CacheTable) Add(citem *CacheItem) {
	ctable.Lock()
	defer ctable.Unlock()

	ctable.addCount++

	if ctable == nil {
		ctable = NewCacheTable(0)
	}
	key := citem.Key
	if k, ok := ctable.CacheMap[key]; ok { //if key already exist
		//val := ctable.itemsMap[k]
		ctable.CacheList.MoveToFront(k)
		k.Value.(*CacheItem).SetValue(citem.Value)
		return
	}

	ele := ctable.CacheList.PushFront(citem)
	ctable.CacheMap[key] = ele

	if ctable.Capacity > 0 && ctable.CacheList.Len() > ctable.Capacity {
		ctable.removeTheOldest()
	}

	return
}

func (ctable *CacheTable) Get(key interface{}) *CacheItem {
	ctable.Lock()
	defer ctable.Unlock()

	ctable.getCount++

	if element, ok := ctable.CacheMap[key]; ok { //found
		ctable.hitCount++
		ctable.CacheList.MoveToFront(element)
		ci := element.Value.(*CacheItem)
		ci.KeepAlive()
		return ci
	}
	return nil
}

func (ctable *CacheTable) RemoveCacheItem(key interface{}) {
	ctable.Lock()
	defer ctable.Unlock()
	//println("removing item :" + key.(string))
	if element, ok := ctable.CacheMap[key]; ok { //命中
		citem := element.Value.(*CacheItem)
		if citem.aboutToExpired != nil {
			citem.aboutToExpired()
		}
		delete(ctable.CacheMap, key)
		ctable.CacheList.Remove(element)
		//	item
	}
}

func (ctable *CacheTable) removeTheOldest() {
	//ct.items.Remove()
	//println("cleaning the oldest")
	ctable.Lock()
	defer ctable.Unlock()
	ele := ctable.CacheList.Back()
	ci := ele.Value.(*CacheItem)
	ctable.RemoveCacheItem(ci.Key)

}
func (ctable *CacheTable) GetHitRate() float64 {
	return float64(ctable.hitCount) / float64(ctable.getCount)
}

//定时启动缓存清理线程
//由于缓存清除线程为定期5秒启动，所以缓存生存周期可能有至多+5秒的误差

func (ctable *CacheTable) clean() {
	ctable.Lock()
	defer ctable.Unlock()
	//println("cache-clean thread started!")
	if ctable.cleanTimer != nil {
		ctable.cleanTimer.Stop()
	}

	//smallestInterval := ctable.cleanInterval
	//now :=time.Now()
	for k, v := range ctable.CacheMap {
		citem := v.Value.(*CacheItem)
		timeToExpired := citem.TimeToExpired()
		if timeToExpired <= 0*time.Second { //已经过期可以清除
			//println("remove item:", citem.Key.(string))
			ctable.RemoveCacheItem(k)
		}

	}

	//println(smallestInterval, ctable.cleanInterval)
	if ctable.cleanInterval > 0 {

		ctable.cleanTimer = time.AfterFunc(ctable.cleanInterval, func() {

			go ctable.clean()
		})
	}
}
