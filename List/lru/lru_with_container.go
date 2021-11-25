package lru

import "container/list"

type Item struct {
	val   interface{}
	item  *list.Element
}

type LRU struct {
	capacity  int
	items    map[string]Item
	lru     *list.List
}

func NewLRU() *LRU {
	l := &LRU{
		items: make(map[string]Item),
		lru: new(list.List),
	}
	return l
}

func(l *LRU) Get(key string) (interface{},bool) {
	//如果存在，则更新节点到链表表头
	if v,f := l.items[key]; f{
		l.lru.MoveToFront(v.item)
		return v.val, true
	}
	//如果存在，返回nil,false
	return nil,false
}

func(l *LRU) Set(key, val string) {
	//进行插入，如果存在，则移动链表到表头
	if v,f:= l.items[key]; f {
		l.lru.MoveToFront(v.item)
		v.val = val
		l.items[key] = v
		return
	}

	//如果cache满了，则淘汰最后
	if l.lru.Len() > l.capacity {
		l.removeOldest()
	}

	//如果不存在，进行插入，然后更新到表头
	item := Item{
		val: val,
	}
	lruPos := l.lru.PushFront(key)
	item.item = lruPos
	l.items[key] = item
	l.lru.MoveToFront(item.item)
}

func(l *LRU) removeOldest() {
	value := l.lru.Remove(l.lru.Back())
	delete(l.items,value.(string))
}
