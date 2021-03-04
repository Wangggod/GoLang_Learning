#### 变量和内置数据类型
##### 变量
Go语言是静态类型的，变量声明的时候必须明确变量的类型。且最重要的是类型写在变量后面。
```
var a int //没有数据默认为0
var a int =1
var a =1
```
或者更简洁的
```
a:=1 //只能用于函数内部，不可用于外部
msg:="Hello World"
```
##### 简单类型
包括空值（nil），整数（int、uint），浮点类型（float32）、字节（byte），字符串（string）
##### 常量
```
const Pi float32 = 3.124223
```

##### 字符串
Go的字符串都是使用UTF-8的字符集编码，用""定义，类型为string。
```
var frenchHello string
var emptyString string='""
func test(){
    no,yes,maybe:="no","yes","maybe"
    japaneseHello：="Konichiwa"
    frenchHello:="Bonjour"
}
```

但是Go中的字符串是不可变的，即如果有var s string ="hello"，是无法用java或者其他语言那样通过s[0]=c来访问修改。
但是可以通过间接法，即声明一个c变量将s转为byte
```
c:=[]byte(s)
c[0]="c"
s2=string(c)
```
即string不可改，但是byte可以改。

![9035f1128ffc36c861d262741b28076a.png](en-resource://database/900:1)

##### 可以分组
例如
```

import "fmt"import "os"

const i = 100const pi = 3.1415const prefix = "Go_"

var i intvar pi float32var prefix string
```
可以写成
```

import(
        "fmt"
        "os"
)

const(
        i = 100
        pi = 3.1415
        prefix = "Go_"
)

var(
        i int
        pi float32
        prefix string
)
```

##### 规则

大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

##### 数组
var arr [1000] int
arr[0]=42
a:=[3]int{1,2,3}
doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

##### silce

在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice

slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度
```
slice := []byte {'a', 'b', 'c', 'd'}

```

slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i
```

/ 声明一个含有10个元素元素类型为byte的数组
var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

// 声明两个含有byte的slice
var a, b []byte

// a指向数组的第3个元素开始，并到第五个元素结束，
a = ar[2:5]
//现在a含有的元素: 
ar[2]、ar[3]和ar[4]

// b是数组ar的另一个sliceb = ar[3:5]
// b的元素是：ar[3]和ar[4]
```

##### map

map也就是Python中字典的概念，它的格式为map[keyType]valueType我们看下面的代码，map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。
```

var numbers map[string]int// 另一种map的声明方式
numbers = make(map[string]int)
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3

fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
// 打印出来如:第三个数字是: 3
```

* map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
* map的长度是不固定的，也就是和slice一样，也是一种引用类型内置的len函数同样适用于map，返回map拥有的key的数量
* map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
* map的初始化可以通过key:val的方式初始化值，同时map内置有判断是否存在key的方式通过delete删除map的元素：

#### 流程
##### if不需要加括号
```

if x > 10 {
        fmt.Println("x is greater than 10")
} else {
        fmt.Println("x is less than 10")
}
```
你甚至可以在if里面声明变量
```
if x := computedValue(); x > 10 {
        fmt.Println("x is greater than 10")
} else {
        fmt.Println("x is less than 10")
}

//这个地方如果这样调用就编译出错了，因为x是条件里面的变量fmt.Println(x)

```
##### goto
```

func myFunc() {
        i := 0
Here:   //这行的第一个词，以冒号结束作为标签
        println(i)
        i++
        goto Here   //跳转到Here去
}
```

##### for
```

package main

import "fmt"

func main(){
        sum := 0;
        for index:=0; index < 10 ; index++ {
                sum += index
        }
        fmt.Println("sum is equal to ", sum)
}
// 输出：sum is equal to 45
```


#### 函数
用func定义，也是一种变量
```

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
        //这里是处理逻辑代码
        //返回多个值
        return value1, value2
}
```

此外Go的函数能够返回多个值
```

package main

import "fmt"

//返回 A+B 和 A*Bfunc SumAndProduct(A, B int) (int, int) {
        return A+B, A*B
}

func main() {
        x := 3
        y := 4

        xPLUSy, xTIMESy := SumAndProduct(x, y)

        fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
        fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
```
还可以变参，即接受变参的函数有着不定数量的参数
```

func myfunc(arg ...int) {}
```
##### defer
延迟语句，执行完之后逆序执行。后进先出。
例如：
```

func ReadWrite() bool {
        file.Open("file")
// 做一些工作
        if failureX {
                file.Close()
                return false
        }

        if failureY {
                file.Close()
                return false
        }

        file.Close()
        return true
}
```
用了defer之后
```

func ReadWrite() bool {
        file.Open("file")
        defer file.Close()
        if failureX {
                return false
        }
        if failureY {
                return false
        }
        return true
}
```
#### Struct
```

type person struct {
        name string
        age int
}
```
```

type person struct {
        name string
        age int
}

var P person  // P现在就是person类型的变量了

P.name = "Astaxie"  // 赋值"Astaxie"给P的name属性.
P.age = 25  // 赋值"25"给变量P的age属性
fmt.Printf("The person's name is %s", P.name)  // 访问P的name属性.
```

##### Struct匿名字段

们上面介绍了如何定义一个struct，定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
```

package main

import "fmt"

type Human struct {
        name string
        age int
        weight int
}

type Student struct {
        Human  // 匿名字段，那么默认Student就包含了Human的所有字段
        speciality string
}

func main() {
        // 我们初始化一个学生
        mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

        // 我们访问相应的字段
        fmt.Println("His name is ", mark.name)
        fmt.Println("His age is ", mark.age)
        fmt.Println("His weight is ", mark.weight)
        fmt.Println("His speciality is ", mark.speciality)
        // 修改对应的备注信息
        mark.speciality = "AI"
        fmt.Println("Mark changed his speciality")
        fmt.Println("His speciality is ", mark.speciality)
        // 修改他的年龄信息
        fmt.Println("Mark become old")
        mark.age = 46
        fmt.Println("His age is", mark.age)
        // 修改他的体重信息
        fmt.Println("Mark is not an athlet anymore")
        mark.weight += 60
        fmt.Println("His weight is", mark.weight)
}
```

##### method
即如果我们定义了一个struct长方形，可以写一个函数area计算面积。但是如果这个时候有struct圆，还是要计算面积，则还需要写一个area_圆，非常麻烦且不好看。
可以用method解决
```

package main

import (
        "fmt"
        "math"
)

type Rectangle struct {
        width, height float64
}

type Circle struct {
        radius float64
}

func (r Rectangle) area() float64 {
        return r.width*r.height
}

func (c Circle) area() float64 {
        return c.radius * c.radius * math.Pi
}


func main() {
        r1 := Rectangle{12, 2}
        r2 := Rectangle{9, 4}
        c1 := Circle{10}
        c2 := Circle{25}

        fmt.Println("Area of r1 is: ", r1.area())
        fmt.Println("Area of r2 is: ", r2.area())
        fmt.Println("Area of c1 is: ", c1.area())
        fmt.Println("Area of c2 is: ", c2.area())
}
```
