package LFU

import "container/list"

type Item struct {
	key   int
	val   int
	count int
}

type LFU3 struct {
	mcache map[int]*Item
	list   map[int]*list.List
	minfrequent int
	capacity    int
}

func(l *LFU3) NewLFU3(c int) *LFU3 {
	return &LFU3{
		mcache: make(map[int]*Item,c),
		list:   make(map[int]*list.List,c),
		minfrequent: 0,
		capacity:  c ,
	}
}

/*
Get 总体步骤如下:
1.判断该key有没有在缓存中，没有返回-1
2.如果在缓存中，则进行缓存的更新操作，之后返回该key
3.缓存更新操作如下：
	1.删除原来频次的节点
	2.更新维护minfrequent次数，如果更新的key对应的count就是minfrequet且原来的频次链表中长度为0的话说明，更新的就是最低频次的key
	3.进行新节点的频次链表节点的添加，如果链表为空则新建，不为空则加入到该链表后
*/
func(l *LFU3) Get(key int) int {
	v,ok := l.mcache[key];
	if ok {
		l.UpdateCache(v)
		return v.val
	}
	return -1
}

func(l *LFU3) UpdateCache(item *Item) {
	//删除原来的节点
	for e := l.list[item.count].Front(); e!= nil; e = e.Next() {
		if e.Value.(int) == item.key {
			l.list[item.count].Remove(e)
			break
		}
	}
	//更新维护minfrequent次数
	if item.count == l.minfrequent && l.list[item.count].Len() == 0 {
		l.minfrequent = item.count+1
	}
	//新节点的添加
	cnt := item.count+1
	if l.list[cnt] == nil {
		l.list[cnt] = list.New()
	}
	l.list[cnt].PushBack(item.key)
}

/*
Set 总体步骤如下：
1. 判断该key有没有在缓存中，如有的话，更新value值，并且更新该key的缓存操作，直接用Get就行
2. 如果没有在缓存中，则判断缓存的大小有没有超出
3. 如果有超出，则进行淘汰策略，淘汰掉最低minfrequent对应的链表中的最后的值,如果使用频率一样，删除掉缓存最久的元素，然后将新节点加入链表
4. 如没有超出,设置count=1,判断1频次的链表是否存在，如不存在则进行新建，然后加入链表。
5. 更新维护minfrequent = 1
*/
func(l *LFU3) Set(key, val int) {
	v,ok := l.mcache[key]
	if ok {
		v.val = val
		l.mcache[key] = v
		l.Get(key)
		return
	}
	if len(l.mcache) >= l.capacity {
		//有超出，进行淘汰策略
		delteList := l.list[l.minfrequent]
		l.list[l.minfrequent].Remove(delteList.Front())
		delete(l.mcache,delteList.Front().Value.(int))
	}
	if l.list[1] == nil {
		l.list[1] = list.New()
	}
	l.list[1].PushBack(key)
	l.minfrequent = 1
}