### 1、目标：实现LRU

* * *


#### 1.1、何为Cache？
cache，即缓存。在数据量少的时候，我们可以通过直接访问数据库的方式来读取数据。但随着数据量增加和访问量的提高，频繁的访问数据库会导致效率下降，并且数据库每秒能接受的请求次数也是有限的。cache的提出就是为了解决这个问题，它将部分数据存放在缓存中，比如直接放在内存里，这样访问到这部分数据的时候便直接从内存读取而不需要请求数据库，从而提高了访问效率。

* * *
#### 1.2、何为cache策略？
有了缓存之后，新的问题随之而来——缓存大小也是有限的，如果缓存满了怎么办？此时就需要一定的策略来决定哪些数据能够继续留在缓存里，而哪些应该被淘汰。实现一个分布式缓存的第一步就是要确定淘汰策略。淘汰策略有什么？回顾考研操作系统，常见替换策略有FIFO、LFU、LRU。
##### 1.2.1、FIFO
FIFO实现起来很简单，就是淘汰缓存中最早进来的记录。实现这一策略我们可以通过设计一个队列，新纪录添加到队尾，旧记录在队头等待淘汰。

FIFO的适用性场景主要是对局数据时效性要求低的时候，优先保证最新数据可用。
##### 1.2.2、LFU
最少使用策略，即每次都清除最少使用的那个，主要比较命中率。在保证高频数据有消息的场景下，可用选择这个策略。
##### 1.2.3、LRU
最近最少使用策略，即清除最后一次被使用的。该策略主要优先保证热点数据的有效性。即热点条件下更适用

#### 1.3、在哪存放cache？
无非是内存和硬盘，但技术上可以分为内存、硬盘文件、数据库。

#### 1.4、缓存分类
##### 1.4.1、本地缓存
即缓存设置在应用内，请求速度非常迅速，无需网络开销。单应用不需要集群支持或者集群情况下各节点无需互相通知的场景下使用本地缓存比较合适，但多个应用之间无法直接共享。
##### 1.4.2、分布式缓存
与应用分离，多个应用可以共享。

#### 1.5、推荐阅读
https://tech.meituan.com/2017/03/17/cache-about.html

* * *


### 2、代码解析

#### 2.1、type Cache struct{  }
代码如下：
```
type Cache struct {
    maxBytes  int64                         //Cache的最大容量
    nBytes    int64                         //当前使用的空间
    ll        *list.List                    //cache队列，Go语言标准库可以实现
    cache     map[string]*list.Element      //键是字符串，值是双向链表中对应的节点指针
    OnEvicted func(key string, value Value) //某条记录被删除时的回调函数
}
```
首先定义了结构体Cache，Cache中总共有maxBytes\nBytes\ll\cache\OnEvicted 五个值。其中值得注意的是OnEvicted是一个函数类型，其中涉及到了类型“Value”还未定义，这个会在后面进行解释。

缓存区有两个主要指标，一个是缓存区的总容量，令一个是当前用了多少缓存。LRU是最近最少使用的淘汰。不同于LFU（最少使用淘汰），LFU是为每个cache元素设置计数器，每次需要淘汰时选择计数器值最小的那个淘汰。LRU的原理并非统计次数，而是命中一次就往后移，每次只淘汰队伍第一个。
因此我们设置最大内存maxBytes，当前使用内存nBytes，可以简单理解为总位置数和已经占用的位置数。接着设置双向链表ll，用于实现淘汰。
##### 2.1.1、为什么要设置map？
如果仅仅是有一个双向链表，我们可以想象一个场景，当由新的数据访问，我们首先要确认这个数据是否在缓存队列中，如果仅仅依靠链表，则需要从头开始遍历链表，此时如果能通过数据的值直接定位到链表的位置是不是更快呢？因此我们为了加速访问还要引入字典map确保查找高效。

注意，map的作用不是统计，而仅仅是为了在O(1)的复杂度下定位到队列的位置。

Go中的map声明方式为`map[ key] value`，key为键值，在这里我们将存储缓存的具体值（比如文中是字符串），value值则为链表元素。这样当我们拿到一个值之后（实际中可以是图像、数据、音频等等），我们可以直接用map的查找确定这个在不在队列中，然后再进行后续的操作。

#### 2.2、type Value interface{  }
相关链接：https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/

```
type Value interface { //interface用来定义对象的一组行为
    Len() int //定义Value接口，用具返回值所用的内存大小
}
```
接口（interface）的引入最直接的就是上下层的模块不再需要依赖下层的具体模块，只需要用接口即可实现。

首先Value即为接口的方法签名。即这个接口从此就叫做Value，它不是成员变量之类的东西。

这个接口它只做一件事情，就是定义了一个函数Len( )，并且Len( )将会返回int。初次接触到我们可能会好奇它如何使用它。在Go官方文档中有Len()函数用来返回队列长度。


