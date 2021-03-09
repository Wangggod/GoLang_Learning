### 目标：实现HTTP服务端
#### 1、关于HTTP的实现
附上参考链接：https://cizixs.com/2016/08/17/golang-http-server-side/

#### 2、HTTPPool（Http连接池）
如果要进行http通信，则客户端client和服务端server之间要建立连接（三次握手四次挥手）。这就导致如果是一些小线程或者是大量节点通信，仍旧使用new一个连接的方式显得不太够用。
链接：https://www.cnblogs.com/xrq730/p/10963689.html

http的实现流程大致为：1、域名解析 2、TCP三次握手 3、客户端发送httpRequest 4、服务端响应httpRequest 5、客户端处理httpResponse 6、TCP四次挥手。

最早期的时候一次TCP只能用一个http（短连接），后来开始有了一个TCP多次连接（长连接）。

长短连接都是通信层的概念，而http是应用层协议。那么为了使通信层知道TCP通道要复用，则需要 1、客户端发送一个keep-alive的header表示需要保持连接 2、服务端收到之后进行回应表示可以保持连接 3、最后一次请求数据客户端声明关闭这个header（一般是15秒，100次）

#### 3、ServeHTTP（服务端实现）
链接：https://blog.csdn.net/weixin_35891744/article/details/113074332
首先定义了

`func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request)`

对于连接池p定义了方法ServeHTTP，传入了`w http.ResponseWriter, r *http.Request`。Go语言中客户端的请求信息都封装给Request，但是发送给客户端的响应不是Response，而是ResponseWriter，这是一个用来创建HTTP响应的接口（官方自带），其会返回一个response指针。因此这也能解释为什么传参的时候request是指针，而它不是。

>response 结构体定义和 ResponseWriter 一样都位于 server.go，感兴趣的同学可以去看下源码，不过由于 response 对外不可见，所以只能通过 ResponseWriter 接口访问它。两者之间的关系是 ResponseWriter 是一个接口，而 http.response 实现了它。当我们引用 ResponseWriter 时，实际上引用的是 http.response 对象实例。


##### 3.1、判断地址是否正确
```
if !strings.HasPrefix(r.URL.Path, p.basePath) {
        panic("HTTPPool serving unexpected path: " + r.URL.Path)
    }
```
用HasPrefix函数对于r.URL.Path进行判断，看起开头是否为p.basePath。

##### 3.2、分隔输出
```
parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
    if len(parts) != 2 {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

```
调用SplitN函数，将传入的地址（正确的）分隔成<basepath><groupname><key>

分隔之后parts分为0,1,2三块，组名可以通过parts[0]访问到，key值可以通过parts[1]访问到。

##### 3.3、获取缓存值
```

group := GetGroup(groupName)
    if group == nil {
        http.Error(w, "no such group: "+groupName, http.StatusNotFound)
        return
    }
    view, err := group.Get(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


```

