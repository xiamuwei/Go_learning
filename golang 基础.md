 golang 基础

## 一、变量

变量声明

```go
var 变量名 [变量数据类型]  = 变量值
```



全局变量、局部变量

```go
// 全局变量
var a int 
var b int = 5
var c = 5 // 自动类型推导
// 同时声明多个变量
var (
    d int
    e string
    f float32
)

func main (){
    // 局部变量可以采用短变量声明
    short_variable := 5
    // 等价于
    var short_variable int 
    short_variable =
}
```



+ 常量 const PI int = 3



## 二、数据类型

+ **基本数据类型**
  + 整型 
  + 浮点型
  + byte （没有单独的字符型）
  + 布尔类型
+ 复合数据类型
  + 数组
  + 结构体
  + 函数
  + 接口
  + 切片
  + 通道
  + 接口
  + map

> 基本数据类型都有一个默认值
>
> | 数据类型 | 默认值 |
> | -------- | ------ |
> | 整型     | 0      |
> | 浮点型   | 0      |
> | 字符串   | ""     |
> | 布尔类型 | false  |



**类型转换**

golang在不同类型的变量之间赋值时需要==显示转换==，也就是说golang中数据类型==不能自动转换==

```go
/*
	基本语法：
		表达式T(v) 将值v转换为类型T
*/
// 一、基本数据类型之间的类型转换
// 数据类型的转化可以从范围小-->范围大，也可以从范围大-->范围小（不会报错，但会有精度损失）
var a int = 100
var b float32 = float32(a)

// 二、基本数据类型和string之间转换
/*	1、基本类型转string类型
	使用fmt.Sprintf("%参数", 表达式) ，根据format参数生成格式化的字符串并返回该字符串
	使用strconv 包的函数Format
*/ 
str := fmt.Sprintf("hello %v", "test")
fmt.Println("str = ", str)

var a int32 = 100
str1 := strconv.FormatInt(int64(a), 10) // 第二个参数为base，是进制，必须在2到36之间
fmt.Println("str1 = ", str1)
	
// 2、string类型转基本类型
/*	将string类型转换为基本数据类型时，要确保string类型能够转换成有效的数据
	使用strconv 包的函数Parse
*/ 
var c string = "true"
d, err := strconv.ParseBool(c)
if err != nil {
	fmt.Println("string转成基本类型失败，报错如下 = ", err)
}
fmt.Printf("d = %v ,d 的数据类型为%T ", d, d)

```



> 为了简化数据类型定义，golang支持自定义数据类型
>
> 基本语法: type 自定义数据类型别名 数据类型   // 相当于一个别名
>
> type myInt int  // 值得注意的是此时int 和 myInt会被golang认为是两种不同类型



## 三、运算符

+ 算术运算符

  ```go
  + - * / % 
  ++ -- //golang中没有前++，后++ 之分
  ```

  

+ 逻辑运算符

  ```go
  ||  //或，同假为假
  &&  //与，同真为真
  !   //非
  ```

  

+ 比较运算符 / 关系运算符

  ```go
  > < >= <= == !=
  ```

  

+ 赋值运算符

  ```go
  = += -= *= /= %= 
  >>= 
  <<=
  &=
  ^=
  !=
  ```

  

+ 位运算符

  ```go
  >> << 
  &
  |
  ^
  ```

  

+ 其他运算符 

  + &
  + *





## 四、流程控制

+ 判断语句

  + if 

    ```go
    func test_if() {
    	var num int = 5
    	if num > 10 {
    		fmt.Println("这个值大于10")
    	} else if num > 0 {
    		fmt.Println("这个值大于0，小于10")
    	} else {
    		fmt.Println("这个值是个负数")
    	}
    }
    ```

  + switch

    ```go
    func test_switch() {
    	var inter interface{}
    	var str string = "hello world"
    	inter = str
    
    	// 使用类型断言
    	switch inter.(type) {
    	case string:
    		fmt.Println("str的数据类型为string")
    	case int:
    		fmt.Println("str的数据类型为int")
    	case float32:
    		fmt.Println("str的数据类型为float32")
    	case bool:
    		fmt.Println("str的数据类型为bool")
    	default:
    		fmt.Println("str的数据类型未知")
    	}
    }
    ```

    

