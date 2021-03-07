### 什么是接口
这两天设计缓存淘汰的时候涉及到了一个优点模糊的概念：interface。

其实仔细想起来并不复杂。在理解接口interface之前，我们需要明确一点是“接口的引入更多是为了多人编程，而不是单人工作”。

首先是形象举例：接口就是个招牌，就像是KFC的牌子，它让程序在不知道对方是什么的情况下通过接口就可以直接了当的知道它里面是卖炸鸡的。

但是为什么我们不直接设置一个类，比如设计一家店（类），然后在定义一个卖炸鸡（函数），这不是也能实现上述功能吗？

当然不是不行，主要是麻烦。

想象一下当有了接口之后，我们两个人合作写代码。今天你不在，明天我出差，后天程序交工。怎么做呢？我负责写接口，比如我写了func Sorter interface{}，并且规定接口下要实现三个函数Len(),Less(i,j int) bool，Swap(i,j int)。剩下的你负责在具体的条件下实现这些接口的具体内容。这比起结构体（类）的成员函数要好实现得多。

> 你写接口你写实现，就不用写接口了。
我写接口你实现，接口不就用上了。我不给你规定好了，你怎么知道该实现哪些内容呢。
更进一步，我写接口你实现，你今天不在，我明天出差，程序后天交工，那我今天必须把调用这个接口的代码写好。所以就需要接口中有函数，有明确的函数签名。我写个接口，再把调用函数写好，明天你把接口实现了，传个实例进来，交工。
interface换个叫法就是contract，有点合同的意思。A实现了这个接口，代表A承诺能做某些事情。
B需要一些能做某些事情的东西，于是B要求，必须实现了A接口，才能被我调用。实际上也就是个“规范”。

举例：
```
package main
import (
    "fmt"
)
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
type Xi []int
type Xs []string
func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }
func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }
func Sort(x Sorter) {
    for i := 0; i < x.Len()-1; i++ {
        for j := i + 1; j < x.Len(); j++ {
            if x.Less(i, j) {
                x.Swap(i, j)
            }
        }
    }
}
func main() {
    ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
    strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
    Sort(ints)
    fmt.Printf("%v\n", ints)
    Sort(strings)
    fmt.Printf("%v\n", strings)
}
```