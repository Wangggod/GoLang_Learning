package lru

import "container/list"

/*
定义struct Cache，即我们要实现的Cache队列。
实现LRU有两个要点：1、是要维护一个Cache队列 2、统计访问数。
1、可以理解为维护一个队列，如果某条数据被访问过，根据LRU其不应该被淘汰，则将这个不被淘汰的元素移至队列末端
反之越少访问的元素则应该移动至队头等待删除。
2、可以在struct中再设置一个map，用来记录统计每个元素的最近被访问的次数
**/
type Cache struct {
	maxBytes  int64                         //Cache的最大容量
	nBytes    int64                         //当前使用的空间
	ll        *list.List                    //cache队列，Go语言标准库可以实现
	cache     map[string]*list.Element      //键是字符串，值是双向链表中对应的节点指针
	OnEvicted func(key string, value Value) //某条记录被删除时的回调函数
}

type entry struct {
	key   string
	value Value
}

type Value interface { //interface用来定义对象的一组行为
	Len() int //定义Value接口，用具返回值所用的内存大小
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//实现查找功能，第一步从字典中找到对应的双向链表节点，第二步将该节点移动到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//缓存淘汰，移除最近最少访问的节点
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//把一个新的cache加入到队列中
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
