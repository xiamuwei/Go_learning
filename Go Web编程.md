# Go Web编程

网络协议



分层模型

OSI七层模型(理论上的标准) / TCP/IP (事实上的标准)



层与协议

![层与协议](E:\learning\Michael\截图\层与协议.jpg)





## socket编程

简介：网络中socket 数据传输是一种特殊的 I/O ，socket 也是一种文件描述符

常见的Socket类型有两种，流式Socket 和 数据报式Socket。流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；数据报式Socket是一种无连接的Socket ，对应于无连接的UDP应用



![](E:\learning\Michael\截图\Socket处理过程.png)



1. golang是实现TCP通信

   TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议

   ```go
   // 客户端
   package main
   
   import (
   	"fmt"
   	"net"
   )
   
   func main() {
   	// 主动连接服务器
   	conn, err := net.Dial("tcp", "127.0.0.1:8000")
   	if err != nil {
   		fmt.Println("net.Dial err = ", err)
   		return
   	}
   	defer conn.Close()
   
   	_, err = conn.Write([]byte("hello golang"))
   	if err != nil {
   		fmt.Println("conn.Write err = ", err)
   		return
   	}
   }
   ```

   ```go
   // 服务端
   package main
   
   import (
   	"fmt"
   	"net"
   )
   
   func main() {
   	// 监听ip端口
   	listener, err := net.Listen("tcp", "127.0.0.1:8000")
   	if err != nil {
   		fmt.Println("net.Listen err = ", err)
   		return
   	}
   	defer listener.Close()
   
   	// 阻塞等待用户链接，一旦收到链接，就返回conn
   	conn, err1 := listener.Accept()
   	if err1 != nil {
   		fmt.Println("net.Listen err1 = ", err1)
   		return
   	}
   
   	defer conn.Close()
   
   	// 接受用户的请求
   	var slice []byte = make([]byte, 1024)
   	count, err2 := conn.Read(slice)
   	if err2 != nil {
   		fmt.Println("net.Listen err2 = ", err2)
   		return
   	}
   
   	fmt.Println("得到的消息 = ", string(slice[:count]))
   }
   ```

   

2. golang是实现UDP通信

   UDP是一种**无连接**的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域

   ```go
   // 客户端
   package main
   
   import (
   	"fmt"
   	"net"
   )
   
   func main() {
   	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
   		IP:   net.IPv4(127, 0, 0, 1),
   		Port: 30000,
   	})
   
   	if err != nil {
   		fmt.Println("net.DialUDP err = ", err)
   		return
   	}
   	_, err = socket.Write([]byte("hello udp protocl"))
   	if err != nil {
   		fmt.Println("发送数据失败，err:", err)
   		return
   	}
   }
   
   ```

   ```go
   // 服务端
   package main
   
   import (
   	"fmt"
   	"net"
   )
   
   func main() {
   	listen, err := net.ListenUDP("udp", &net.UDPAddr{
   		IP:   net.IPv4(127, 0, 0, 1),
   		Port: 30000,
   	})
   	if err != nil {
   		fmt.Println("net.Listen err = ", err)
   		return
   	}
   	defer listen.Close()
   
   	buf := make([]byte, 1024)
   	num, addr, err1 := listen.ReadFromUDP(buf)
   	if err1 != nil {
   		fmt.Println("listen.ReadFromUDP err = ", err1)
   		return
   	}
   	fmt.Printf("data:%v addr:%v count:%v\n", string(buf[:num]), addr, num)
   
   }
   ```

   





## http编程

Http (超文本传输协议，hyperText Transfer Protocol) 协议通常承载于TCP协议之上，有时也承载于TLS或SSL 协议层之上，这个时候就成了我们常说的Https



http协议工作流程

1. 客户输入URL
2. 客户端会先请求DNS服务器，通过DNS服务器将URL解析获取相应的域名对应IP
3. 通过IP地址访问对应的服务器
4. 建立TCP链接，进行通信

![web工作流程](E:\learning\Michael\截图\web工作流程.jpg)

> 注：地址URL，全称为Unique Resource Location，用来表示网络资源，可以理解为网络文件路径
>
> URL格式如下:
>
> ```
> http://host[":"port][abs_path]
> 
> http://192.168.31.1/html/index
> ```
>
> URL的长度是有限制的，不同的服务器的限制值不太相同，但不能无限长





+ 请求报文格式

![请求报文格式](E:\learning\Michael\截图\请求报文格式.jpg)格式说明：

Http请求报文由==请求行，请求头，空行，请求体==4个部分组成，如下图所示

![http请求报文格式](E:\learning\Michael\截图\http请求报文格式.jpg)

