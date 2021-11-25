package LFU

//双向链表的节点
type Node2 struct{
	key       int
	val       int
	count     int
	prev      *Node2
	next      *Node2
}

type ListNode2 struct {
	size     int
	//哨兵头尾节点
	head     *Node2
	tail     *Node2
}

type LFU4 struct {
	mcache    map[int]*Node2
	mlist     map[int]*ListNode2
	capacity      int
	minfrequent   int
}

func NewLFU4(c int) *LFU4 {
	return &LFU4{
		mcache: make(map[int]*Node2,c),
		mlist: make(map[int]*ListNode2,c),
		capacity: c,
		minfrequent: 0,
	}
}
/*
Get函数总体步骤：
1. 访问cache，判断key有没有在缓存中，如果不存在，返回-1
2. 如果key存在，则进行key对应的频次链表的更新，然后返回value值
3. 频次链表的更新如下：
	1.删除原来频次链表中的节点,更新长度值
	2.更新维护minfrequent的值，如果节点未更新的值等于minfrequent并且它对应的频次链表为空了，说明，该节点确实是最低频次对应的节点，进行更新最新频次+1
	3.添加该节点到对应的频次链表中，看频次链表是否为空，如果为空则新建链表加入，不为空则加入尾节点即可。
*/
func(l *LFU4) Get(key int) int {
	node,ok := l.mcache[key]
	if ok {
		//删除节点值
		l.remove(node)
		//更新节点
		l.Update(node)
		return node.val
	}
	return -1
}

func(l *LFU4) remove(node *Node2) {
	l.mlist[node.key].removeNode(node)
	l.mlist[node.key].size --
}

func(l *LFU4) Update(node *Node2) {
	if node.count == l.minfrequent && l.mlist[node.count].size == 0 {
		l.minfrequent ++
	}
	v,ok := l.mlist[node.key];
	if !ok {
		//创建链表
		v = &ListNode2{
			size : 0,
			head: &Node2{key:0,count: 0,val: 0, prev: nil,next: nil},
			tail: &Node2{key:0,count: 0,val: 0, prev: nil,next: nil},
		}
		v.head.next = v.tail
		v.tail.prev = v.head

		l.mlist[node.key] = v
	} else {
		//加入新的节点
		v.addNode(node)
		//更新长度值
		v.size++
	}
}

//从双向链表中删除一个节点
func(list *ListNode2) removeNode(node *Node2) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//添加节点到双向链表的尾端
func(list *ListNode2) addNode(node *Node2) {
	node.next = list.tail
	node.prev = list.tail.prev
	node.prev.next = node
	list.tail.prev = node
}

/*
Set函数的整体步骤如下：
1. 访问cache,判断key有没有在缓存中，如果存在，则更新频次链表以及更新value值
2. 如果不存在，判断缓存是否已满，如果满了则进行淘汰策略，淘汰掉minfrequent对应链表的头节点，然后进行新节点的插入
3. 如果存在，忽略第二步，直接进行新节点的插入
4. 新节点插入，判断频次为1的链表是否为空，如果为空则进行新建，不为空则直接进行插入
5. 更新minfrequent的值
*/
func(l *LFU4) Set(key, val int) {
	node,ok := l.mcache[key]
	if ok {
		node.val = val
		l.mcache[key] = node
		l.Get(key)
		return
	}
	if len(l.mcache) >= l.capacity {
		deleteNode := l.mlist[l.minfrequent].head.next
		l.mlist[l.minfrequent].removeNode(deleteNode)
		delete(l.mcache,key)
	} else {

	}
	if l.mlist[1] == nil {
		//新建
		l.mlist[1] = &ListNode2{
			size: 0,
			head: &Node2{key:0,count: 0,val: 0, prev: nil,next: nil},
			tail: &Node2{key:0,count: 0,val: 0, prev: nil,next: nil},
		}
		l.mlist[1].head.next = l.mlist[1].tail
		l.mlist[1].tail.prev = l.mlist[1].head
	}
	//插入新的节点
	list := l.mlist[1]
	newNode := &Node2{
		key: key,
		val: val,
		count: 1,
		prev: nil,
		next: nil,
	}
	list.addNode(newNode)
	list.size++
	l.mlist[1] = list

	l.minfrequent = 1
}
