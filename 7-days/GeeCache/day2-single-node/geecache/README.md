### 前言
在day-1中，我们简单实现了LRU算法。但目前还存在许多问题，即无法确保缓存数据被互斥访问。我们还需要添加锁来确保用来存放缓存的map在某一时刻不会被多方修改。

* * *

### 目标：使用互斥锁和更新缓存
#### 1、sync.Mutex
在第一天中我们实现了LRU的几个方法，分别是：
```
查找功能
func (c *Cache) Get(key string) (value Value, ok bool) {  }
淘汰功能
func (c *Cache) RemoveOldest(){  }
添加功能
func (c *Cache) Add(key string, value Value){  }
计算长度
func (c *Cache) Len() int{  }
```
这些方法都没有上锁，可能会出现一些奇怪的问题，因此我们要使用sync.Mutex来确保互斥访问。

#### 2、支持并发读写（byteview.go）
首先我们需要一个只读的结构来存放缓存值：
```
type ByteView struct {
    b []byte //b将会存储真实的缓存值，且为了防止被修改所以是只读的
}

```
struct ByteView中只有一个成员变量b，而且是byte型的数组，这样可以存储任意类型的数据，例如字符串、图片。（这个才是这种的缓存）。

##### 2.1、func (v ByteView) Len() int
在lru.go中，我们实现了Cache结构，并且要求被缓存的对象需要使用Value结构，即用Len int方法返回其占用内存的大小。
##### 2.2、func (v ByteView) ByteSlice() []byte
由于b被设定为只读，但我们仍然需要修改它应该怎么办呢？那就是新建一个同样为ByteView类型的v，将b直接拷贝给v。

#### 3、为lru.Cache实现并发结构（cache.go）
cache.go实现非常简单，通过import lru.gp（其中包含struct Cache），并在声明一个struct cache。其中cache的成员函数包括实现锁用的mu和实现lru用的lru（由于引入了lru.go，可以定义为lru）。

接着封装get和add。
```

func (c *cache) add(key string, value ByteView) {
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.lru == nil {
        c.lru = lru.New(c.cacheBytes, nil)
    }
    c.lru.Add(key, value)
}
func (c *cache) get(key string) (value ByteView, ok bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.lru == nil {
        return
    }
    if v, ok := c.lru.Get(key); ok {
        return v.(ByteView), ok
    }
    return
}


```

add中首先使用`c.mu.Lock()`来上锁，再使用`defer c.mu.Unlock()`进行延迟解锁。再通过`c.lru = lru.New(c.cacheBytes, nil)`实例化结构cache c(c包含了lru)。实例化结束之后便可以直接调用lru.go中定义的Add方法实现添加。

get也是同理。

#### 4、主体结构（geecache.go）
快速浏览一下当前的文件结构：
![fb7e6662ce260246c1d5f7bebf5da364.png](en-resource://database/908:0)

首先我们定义了struct Group。一个Group可以认为是一个缓存命名空间，拥有name、getter（缓存未命中时候的回到）和mainCache（属于cache.go中定义的cache结构）

##### 4.1、结构梳理。
在开始读代码前，我们有必要再次理清一下目前的结构。
首先在`lru.go`中，我们实现了LRU淘汰策略，Cache可以理解为一个缓存区（即所有的缓存数据、控制结构都在这里）。

接着为了实现并发读写，我们建立了byteview.go来把缓存值封装起来（即具体缓存的数据），也确保了缓存数据不会被随意更改。cache.go则在封装了lru.go的基础上实例化了cache（通过`c.lru = lru.New(c.cacheBytes, nil`)，因此cache中的 lru可以直接使用lru.go中定义的方法来实现并发存储。
最后我们要设计geecache.go，负责与外部进行交互，控制缓存的存取和获取的主流程。

##### 4.2、Getter（重要，待填坑）
当缓存不存在的时候，我们应该从哪里获取数据？