1. 请求行

   请求行由方法字段、URL字段、HTTP协议版本字段3个部分组成，它们之间使用空格分隔。常用的HTTP请求方法有POST、GET

   + GET：
     + 当客户端要从服务器中读取某个资源时，使用GET方法。GET方法要求服务器将URL定位的资源放在响应报文的数据部分，返回给客户端
     + 使用GET方法时，请求参数和对应的值附加在URL后面，利用一个问号("?") 表示URL的结尾和请求参数的开始，传递参数长度受限制，因此GET方法不适合用于上传数据
     + 通过GET方法来获取网页时，参数会显示在浏览器地址栏上，因此保密性很差

   + POST：
     + 当客户端给服务端提供信息较多时可以使用POST方法，POST方法向服务器提交数据，比如完成表单数据的提交，将数据提交给服务器处理
     + GET一般用于获取/查询资源信息，POST会附带用户数据，一般用于更新资源信息。POST方法将请求参数封装在HTTP请求数据中，而且长度没有限制，因为POST携带的数据，在HTTP请求正文中，以key/value的形式出现，可以传输大量数据

2. 请求头

   请求头部为请求报文添加一些附加信息，由key/value对组成，每行一对，key和value之间使用冒号分隔

3. 空行

   最后一个请求头之后是空行，发送回车符和换行符，通知服务器以下不会再有请求头了

4. 请求体

   请求体不在GET方法中使用，而是在POST方法中使用。

   POST方法使用于需要用户填写表单的场景



+ 响应报文格式

  + 成功请求

  ![image-20220323142736650](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220323142736650.png)

  + 请求失败

  ![image-20220323142748608](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220323142748608.png)

格式说明：

HTTP响应报文由==状态行、响应头部、空行、响应体==4个部分组成，如下图所示

![Http响应报文格式](E:\learning\Michael\截图\Http响应报文格式.png)

1. 状态行

   状态行由HTTP协议版本字段、状态码、状态码的描述文本3个部分组成，它们之间使用空格分隔

   状态码分类

   | 状态码 | 含义                                               |
   | ------ | -------------------------------------------------- |
   | 1xx    | 表示服务器已接收了客户端请求，客户端可继续发送请求 |
   | 2xx    | 表示服务器已成功接收到请求并进行请求               |
   | 3xx    | 表示服务器要求客户端重定向                         |
   | 4xx    | 表示客户端的请求有非法内容                         |
   | 5xx    | 表示服务器未能正常处理客户端的请求而出现意外错误   |

   常见状态码

   | 状态码                    | 含义                                       |
   | ------------------------- | ------------------------------------------ |
   | ==200 OK==                | 客户端请求成功                             |
   | 400 Bad Request           | 表示服务器                                 |
   | 401 Unauthorized          | 未授权                                     |
   | 403 Forbidden             | 服务器拒绝服务                             |
   | ==404 Not Found==         | 请求的资源不存在                           |
   | 500 Internal Server Error | 服务器内部错误                             |
   | 503 Server Unavailable    | 服务器临时不能处理客户端请求(稍微可能可以) |

2. 响应头部

3. 空行

   最后一个响应头部之后是一个空行，发送回车符和换行符，通知服务器以下不再有响应头部

4. 响应体

   服务器返回给客户端的文本消息





golang标准库内建了net/http包，覆盖了HTTP客户端和服务端的具体实现。使用net/http包，我们可以不用考虑手动编写请求或者响应格式，很方便的编写客户端或服务端程序

服务端程序

```go
package main

import (
	"net/http"
)

// w ,给客户端回复数据；r，读取客户端发送的数据
func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("r.Method = ",r.Method)
	w.Write([]byte("hello world"))
}

func main() {
    // 注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/hello", hello)
    // 监听绑定
	http.ListenAndServe("127.0.0.1:8000", nil)
}
```