+ 循环语句

  + for 

    ```go
    func test_for() {
    	for i := 0; i < 10; i++ {
    		fmt.Println("i的值 = ", i)
    	}
        
        // for-range 语句遍历数组
        var arr [5]string = [5]string{"hi", "hello", "world", "go", "!"}
    	for k, v := range arr {
    		fmt.Printf("key = %v , value = %v\n", k, v)
    	}
        
        
        // golang 中没有 while 和 do while语法，但可以通过for实现类似效果
    	// 实现while
    	var count int = 5
    	for {
    		if count == 0 {
    			fmt.Println("循环结束...")
    			break
    		}
    		fmt.Println("count = ", count)
    		count--
    	}
    
    	count = 5
    	fmt.Println("count = ", count)
    	// 实现do while
    	for {
    		fmt.Println("count = ", count)
    		count--
    		if count == 0 {
    			fmt.Println("循环结束...")
    			break
    		}
    	}
    }
    ```

    

注意：go语言中没有 while 和 do while 语法，但是可以通过for循环来实现其使用效果

+ while循环的实现

  ![image-20220322163100323](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220322163100323.png)

+ do while循环的实现

  ![image-20220322163113733](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220322163113733.png)



> goto 
>
> break
>
> continue



## 五、函数

为完成某一功能的程序指令的集合，称为函数。作用可以提高代码复用性。

在golang中，函数分为：**自定义函数，系统函数**

语法：

```go
func 函数名(形参列表) (返回值列表){
    执行语句...
    return 返回值列表 //返回值按照情况可以没有，且可以同时返回多个值
} 
```



+ 递归函数

+ 匿名函数

  **golang支持匿名函数，匿名函数就是没有名字的函数**

  两种使用方式

  1. 在定义匿名函数时就直接调用，这种方法匿名函数只能调用一次

     ```go
     res := func(n1 int, n2 int) int {
     		return n1 + n2
     	}(2, 3)
     fmt.Println("两数的总和 res = ",res)
     ```

  2. 将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数

     ```go
     a := func(n1 int, n2 int) int {
         return n1 - n2
     }
     
     res2 := a(20, 10)
     res3 := a(30, 10)
     fmt.Println("res2 = ", res2)
     fmt.Println("res3 = ", res3)
     ```

+ init函数

  **完成一些初始化的工作**

  每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被golang运行框架调用，也就是说init会在main函数前被调用

  执行顺序是：==全局变量定义 --> init函数 --> main函数==

