 ### 安装Go
#### Windows安装
>  直接默认安装位置，注意不要修改
  安装结束后要配置GOROOT、GOPATH、Path。
  其中GOROOT是指Go的安装路径，GOPATH是指Go语言项目的路径，Path是bin的位置。
  
>  安装结束之后可以通过在cmd中输入Go来确认是否有安装。
```
$ go env //打印Go默认所有的环境变量
$ go env GOPATH //打印具体某个环境变量的值
```

>GOROOT
即安装的绝对路径。ROOT表示Go语言的安装目录。在Windows中，GOROOT的默认值是C:/go，而在Mac OS或Linux中GOROOT的默认值是usr/loca/go，如果将Go安装在其他目录中，而需要将GOROOT的值修改为对应的目录。另外，GOROOT/bin则包含Go为我们提供的工具链，因此，应该将GOROOT/bin配置到环境变量PATH中，方便我们在全局中使用Go工具链

>GOPATH
>环境变量GOPATH用于指定我们的开发工作区(workspace),是存放源代码、测试文件、库静态文件、可执行文件的工作。在类Unix(Mac OS或Linux)操作系统中GOPATH的默认值是$home/go。而在Windows中GOPATH的默认值则为%USERPROFILE%\go(比如在Admin用户，其值为C:\Users\Admin\go)。
>之后所有的源码都会在GOPATH下，一般就一个目录一个项目。比如GOPATH/src/mymath，表示mymath这个应用或者可执行应用，这个根据package或者main来决定

>GOBIN
>环境变量GOBIN表示我们开发程序编译后二进制命令的安装目录。当我们使用go install命令编译和打包应用程序时，该命令会将编译后二进制程序打包GOBIN目录，一般我们将GOBIN设置为GOPATH/bin目录。

>交叉编译
>什么是交叉编译？所谓的交叉编译，是指在一个平台上就能生成可以在另一个平台运行的代码，例如，我们可以32位的Windows操作系统开发环境上，生成可以在64位Linux操作系统上运行的二进制程序。在其他编程语言中进行交叉编译可能要借助第三方工具，但在Go语言进行交叉编译非常简单，最简单只需要设置GOOS和GOARCH这两个环境变量就可以了。

#### 设置代理
>设置代理的主要目的是为了顺利安装Go Module
```

$ go env -w GOPROXY=https://goproxy.cn,direct //设置代理
$ go env -w GO111MODULE=on //开启go modules
```
关于module的更多信息在https://blog.csdn.net/u011897301/article/details/105758056/
>GOPROXY的意义
>一般来说，需要用到某些依赖包的时候，go会去访问github的仓库，然后clone到本地来，所以我们需要设置代理来加速访问。

##### 问题

>~~我的GOPATH设置在E:/Go，参照教程我在桌面新建了hello文件夹并设置了，将代码保存为main.go之后，在目录下运行go build发现无法成功，提示“go: cannot find main module; see 'go help modules'”
原因在于我之前设置的时候将module功能常开，这导致我这个hello工程里面并没有mod文件，导致检测失败。解决办法是设为自动开启，即有mod在开，没有就算了。~~

###### 疑问（2021-3-3）
>我还是不太清楚module的意义和GOPATH的关系。

>Module的出现就是为了弱化GOPATH，在此之前任何代码都要防止在PATH下，但实用Module的话则更便于管理。最大的意义就是不必一定在GOPATH中创建项目
