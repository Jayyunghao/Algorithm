package lru

type LinkNode struct {
	key  string
	val  interface{}
	pre  *LinkNode
	next *LinkNode
}

type LRUCache struct {
	m   map[string]*LinkNode
	capacity  int
	//当做头尾的两个哨兵节点
	head  *LinkNode
	tail *LinkNode
}

func Constructor(c int) *LRUCache {
	l := &LRUCache{
		m: make(map[string]*LinkNode),
		capacity: c,
		head: &LinkNode{"",0,nil,nil},
		tail: &LinkNode{"",0,nil,nil},
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (l *LRUCache) Get(key string) (interface{},bool) {
	//如果访问存在，则移动节点到链表表头
	cache := l.m
	if v,f := cache[key]; f {
		l.MoveToFront(v)
		return v.val,true
	}
	//不存在返回nil,false
	return nil, false
}

func(l *LRUCache) Set(key string, val interface{}) {
	cache := l.m
	//查找该key是否存在，如存在就移动到头结点处,并更新值
	if v,f := cache[key];f{
		l.MoveToFront(v)
		v.val = val
		cache[key] = v
		return
	}
	//不存在，则判断是否超出缓存大小，超出进行淘汰值和该节点
	if len(cache) >= l.capacity {
		delete(cache,l.tail.key)
		l.tail.pre.pre.next = l.tail
		l.tail.pre = l.tail.pre.pre
	}
	//进行新节点的插入，并插入到头结点处
	newNode := new(LinkNode)
	newNode.key = key
	newNode.val = val

	newNode.next = l.head.next
	l.head.next.pre = newNode
	l.head.next = newNode
	newNode.pre = l.head

	l.m[key] = newNode
	return
}

func(l *LRUCache) MoveToFront(node *LinkNode) {
	//删除该节点
	node.pre.next = node.next
	node.next.pre = node.pre
	//移动到头结点位置
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
	node.pre = l.head
}