+ 闭包

  **闭包就是一个函数 和与其相关的引用环境组合的一个整体**

  > [闭包的使用场景](https://www.html.cn/qa/other/23216.html)
  >
  > 1. 使用闭包代替全局变量；
  > 2. 函数外或在其他函数中访问某一函数内部的参数
  > 3. 在函数执行之前为要执行的函数提供具体参数
  > 4. 在函数执行之前为函数提供只有在函数执行或引用时才能知道的具体参数
  > 5. 暂停执行
  > 6. 包装相关功能



**可变参数**



**引用传递和值传递**

值类型：基本数据类型、string、数组、结构体

引用类型：接口、函数、map、通道、切片、指针等



string中常用的系统函数

时间和日期中常用的系统函数



内置函数	

​	len： 用来求长度，比如string、array、slice、map、channel

​	new：用来分配内存，主要用来分配值类型，比如int、float32、struct...返回的是指针

​	make：用来分配内存，主要用来分配引用类型，比如channel、map、slice



## 六、错误处理

golang中引入的错误处理方式为：panic、defer、recover

golang中抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

```GO
func testForPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover err = ", err)
		}
	}()
	var a = 10
	var b = 0
	var c = a / b
	fmt.Println("c = ", c)

}
```





自定义错误

使用errors.New 内置函数

```go
func readConf(name string) (err error){
    if name == "config.ini"{
        return nil
    }else {
        return errors.New("读取文件失败...")
    }
}

fn test(){
    err := readConf("config2.ini")
    if err != nil {
        panic(err)
    }
}
```



## 七、结构体

**golang中结构体是值类型**

golang 中的结构体 和其他编程语言的类(class) 有同等地位，但没有传统面向对象编程里的继承、方法重载、构造函数 和 析构函数、隐藏的this指针等等

基本语法

type 结构体名称 struct{

​	field1 type   // **字段可以是基本数据类型，数组或者引用类型**

​	field2 type

}



单元结构体

元组结构体



方法

golang中的方法是作用在指定的数据类型上的(即：和指定的数据类型绑定)，因此自定义类型都可以有方法，而不仅仅是struct，比如int，float32等都可以有方法







## 八、数组

数组可以存放多个同一类型的数据

创建数组

```go
// 声明数组
var arr1 [5]int
// 往数组添加元素
arr1[0] = 1
arr1[1] = 2
arr1[2] = 3

// 定义数据
var arr2 [3]int = [3]int{1, 2, 3}
// 自动类型推导
var arr3 = [3]int{4, 5, 6}
// 定义数组时不显示指定数组大小，而是让编译器自动根据后面赋值推导
var arr4 = [...]int{7, 8, 9}
var arr5 = [...]int{0: 10, 1: 11, 2: 12}
fmt.Println("arr2 = ", arr2)
fmt.Println("arr3 = ", arr3)
fmt.Println("arr4 = ", arr4)
fmt.Println("arr5 = ", arr5)

// 使用for-range 遍历数组
for k, v := range arr1 {
    fmt.Printf("key = %v , value = %v\n", k, v)
}
```





## 九、切片

切片就是数组的一个引用，因此切片是引用类型。但是切片的长度可以变化，**因此切片是一个长度可变的动态数组**

底层来说，其实就是一个结构体

type slice struct {

​	ptr *[2]int

​	len int    // 长度

​	cap int   // 容量

}

声明切片语法

var slice []int

创建切片

```go
// 声明切片
// 1、使用切片引用一个已经创建好的数组
var slice1 = arr2[2:3]
fmt.Println("slice1 = ", slice1)
fmt.Println("slice1 = ", &slice1[0])

// 2、使用make创建一个切片
slice2 := make([]int, 5, 10) // 创建了一个[]int类型，大小为5，容量为10的切片
/*
	等价于
	var slice2 []int
	slice2 = make([]int,5,10)
	需要注意的是切片使用之前必须使用make分配内存空间！因为只是声明并不会分配内存，且只有分配内存后才能赋值和使用
*/
slice2[0] = 50 
// 使用append函数进行追加
slice2 = append(slice2, 100) 
fmt.Println("slice2 = ", slice2)
// 结果 slice2 =  [50 0 0 0 0 100]


// 切片的遍历，请参照数组的遍历
```

> 切片append操作的底层原理
>
> 切片append操作的本质就是对数组扩容，golang底层会创建一个新的数组newArr (安装扩容后大小)，将slice原来包括的元素拷贝到新的数组newArr ，slice重新引用到newArr



## 十、map

map 是key-value 数据结构，类似其他编程语言的集合

基本语法：

var 变量名 map[key_type]value_type

注：key除了slice、map、function不可以外，其他数据类型都可以，因为这几个无法使用==判断。通常为int、string

value的类型与key基本一样，通常为数字、string、map、struct



map声明

```go
// 方式1 
var map1 map[string]string
map1 = make(map[string]string, 5)
map1["a"] = "value1"
map1["b"] = "value2"
fmt.Println("map1 = ", map1)

// 方式2
map2 := map[string]string{
    "test1":"a",
    "test2":"b",
}
fmt.Println("map2 = ", map2)
```



map的增删改查

```go
// 增删改查
// 增或者改,没有就添加，有就使用value进行修改
map2["test1"] = "value"
map2["test3"] = "c"
fmt.Println("map2 = ", map2)
// 删除，使用delete，无法一次性全部清空map，可以遍历逐个删除
delete(map2, "test1")
fmt.Println("map2 = ", map2)
// 查
res, ok := map2["test3"]
if !ok {
    fmt.Println("没有找到查询结果")
} else {
    fmt.Println("查询结果 res = ", res)
}
```



map遍历参照数组的遍历



map排序

golang中的map默认是无序的，注意也不是按照添加的顺序存放的





## 十一、接口

golang中的多态特性主要是通过接口来体现的

语法：

type 接口名 interface {

​	method1(参数列表) 返回值列表

​	method2(参数列表) 返回值列表

}

注：

1. 接口体中的所有方法都没有方法体，即接口的方法都是没有实现的方法；
2. golang中的接口，不需要显式的实现，只要有一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口;
3. golang接口中不能有任何变量
4. 空接口interface{} 没有任何方法，所有类型都是实现了空接口，即我们可以把任何一个变量赋给空接口





类型断言 assert

由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言

## 



## 十二、指针

指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值

var num int = 5

var ptr *int = &num





## 面向对象

+ 特性
  + 封装
  + 继承
  + 多态



## 单元测试

![image-20220329151211396](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220329151211396.png)



测试单个文件，一定要带上被测试的源文件

go test -v cal_test.go cal.go

测试单个方法

go test -v -run TestAddUpper



## 文件读写

文件信息

使用os.Stat(fileName) 查看关于文件的各种信息

> ```go
> func Stat(name string) (fi FileInfo, err error)
> // Stat返回一个描述name指定的文件对象的FileInfo
> ```
>
> ```
> type FileInfo interface {
>     Name() string       // 文件的名字（不含扩展名）
>     Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
>     Mode() FileMode     // 文件的模式位
>     ModTime() time.Time // 文件的修改时间
>     IsDir() bool        // 等价于Mode().IsDir()
>     Sys() interface{}   // 底层数据来源（可以返回nil）
> }
> ```



读文件

写文件

序列化 及 反序列化



## 并发协程

同步锁：不同goroutine对公共资源进行访问

管道：不同goroutine之间交换信息



## 反射

动态代理及框架中大量使用反射机制







## 泛型--Go1.18新特性

泛型为Go添加了三个新的重要内容

1. 面向函数和类型的“类型形参”(type parameters)
2. 将接口类型定义为类型集合，包括没有方法的接口类型
3. 类型推断：在大多数情况下，在调用泛型函数时可省略“类型实参”



+ 类型参数 Type Parameters

现在函数和类型都具有类型形参” (type parameters)，类型形参列表看起来就是一个普通的参数列表，除了它使用的是方括号而不是小括号。

先从浮点值的基本非泛型 Min 函数开始

```go
func Min(x ,y float){
    if x > y {
        return x
    }
    return y
}
```

通过添加类型形参列表来使这个函数泛型化 -- 使其适用于不同的类型。在此实例中，添加了一个带有单个类型形参T的类型参数列表，并替换了float64

```go
import "golang.org/x/exp/constraints"

func GMin[T constraints.Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}
// 然后就可以使用类型实参调用此函数
// 向GMin提供类型参数，在这种情况下int称为实例化。实例化份两步进行，首先，编译器在泛型函数或泛型类型中用所有类型形参替换它们各自的类型实参，然后，编译器验证每个类型形参是否满足各自的约束。如果第二步失败，实例化就会失败并且程序无效。
x := GMin[int](2,3)


// 成功实例化后，即可产生非泛型函数，它可以像任何其他函数一样被调用。比如：
fmin := GMin[float64]
m := fmin(2.71, 3.14)

```



+ 类型推断

  此项功能是最复杂的变更，主要包括：

  + 函数参数类型推断
  + 约束类型推断

  虽然类型推断的工作原理细节很复杂， 但使用它并不复杂：类型推断要么成功，要么失败。如果它成功，类型实参可以被省略，调用泛型函数看起来与调用普通函数没有什么不同。如果类型推断失败，编译器将给出错误消息，在这种情况下，只需提供必要的类型实参





# 数据库

go语言中`database/sql`包提供了保证SQL 或类SQL数据库的泛用接口，并不提供具体的数据库驱动。使用`database/sql`包时必须注入(至少)一种数据库驱动

下载依赖

```shell
go get -u github.com/go-sql-driver/mysql
```

使用

1. sql.Open() 使用数据库驱动的名称 / 数据源名称 得到一个指向sql.DB的指针 

2. sql.DB 是用来操作数据库的，它代表了0个或者多个底层连接的池，这些连接由sql包来维护，sql包会自动创建和释放这些连接 

   > 注：Open()函数并不会连接数据库，甚至不会验证其参数。它只是把后续连接到数据库所必需的structs给设置好了
   >
   > 而真正的连接是被需要的时候才进行懒设置的
   >
   > sql.DB可以不进行关闭，它就是用来处理数据库的，而不是实际的连接，里面包含了数据库连接的池，而且会对此进行维护

```go
package main

import (
	"database/sql"

    /*
    正常做法是使用sql.Register()函数、数据库驱动名称和一个实现了driver.Driver 接口的struct，来注册数据库的驱动
    但是在包被引入的时候进行了自我注册
    */
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
    /*
    func Open(driverName, dataSourceName string) (*DB, error)
    Open打开一个driverName指定的数据库，dataSourceName指定数据源，一般至少包括数据库文件名和其他连接的必要条件
    */
    dsn := "root:wr19980211@tcp(127.0.0.1:3306)/mysql"
	db, err := sql.Open("mysql", dsn)
    
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
```



### CRUD

+ retrieve

  + 单行查询

    单行查询`db.QueryRow()`执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）

    ```go
    package main
    
    import (
    	"database/sql"
    	"fmt"
    
    	_ "github.com/go-sql-driver/mysql"
    )
    
    var db *sql.DB
    
    type User struct {
    	id   int
    	age  int
    	name string
    }
    
    func main() {
    	initDB()
    	queryDemo(1)
    }
    
    func initDB() {
    	dsn := "root:wr19980211@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True"
    	var err error
    	db, err = sql.Open("mysql", dsn)
    	if err != nil {
    		fmt.Println("sql.Open err =", err)
    		return
    	}
    	err = db.Ping()
    	if err != nil {
    		fmt.Println("db.Ping err =", err)
    		return
    	}
    }
    
    func queryDemo(id int) {
    	str := "select id,name ,age from user where id = ? "
    	var user User
    	err := db.QueryRow(str, id).Scan(&user.id, &user.name, &user.age)
    	if err != nil {
    		fmt.Println("db.QueryRow err = ", err)
    	}
    	fmt.Println(user)
    }
    
    ```

  + 多行查询

    多行查询`db.Query()`执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数

    ```go
    package test
    
    import (
    	"database/sql"
    	"fmt"
    	"testing"
    
    	_ "github.com/go-sql-driver/mysql"
    )
    
    type user struct {
    	id   int
    	age  int
    	name string
    }
    
    // 定义一个全局对象db
    var db *sql.DB
    
    // 定义一个初始化数据库的函数
    func initDB() (err error) {
    	// DSN:Data Source Name
    	dsn := "root:wr19980211@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True"
    	// 不会校验账号密码是否正确
    	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
    	db, err = sql.Open("mysql", dsn)
    	if err != nil {
    		return err
    	}
    	// 尝试与数据库建立连接（校验dsn是否正确）
    	err = db.Ping()
    	if err != nil {
    		return err
    	}
    	return nil
    }
    
    // 查询多条数据示例
    func queryRowAllDemo() {
    	sqlStr := "select id, name, age from user"
    	rows, err := db.Query(sqlStr)
    	if err != nil {
    		fmt.Printf("db.Query, err:%v\n", err)
    		return
    	}
    	for rows.Next() {
    		var u user
    		// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
    		err = rows.Scan(&u.id, &u.name, &u.age)
    		if err != nil {
    			fmt.Printf("scan failed, err:%v\n", err)
    			return
    		}
    		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
    	}
    }
    
    func TestMain(t *testing.T) {
    	err := initDB() // 调用输出化数据库的函数
    	if err != nil {
    		fmt.Printf("init db failed,err:%v\n", err)
    		return
    	}
    
    	queryRowAllDemo()
    }
    
    ```

+ create

+ update

+ delete

  > 更新，插入，删除操作都使用的是`Exec`方法

  ```go
  package test
  
  import (
  	"database/sql"
  	"fmt"
  	"testing"
  
  	_ "github.com/go-sql-driver/mysql"
  )
  
  type user struct {
  	id   int
  	age  int
  	name string
  }
  
  // 定义一个全局对象db
  var db *sql.DB
  
  // 定义一个初始化数据库的函数
  func initDB() (err error) {
  	// DSN:Data Source Name
  	dsn := "root:wr19980211@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True"
  	// 不会校验账号密码是否正确
  	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
  	db, err = sql.Open("mysql", dsn)
  	if err != nil {
  		return err
  	}
  	// 尝试与数据库建立连接（校验dsn是否正确）
  	err = db.Ping()
  	if err != nil {
  		return err
  	}
  	return nil
  }
  
  // create
  // 插入数据示例
  func insertDemo() {
  	str := "insert into user (name ,age ) values (?,?)"
  	res, err := db.Exec(str, "Candy", 18)
  	if err != nil {
  		fmt.Println("Insert data err = ", err)
  	}
  	id, err1 := res.LastInsertId() // 新插入数据的id
  	if err1 != nil {
  		fmt.Println("LastInsertId err = ", err1)
  	}
  	affect, err2 := res.RowsAffected() // 此次操作受影响的行数
  	if err1 != nil {
  		fmt.Println("RowsAffected err = ", err2)
  	}
  	fmt.Printf("result id=%v , affect rows num = %v \n", id, affect)
  }
  
  // update
  // 更新数据示例
  func updateDemo() {
  	str := "update user set age = ? where id = ?"
  	res, err := db.Exec(str, 100, 2)
  	if err != nil {
  		fmt.Println("update data err = ", err)
  	}
  
  	id, err1 := res.LastInsertId()
  	if err1 != nil {
  		fmt.Println("LastInsertId err = ", err1)
  	}
  	affect, err2 := res.RowsAffected()
  	if err1 != nil {
  		fmt.Println("RowsAffected err = ", err2)
  	}
  	fmt.Printf("result id=%v , affect rows num = %v \n", id, affect)
  }
  
  // delete
  // 删除数据示例
  func deleteDemo() {
  	str := "delete from user where id = ?"
  	res, err := db.Exec(str, 1)
  	if err != nil {
  		fmt.Println("delete data err = ", err)
  	}
  
  	id, err1 := res.LastInsertId()
  	if err1 != nil {
  		fmt.Println("LastInsertId err = ", err1)
  	}
  	affect, err2 := res.RowsAffected()
  	if err1 != nil {
  		fmt.Println("RowsAffected err = ", err2)
  	}
  	fmt.Printf("result id=%v , affect rows num = %v \n", id, affect)
  }
  
  func TestMain(t *testing.T) {
  	err := initDB() // 调用输出化数据库的函数
  	if err != nil {
  		fmt.Printf("init db failed,err:%v\n", err)
  		return
  	}
  
  	insertDemo()
  }
  
  ```

### 预处理

以上为普通SQL语句执行的过程：

1. 客户端对SQL语句进行占位符替换得到完整的SQL语句
2. 客户端发送完整的SQL语句到MySQL服务端
3. MySQL服务端执行完整的SQL语句并将结果返回给客户端

但以上结果存在执行重复SQL语句，多次编译，浪费服务器性能及成本；且带来了SQL注入问题



=》 预处理SQL语句：

1. 将SQL语句分成两部分，命令部分与数据部分
2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理
3. 然后把数据部分发送非MySQL服务端，MySQL服务端对SQL语句进行占位符i换
4. MySQL服务端执行完整的SQL语句并将结果返回给客户端



`database/sql`中使用下面的Prepare方法来实现预处理操作

`Prepare`方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令

+ 查询操作

  ```go
  package prepare
  
  import (
  	"database/sql"
  	"fmt"
  	"testing"
  
  	_ "github.com/go-sql-driver/mysql"
  )
  
  var db *sql.DB
  
  type User struct {
  	id   int
  	age  int
  	name string
  }
  
  // 获取数据库连接
  func initDB() {
  	dsn := "root:wr19980211@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True"
  	var err error
  	db, err = sql.Open("mysql", dsn)
  	if err != nil {
  		fmt.Println("sql.Open err", err)
  		return
  	}
  
  	err = db.Ping()
  	if err != nil {
  		fmt.Println("db.Ping err", err)
  		return
  	}
  }
  
  func TestPrepare(t *testing.T) {
  	initDB()
  	// prepareRowDemo()
  	prepareRowAllDemo()
  }
  
  // 使用预处理查询单行数据示例
  func prepareRowDemo() {
  	str := "select id ,name ,age from user where id = ?"
  	stmt, err := db.Prepare(str)
  	if err != nil {
  		fmt.Println("db.Prepare err = ", err)
  	}
  	defer stmt.Close()
  
  	var user User
  	err = stmt.QueryRow(2).Scan(&user.id, &user.name, &user.age)
  	if err != nil {
  		fmt.Println("stmt.QueryRow(1).Scan = ", err)
  	}
  	fmt.Printf("id:%d name:%s age:%d\n", user.id, user.name, user.age)
  }
  
  // 使用预处理查询多行数据示例
  func prepareRowAllDemo() {
  	str := "select id , name ,age from user"
  	stmt, err := db.Prepare(str)
  	if err != nil {
  		fmt.Println("db.Prepare err= ", err)
  	}
  
  	defer stmt.Close()
  	rows, err1 := stmt.Query()
  	defer rows.Close()
  	if err1 != nil {
  		fmt.Println("stmt.Query err1= ", err1)
  	}
  
  	for rows.Next() {
  		var user User
  		err = rows.Scan(&user.id, &user.name, &user.age)
  		if err != nil {
  			fmt.Println("stmt.QueryRow(1).Scan = ", err)
  		}
  		fmt.Printf("id:%d name:%s age:%d\n", user.id, user.name, user.age)
  	}
  
  }
  ```

+ 更新、删除、添加操作，类似

  ```go
  // 使用预处理插入数据示例
  func InsertDemo() {
  	str := "insert into user (id ,name ,age ) values (?,?,?)"
  	stmt, err := db.Prepare(str)
  	if err != nil {
  		fmt.Println("db.Prepare err= ", err)
  	}
  	defer stmt.Close()
  
  	res, err1 := stmt.Exec(1, "Jack~", 30)
  	id, err1 := res.LastInsertId() // 新插入数据的id
  	if err1 != nil {
  		fmt.Println("LastInsertId err = ", err1)
  	}
  	affect, err2 := res.RowsAffected() // 此次操作受影响的行数
  	if err1 != nil {
  		fmt.Println("RowsAffected err = ", err2)
  	}
  	fmt.Printf("result id=%v , affect rows num = %v \n", id, affect)
  }
  
  ```

  

> SQL注入问题
>
> 任何时候都不应该自己拼接SQL语句



### 事务

事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，同时这个完整的业务需要执行多次的DML(insert \ update \ delete) 语句共同联合完成

在Mysql中只有使用InnoDB数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行

+ 事务的ACID

  通常事务必须满足4个条件：原子性(Atomicity，或称为不可分割性)、一致性(Consistency)、隔离性(Isolation，又称独立性)、持久性(Durability)

  | 条件   | 解释                                                         |
  | ------ | ------------------------------------------------------------ |
  | 原子性 | 一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。 |
  | 一致性 | 在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。 |
  | 隔离性 | 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable） |
  | 持久性 | 事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。 |

+ 事务的方法

  + 开始事务

    ```go
    func (db *DB) Begin() (*Tx, error)
    ```

  + 提交事务

    ```go
    func (tx *Tx) Commit() error 
    ```

  + 回滚事务

    ```go
    func (tx *Tx) Rollback() error
    ```

    ```go
    // 事务操作示例
    func transactionDemo() {
    	tx, err := db.Begin() // 开启事务
    	if err != nil {
    		if tx != nil {
    			tx.Rollback() // 回滚
    		}
    		fmt.Printf("begin trans failed, err:%v\n", err)
    		return
    	}
    	sqlStr1 := "Update user set age=30 where id=?"
    	ret1, err := tx.Exec(sqlStr1, 2)
    	if err != nil {
    		tx.Rollback() // 回滚
    		fmt.Printf("exec sql1 failed, err:%v\n", err)
    		return
    	}
    	affRow1, err := ret1.RowsAffected()
    	if err != nil {
    		tx.Rollback() // 回滚
    		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
    		return
    	}
    
    	sqlStr2 := "Update user set age=40 where id=?"
    	ret2, err := tx.Exec(sqlStr2, 3)
    	if err != nil {
    		tx.Rollback() // 回滚
    		fmt.Printf("exec sql2 failed, err:%v\n", err)
    		return
    	}
    	affRow2, err := ret2.RowsAffected()
    	if err != nil {
    		tx.Rollback() // 回滚
    		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
    		return
    	}
    
    	fmt.Println(affRow1, affRow2)
    	if affRow1 == 1 && affRow2 == 1 {
    		fmt.Println("事务提交啦...")
    		tx.Commit() // 提交事务
    	} else {
    		tx.Rollback()
    		fmt.Println("事务回滚啦...")
    	}
    
    	fmt.Println("exec trans success!")
    }
    ```

    



# 标准库

### http/template

`html/template`包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。它提供了和`text/template`包相同的接口，Go语言中输出HTML的场景都应使用`html/template`这个包



+ 模板与渲染

  模板可以理解为事先定义号的HTML文档文件，模板渲染的作用机制可以简单理解为文本替换操作 -- 使用相应的数据去替换HTML文档中事先准备好的标记

+ 使用

  + 定义模板文件

    根据模板文件的语法规则去编写

  + 解析模板文件

    定义好模板文件之后，可以使用下面的常用方法去解析模板文件，得到模板对象

    ```go
    func (t *Template) Parse(src string) (*Template, error)
    func ParseFiles(filenames ...string) (*Template, error)
    func ParseGlob(pattern string) (*Template, error)
    ```

  + 模板渲染

    使用数据去填充模板

    ```go
    func (t *Template) Execute(wr io.Writer, data interface{}) error
    func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
    ```

+ 模板语法

  + {{ . }} 

    模板引擎都包含在 {{ 和 }} 中，其中{{ . }} 中的点表示当前对象

  + {{/*  这是模板语言中的注释 */}}

  + 变量：

  + 移除空格：{{-  移除模板内容左右两侧的空白符号 -}}

  + 条件判断

    ```html
    {{ if pipline }}
    	t1
    {{ end }}
    
    
    {{ if pipline }}
    	t1
    {{ else if pipline }}
    	t0
    {{ end }}
    ```

  + range

  + with

    相当于造了一个局部变量，可以节省代码。比如此代码中将.name 中的点替换成上面的 .m1

    ![image-20220330141228058](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220330141228058.png)

  + 预定义函数

    ```
    and
        函数返回它的第一个empty参数或者最后一个参数；
        就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
    or
        返回第一个非empty参数或者最后一个参数；
        亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
    not
        返回它的单个参数的布尔值的否定
    len
        返回它的参数的整数类型长度
    index
        执行结果为第一个参数以剩下的参数为索引/键指向的值；
        如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
    print
        即fmt.Sprint
    printf
        即fmt.Sprintf
    println
        即fmt.Sprintln
    html
        返回与其参数的文本表示形式等效的转义HTML。
        这个函数在html/template中不可用。
    urlquery
        以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
        这个函数在html/template中不可用。
    js
        返回与其参数的文本表示形式等效的转义JavaScript。
    call
        执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
        如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
        其中Y是函数类型的字段或者字典的值，或者其他类似情况；
        call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
        该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
        如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
    ```

  + 比较函数 ： 比较函数只适合于基本类型

    + eq  等于 `(只有eq可以接收2个及以上的参数，它会将第一个参数和其他参数依次比较)`
    + ne  不等于
    + lt  小于
    + le 小于等于
    + gt  大于
    + ge 大于等于

  + 自定义函数

    除了预定义函数外，你也可以使用自定义函数，在模板中跟预定义函数一样调用

    一定要在解析模板之前注册函数

    ![image-20220330143346460](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20220330143346460.png)

  + 嵌套template

  + block









# 其他

## 依赖注入

+ 简介

  依赖注入是一种编程模式。比较适合面向对象编程，在函数式编程中则不需要。golang是一门支持多范式编程的语言，所以在使用面向对象的大型项目中，还是建议按照实际情况判断是否应该使用依赖注入模式

+ 依赖注入库

  + uber开发的dig
  + elliotchance开源的dingo
  + sarulabs开源的di
  + google开源的wire
  + facebook开源的inject等

+ 基本用法

  1. 创建容器

     ```go
     container := dig.New()
     ```

     容器用来管理依赖

  2. 注入依赖

     调用容器的Provide方法，传入一个工厂函数，容器会自动调用该工厂函数创建依赖，并保存到container

     ```go
     type DBClient struct {}
     
     func NewDBClient() {
         return *DBClient{}
     }
     
     func InitDB() *DBClient {
         return NewDBClient()
     }
     
     container.Provide(InitDB)
     ```

     注入的依赖会被dig所管理，每种类型的对象只会被创建一次，可以理解为单例。如果再注入同一类型的依赖，工厂函数则不会被执行

     ```go
     type DBClient struct {}
     
     func NewDBClient() {
         return *DBClient{}
     }
     
     func InitDB() *DBClient {
         return NewDBClient()
     }
     
     func InitDB2() *DBClient {
         return NewDBClient()
     }
     
     // 相同类型的对象只会被创建一次
     container.Provide(InitDB)
     container.Provide(InitDB2)// 不会执行
     ```

     

  3. 使用依赖

     如果需要使用依赖，使用container的Invoke方法。并在函数的形式参数中定义参数，container会自动把单例注入

     ```go
     func UseOption( db *DBClient ){
         
     }
     
     container.Invoke(UseOption)
     ```

     