#### 2.3、type entry struct{ }
```
type entry struct {
    key   string
    value Value
}
```
定义结构体entry中有两个数据，一个是string型的key，和map的key是对应的。另一个是Value型的value，这样方便使用接口进行长度计算。

我们可以理解为，Cache是定义了“双向链表+map”的结构体，但当我们要插入一个新的Cache记录时，必然要包含“链表节点+map键值”

有一点好奇的是，为什么有了map的键值对应还要在链表中保留key值呢？map中的结构为[key: value]，key可以理解为map的编号，value理解为实际cache队列中的值。如果我们要淘汰头结点的时候，同时需要把map中的元素删除。假设链表没有设置key，那么则无法在map中找到对应的元素。所以这其实是一个双向对应的关系。



#### 2.4 func New(maxBytes int64, onEvicted func(string, Value)) *Cache{  }
```

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
    return &Cache{
        maxBytes:  maxBytes,
        ll:        list.New(),
        cache:     make(map[string]*list.Element),
        OnEvicted: onEvicted,
    }
}


```
当我们定义好Cache结构体之后，需要调用New函数为其分配对应的空间。其中链表自带new，map使用make分配空间。函数名直接定义即可。

#### 2.5、func (c *Cache) Get(key string) (value Value, ok bool)

```
func (c *Cache) Get(key string) (value Value, ok bool) {
    if ele, ok := c.cache[key]; ok {
        c.ll.MoveToFront(ele)
        kv := ele.Value.(*entry)
        return kv.value, true
    }
    return
}
```
当我们拿到一个值后，首先要判断它有没有在缓存中，因此第一步要设计一个查找方法。
首先查找的对象是` (c *Cache)`，输入为`（key string）`即通过键值直接查找。返回值为`(value Value, ok bool)`，

查找开始时首先抛出判断语句`if ele, ok := c.cache[key]; ok`，这句话的意思是首先我们输入的是键值key，即要查询的缓存值，可以是图像或者数值，这里我们就当做是定义里面的字符串。

拿到待查询的key之后，我们先传入`c.cache[key]`，即将一个Cache类型的c的成员变量cache[key]分别赋值给ele和ok。由于ok是bool型，如果这个值不存在，那么ok的值就是否，因此无法进入后续语句。。因为map中存放的value是链表节点，所以我们通过key访问map得到的ele就是链表中的节点。

如果key值查询成功，那么将这个`c.cache[key]`传给ele。然后调用`MoveToFront`将ele移动到队尾（注意这个函数的命名，不要被迷惑了）。然后将kv赋值为`ele.Value.(*entry)`，因为此时ele被是一个Cache类型，可以使用接口Value来获得这个节点的值。即用key和该节点的value，生成一个entry（可以理解为真实的cache）

#### 2.6、func (c *Cache) RemoveOldest()
该方法用于删除节点，即把队列的第一个弹出去。
```
func (c *Cache) RemoveOldest() {
    ele := c.ll.Back()
    if ele != nil {
        c.ll.Remove(ele)
        kv := ele.Value.(*entry)
        delete(c.cache, kv.key)
        c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
        if c.OnEvicted != nil {
            c.OnEvicted(kv.key, kv.value)
        }
    }
}
```
首先` ele := c.ll.Back()`直接获取队列第一个元素，暂时还不知道它是否符合条件。因此立刻使用` if ele != nil`判定其是否为空，如果不是则继续。

继续之后使用队列函数直接删除ele的节点，`c.ll.Remove(ele)`，删除之后用`kv := ele.Value.(*entry)`获取到这个被删除的节点的值，包括它的key和value，因为我们同时还要去删除map中的映射关系。接着使用`delete(c.cache, kv.key)`删除map与该节点的映射。删除结束后此时链表和map中都不存在这个数据了，因此我们要更新占用内存nBytes。我们把当前的nBytes - （key的长度+value的长度）即获得当前内存的值。（因为实际中我们存储cache时除了值还有映射关系，所以要减去两个）。

#### 2.7、func (c *Cache) Add(key string, value Value)
当我们往cache中新增值时应该怎么做呢？
```
func (c *Cache) Add(key string, value Value) {
    if ele, ok := c.cache[key]; ok {
        c.ll.MoveToFront(ele)
        kv := ele.Value.(*entry)
        c.nBytes += int64(value.Len()) - int64(kv.value.Len())
        kv.value = value
    } else {
        ele := c.ll.PushFront(&entry{key, value})
        c.cache[key] = ele
        c.nBytes += int64(len(key)) + int64(value.Len())
    }
    for c.maxBytes != 0 && c.maxBytes < c.nBytes {
        c.RemoveOldest()
    }
}
```
首先和查找一样，先判断是否有，如果有的话则移到队列后面。不一样的是这个时候可能要修改值，所以可能要修改大小和更新值。

如果这个值目前不在cache中，则把这个值插入，并且计算大小。

最后记得判断是否超出，如果超出的话就要移除掉队头的那个。
