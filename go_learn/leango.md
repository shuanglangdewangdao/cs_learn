# 变量声明

## 完整声明

    var a int

有默认值

## 声明并初始化

    var a int = 100

## 简化声明

    var a = 100

## 最简声明

    a:=100

## 多重变量声明

    var a,b,c,=1,"ss",true

    var (
        a int = 10
        b string="ss"
    )


# 函数

    // r1 r2的默认值为0
    func foo1(a,b,c,e int)(r1 int,r2 string){

    }

# import与init函数

![](init函数.png)

首字母大写表示接口对外开放

现在要完整路径才能正确导入包

所有包都要init函数

## 匿名别名

只使用init方法

    import _ "baoming"

## 别名

    import aa "baoming"

## 全部导入

    import . "baoming"

# defer

程序最后执行的语句

    defer fmt.Println("1")
    defer fmt.Println("2")

    输出 2 1

执行过程符合程序栈的习惯

## defer 与 return

return 早于 defer 执行

return有两个动作，先进行值返回，再执行defer，最后执行RET退出函数

# slice

## 数组

    var array [10] int

### 赋值

#### 索引赋值

    array[i]=k

#### 直接赋值

    array := [10] int {1,2,3,4}

### 在函数中的使用

本身长度不同数据类型不同，只能传同类型的参数

值传递

## 动态数组

    array：= []int{}

### 在函数中的使用

传指针

## Slice

    var numbers = make([]int,3,5)

    len=3,cap=5

![](Slice.png)

### 追加

    numbers = append(numbers,1)

    容量不满时，值增加len不增加cap

    容量满时重新开辟2*cap的容量复制进去,这里的cap时当前的cap

### 新切片

    s1:=ss[0:2]

s1首地址与ss相同

    s1:=make([]int,3)

    copy(s1,ss)


# map

是一种哈希表(使用过程中为指针的副本)

## 声明与赋值
第一种

    var myMap map[string]string

    myMap=make(map[string]string,10)

第二种

    myMap:=make(map[string]string)


第三种

    myMap:=make(map[string]string){
        "one": "php"
    }


## 使用

添加

    mayMap["two"]="python"

遍历

    for key,value:=range myMap

删除

    delete(myMap,"one")

修改

    myMap["one"]="opo"


# 结构体

    type Oop struct {

    }

    var oop1 Oop

fun(oop1 Oop):是值传递


## 类

通过结构体绑定方法

    func (t *Oop) show{

    }

不建议使用this和self

    oop1.show()

### 注意

类名首字母大写表示可以在其他包访问

类内属性首字母大写表示可以在类外访问，否则只能在类内访问

方法名首字母大写可以在包外访问，首字母小写在包内访问


### 继承

    type NewOop struct {
        Oop
        level int
    }

#### 使用

    var oop2 NewOop

### 多态

父类使用子类方法

#### interface
interface 本质上是一个指向具体类型和方法的指针
    type AnimalIF interface {
        Sleep()
        GetColor()
        GetType()
    }

##### 使用

只要实现上述全部方法就使用了这个接口

# interface

interface{}:万能类型可以应用任意类型

## 判断是什么类型

类型断言
    arg interface{}

    value,ok:=arg.(string)

## 变量

type：static type(int,string),concrete type(interface指向的类型)二选一

value:

### 断言
相同静态类型可以通过断言转换

    type Reader interface{
        ReadeBook()
    }

    type Writer interface{
        WriteBook()
    }

    Type Book struct{

    }

    func (t *Book) ReadeBook(){

    }
    func (t *Book) ReadeBook(){

    }

    func main(){
        b:=&book{}

        var r Reader
        r=b

        r.ReadeBook()

        w=r.(Writer)
    }

# 反射

reflect包有两个函数

func ValueOf(i interface) Value {...}:返回数据中的值，若空则返回0
func TypeOf(i interface) Type {...}:返回数据中的类型，若空则返回nil，本身是一个断言


# 结构体标签

## 标签定义

    type resume struct{
        Name string 'info:"name" doc:"我的名字"'
        Sex string 'info:"sex"'
    }

## 获取

使用反射获取标签

## 使用

josn文件的编解码

orm映射关系