客户端程序

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
    // 使用get方法请求一个资源
	response, err := http.Get("http://127.0.0.1:8000/hello")
	if err != nil {
		fmt.Println("http.Get err = ", err)
	}
	defer response.Body.Close()

	var buf []byte = make([]byte, 1024)
    var tmp string
    for {
        count, err1 := response.Body.Read(buf)
		if err1 != nil {
			fmt.Println("http.Get err1 = ", err1)
		}
        if count == 0 {
            fmt.Println("读取内容结束")
            break
        }
        tmp += string(buf[:count])    
    } 
	
	fmt.Println("buf = ",tmp)
}
```





## Gin框架

> Cookie和Session简介
>
> + Cookie
>
>   + 作用：Http协议是无状态，这就意味着每次请求都是独立的，对于服务器来说，每次的请求都是全新的。状态可以理解为客户端和服务端在某次会话中产生的数据，那无状态的就以为这些数据不会被保留。会话中产生的数据又是我们需要保存的。为此，Cookie技术诞生。
>
>     在internet中，Cookie实际上是指小量消息，是由Web服务器构建的，将信息存储在用户计算机上(客户端)的数据文件。
>
>   + 原理
>
>     Cookie是由服务器端生成，发送给User-Agent(一般是浏览器)，浏览器会将Cookie 的key/value 保存到某个目录的文本文件内，下次请求同一网站时就发送该Cookie给服务器(前提是浏览器设置为启用Cookie)。Cookie 名称和值可以由服务器端开发自己定义，这样服务端可以知道该用户是否是合法用户以及是否需要重新登录等，服务器可以设置或读取Cookies中的包含信息，借此维护用户跟服务器会话中的状态
>
>   + 特点
>
>     1. 浏览器发送请求的时候，自动携带该站点之前存储的Cookie信息
>     2. 服务端可以设置Cookie数据
>     3. Cookie是针对单个域名的，不同域名之间的Cookie是独立的
>     4. Cookie数据可以设置过期时间，过期的Cookie数据会被系统清除
>
>   + 标准库net/http 中定义Cookie，它代表一个出现在HTTP响应头中Set-Cookie 的值里或者HTTP请求头中Cookie 的值的HTTP Cookie
>
>     ```go
>     type Cookie struct {
>         Name       string
>         Value      string
>         Path       string
>         Domain     string
>         Expires    time.Time
>         RawExpires string
>         // MaxAge=0表示未设置Max-Age属性
>         // MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
>         // MaxAge>0表示存在Max-Age属性，单位是秒
>         MaxAge   int
>         Secure   bool
>         HttpOnly bool
>         Raw      string
>         Unparsed []string // 未解析的“属性-值”对的原始文本
>     }
>     ```
>
>     
>
> + Session
>
>   + 作用
>
>     Cookie 虽然在一定程度上解决了"保持状态"的需求，但是由于Cookie本身最大支持4069字节，以及Cookie本身保存在客户端，可能被拦截或窃取，因此就需要一种新东西，能支持更多字节，并且保存在服务器，有较高的安全性，这就是Session
>
>   + 原理
>
>     基于HTTP协议的无状态特征，服务器根本就不知道访问者是"谁"，那么上述的Cookie就起到桥接的作用
>
>     用户登录成功之后，我们在服务端为每个用户创建一个特定的session和一个唯一的标识，它们一一对应。其中：
>
>     1. Session是在服务端保存的一个数据结构，用来跟踪用户的状态，这个数据可以保存在集群、数据库、文件中
>     2. 唯一标识通常称为Session ID 会写入用户的Cookie中
>
> + 总结
>
>   Cookie弥补了HTTP无状态的不足，让服务器知道来的人是"谁"；但是Cookie以文本形式保存在本地，自身安全性较差；所以我们就通过Cookie识别不同的用户，对应的在服务端为每个用户保存一个Session数据，该Session数据中能够保存具体的用户数据信息



### 简介

Gin是使用Go语言编写，基于httprouter开发的web框架

下载并安装

```shell
go get -u github.com/gin-gonic/gin
```

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
    // 相当于对http库中的listenAndServer以及handleFunc进行封装，用默认路由实现
	r := gin.Default()
	// GET请求方式， /hello 请求路径
	// 当客户端以GET方式请求 /hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
        // 返回JSON格式的数据，除此之外还可以是HTML之类的
        // gin.H 是一个map[string]interface{}
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// 启动HTTP服务器，默认在0.0.0.0：8080启动服务
    // 如果想改变服务端口及地址，可以在此处修改
	r.Run()
}


// GET中的函数可以不适用匿名函数,如下
func sayHello(c *gin.Context){
    // c.JSON(http.StatusOK, gin.H{
    c.JSON(200,gin.H{
        "message":"hello world",
    })
}
```







### RESTful API

+ 简介

  REST是Representational State Transfer的简称，”表征状态转移“或者”表现层状态转化“

  REST的含义就是客户端与Web服务器之间进行交互的时候，使用Http协议中的4个请求方法代表不同的动作

  1. GET用来获取资源
  2. POST用来新建资源
  3. PUT用来更新资源
  4. DELETE用来删除资源

因为RESTful API，可以简化URL设计

原本设计

| 请求方法 |     URL      |     含义     |
| :------: | :----------: | :----------: |
|   GET    |    /book     | 查询书籍信息 |
|   POST   | /create_book | 创建书籍记录 |
|   POST   | /update_book | 更新书籍信息 |
|   POST   | /delete_book | 删除书籍信息 |

按照RESTful API设计如下

| 请求方法 |  URL  |     含义     |
| :------: | :---: | :----------: |
|   GET    | /book | 查询书籍信息 |
|   POST   | /book | 创建书籍记录 |
|   PUT    | /book | 更新书籍信息 |
|  DELETE  | /book | 删除书籍信息 |

```go
func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
}
```







### Gin渲染





### 获取参数





### 文件上传





### 重定向



### Gin路由





### Gin中间件











# GORM 

类似于mybatis的orm框架，能够帮助你编写sql语句

