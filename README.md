# go-object

本项目主要目的是介绍一种在Golang语法上实现面向对象代码风格，项目的代码并非重点。为了方便描述，本项目介绍的面向对象代码风格称为**go-object风格**。

> [!NOTE]
>
> **go-object风格的**经典实现例子：https://github.com/jishaocong0910/go-sql-parser

# 类

**go-object风格**中类只有两种：抽象类和非抽象类，抽象类只能被继承不能单独创建，非抽象类则相反。

## 抽象类

***抽象类***由***抽象接口***和***抽象成员***组成，***抽象接口***是一个接口，***抽象成员***是一个结构体，它们是一个整体相辅相成。

***抽象接口***是**抽象类**的代表，多态关系的父类表现形式（**go-object风格**的多态关系实际为Golang接口的实现关系），还用于声明子类必须实现的方法和可重写方法。声明方式：名称格式为`I_<类名>`，必须声明一个用于转化为***抽象成员***的方法，称为***父类成员方法***，名称格式为：`M_<类名>_`，无参数，返回***抽象成员***指针。

***抽象成员***用于存放**抽象类**的成员变量和方法，即使没有成员变量或方法，也需要定义它，方便以后扩展。声明方式：名称格式为`M_<类名>`（与***父类成员方法***相差一个下划线，原因是避免编译错误），必须声明一个***抽象接口***的字段`I`，称为***子类对象字段***，还需实现***父类成员方法***，实现逻辑为将本身指针返回。

*抽象类声明*

```go
// 抽象接口
type I_Biology interface {
    // 父类成员方法
    M_Biology_() *M_Biology
}

// 抽象成员
type M_Biology struct {
    // 子类对象字段
    I I_Biology
}

// 实现父类成员方法，返回本身指针
func (this *M_Biology) M_Biology_() *M_Biology {
    return this
}
```

## 非抽象类

***非抽象类***非常简单，它形式仅是一个简单的结构体。

# 继承

**go-object风格**的继承遵循以下规则。

1. ***抽象类***继承***抽象类***，通过***抽象接口***内嵌另一个***抽象接口***方式实现，可多继承。
2. ***非抽象类***继承***抽象类***，通过内嵌***抽象成员***指针方式方式实现，可多继承。
3. <font color=Red>***非抽象类***继承的***抽象类***具有父类时，***非抽象类***必须内嵌这些父类的***抽象成员***。</font>

*抽象类继承抽象类*

```go
type I_Animal interface {
    I_Biology // 内嵌Biology的抽象接口
    M_Animal_() *M_Animal
}

type M_Animal struct {
    I I_Animal
}

func (this *M_Animal) M_Animal_() *M_Animal {
    return this
}
```

*非抽象类继承顶级抽象类*

```go
type Cat struct {
    *M_Biology // 内嵌Biology的抽象结构
}
```

*非抽象类继承非顶级抽象类*

```go
// 错误方式（Animal不是顶级抽象类，Animal父类的抽象结构也需要内嵌）
type Dog struct {
    *M_Animal // 只内嵌Animal的抽象结构是不合法的
}
```

```go
// 正确方式
type Dog struct {
    *M_Biology // 内嵌Biology的抽象结构
    *M_Animal  // 内嵌Animal的抽象结构
}
```

# 构造器

所有类都必须具有构造器，构造器是一个函数，***抽象类***和***非抽象类***规则不同。

## 抽象类构造器

声明规则：名称格式为`Extend<类名>`，创建并返回***抽象成员***指针，参数列表首个参数为***抽象接口***，用于初始化***子类对象***，其他参数根据实际需要自定。

*完整的抽象类声明*

```go
type I_Biology interface {
    M_Biology_() *M_Biology
}

type M_Biology struct {
    I I_Biology
}

func (this *M_Biology) M_Biology_() *M_Biology {
    return this
}

func ExtendBiology(i I_Biology) *M_Biology {
    return &M_Biology{I: i}
}
```

## 非抽象类构造器

声明规则：名称格式为`New<类名>`，无必要构造参数，根据实际情况自定，返回类的结构体指针。构造基本逻辑是，先创建本类的对象（本结构体指针），再调用抽象类的构造函数，首个参数传入本类对象，其他参数根据实际需要传值，创建的变量用来初始化自身内嵌的***抽象成员***。

*完整的非抽象类声明*

```go
type Dog struct {
    *M_Biology
    *M_Animal
}

// 构造器
func NewDog() *Dog {
    // 先创建本类对象
    c := &Helicopter{}
    // 使用对应的抽象类的构造器，初始化所有内嵌抽象成员。
    c.M_Vehicle = ExtendVehicle(c)
    c.M_Aircraft = ExtendAircraft(c)
    return c
}
```

# 类的其他规则

- 类的名称、字段、方法和构造器等等，都没有强制是否暴露，可根据实际需要选择首字母的大小写。例如抽象类只在包内使用，抽象接口和抽象结构也命名为`i_<类名>`和`m_<类名>`。
- 方法接受者使用指针，统一名称为`this`。
- 使用单独文件声明类，以区分普通Go代码，文件名格式为，抽象类：`<类名>.aclass.go`，非抽象类`<类名>.class.go`。


# 中间类

将即可以被继承，又可以实例化的类，在**go-object**风格中称为***中间类***。尽管**go-object**风格没有***中间类***的声明方式，但还是能变通实现这样的类。

## 抽象类实例化



## 多继承二义性

多继承一般都会产生二义性问题，其中一个常见的场景：所有类的方法都不相同，但子类继承的多个父类又继承了相同父类，导致出现二义性问题，例如B和C继承了A，D继承B和C，D调用A的方法时出现二义性问题。

*C++多继承二义性*

```c++
#include <iostream>
using namespace std;

class A {
public:
    void foo()
    {
        cout << "A::foo()" << endl;
    }
};

class B : public A {
public:
    void bar()
    {
        cout << "B::bar()" << endl;
    }
};

class C : public A {
public:
    void baz()
    {
        cout << "C::baz()" << endl;
    }
};

class D : public B, public C {
public:
    void qux()
    {
        cout << "D::qux()" << endl;
    }
};

int main()
{
    D d;
    d.qux();
    d.bar();
    d.baz();
    d.foo(); // error: request for member ‘foo’ is ambiguous
    return 0;
}
```

**go-object风格**解决了这种场景的二义性问题，因为在它的规则中，***抽象类***之间的继承只是将***抽象接口***进行内嵌，***抽象成员***则没有，所以并不会拷贝父类的成员方法，因此不会产生二义性，并且由于继承规则的最后一条，子类必须继承每一级父类，也就是内嵌每一级的***抽象成员***，所以保证子类会拥有所有父类的成员方法。现在我们用**go-object风格**来改写上面的C++例子。

*go-object风格改写*

```go
package main

import "fmt"

// 类A
type I_A interface {
    M_A_() *M_A
}

type M_A struct {
    I I_A
}

func (this *M_A) M_A_() *M_A {
    return this
}

func (this *M_A) Foo() {
    fmt.Println("A::foo()")
}

func ExtendA(i I_A) *M_A {
  return &M_A{I: i}
}

// 类B
type I_B interface {
    I_A // 继承类A
    M_B_() *M_B
}

type M_B struct {
    I I_B
}

func (this *M_B) M_B_() *M_B {
    return this
}

func (this *M_B) Bar() {
    fmt.Println("B::bar()")
}

func ExtendB(i I_B) *M_B {
  return &M_B{I: i}
}

// 类C
type I_C interface {
    I_A // 继承类A
    M_C_() *M_C
}

type M_C struct {
    I I_C
}

func (this *M_C) M_C_() *M_C {
    return this
}

func (this *M_C) Baz() {
    fmt.Println("C::baz()")
}

func ExtendC(i I_C) *M_C {
  return &M_C{I: i}
}

// 类D
type D struct {
    *M_A // 继承类A
    *M_B // 继承类B
    *M_C // 继承类C
}

func (this *D) Qux() {
    fmt.Println("D::qux()")
}

func NewD() *D {
  d := &D{}
  d.M_A = ExtendA(d)
  d.M_B = ExtendB(d)
  d.M_C = ExtendC(d)
  return d
}

func main() {
    // 以下代码正常运行
    d := NewD()
    d.Qux()
    d.Bar()
    d.Baz()
    d.Foo()
}
```

# 编译器的利用

**go-object风格**是一种自我约定的代码风格，没有编译器的天然约束，但它的一些规则具有Golang的语法约束。

## 子类必须内嵌所有父类的抽象结构

即继承规则的最后一条。

*错误写法*
```go
type Helicopter struct {
    *M_Aircraft // 只内嵌直接上级父类的成员结构
}

func NewHelicopter() *Helicopter {
    c := &Helicopter{}
    c.M_Aircraft = ExtendAircraft(c) // some methods are missing: M_Vehicle() *M_Vehicle
    return c
}
```

*正确写法*
```go
type Helicopter struct {
    *M_Vehicle  
    *M_Aircraft
}

func NewHelicopter() *Helicopter {
    c := &Helicopter{}
    c.M_Vehicle = ExtendVehicle(c)
    c.M_Aircraft = ExtendAircraft(c)
    return c
}
```

## 子类必须实现抽象方法

*在抽象接口中增加方法。*

```go
type I_Vehicle interface {
    M_Vehicle_() *M_Vehicle
    Takeoff() string // 增加此方法
}
```
*子类的构造器将会出现编译错误，显示未实现方法。*

```go
func NewHelicopter() *Helicopter {
    c := &Helicopter{}
    c.M_Vehicle = ExtendVehicle(c) // some methods are missing: Takeoff() string
    c.M_Aircraft = ExtendAircraft(c) // some methods are missing: Takeoff() string
    return c
}
```
遇到此错误时，子类实现指定方法即可修复。

# 多态

**go-object风格**的多态规则：

- 子类转化为父类表现形式为子类转化为***抽象接口***。
- 父类（***抽象接口***）使用Golang的类型断言转化为子类。

*多态转换*

```go
h := NewHelicopter()
var v I_Vehicle = h  // Helicopter转Vehicle
var a I_Aircraft = h // Helicopter转Aircraft
a = v.(I_Aircraft)   // Vehicle转Aircraft
h = v.(*Helicopter)  // Vehicle转Helicopter
v = a                // Aircraft转Vehicle
v = a.(*Helicopter)  // Aircraft转Helicopter
```

# 方法重写

方法的重写，即子类重写父类的方法，**go-object风格**的规则是，重写的方法应该也在***抽象接口***中声明。

*方法重写例子*
```go
package main

import "fmt"

type I_Car interface {
    M_Car_() *M_Car
    Speed() int // 声明一个和成员结构相同方法，表示此方法可重写
}

type M_Car struct {
    I I_Car
}

func (this *M_Car) M_Car_() *M_Car {
    return this
}

func (this *M_Car) Speed() int {
    return 100
}

func ExtendCar(i I_Car) *M_Car {
  return &M_Car{I: i}
}

// 重写父类方法的子类
type Jeep struct {
    *M_Car
}

func (this *Jeep) Speed() int {
    return 200
}

func NewJeep() *Jeep {
  j := &Jeep{}
  j.M_Car = ExtendCar(j)
  return j
}

// 没有重写父类方法的子类
type Trucks struct {
    *M_Car
}

func NewTrucks() *Trucks {
  t := &Trucks{}
  t.M_Car = ExtendCar(t)
  return t
}

func main() {
    j := NewJeep()
    t := NewTrucks()
    cars := []I_Car{j, t}

    // 通过抽象接口调用，将调用子类重写的方法，如果子类没有重写则调用父类方法。
    for _, c := range cars {
        fmt.Println(c.Speed())
    }
    // Output:
    // 200
    // 100

    // 通过成员结构调用，可直接调用父类实现的方法。
    for _, c := range cars {
        fmt.Println(c.M_Car_().Speed())
    }
    // Output:
    // 100
    // 100
}
```

# 枚举

本项目用**go-object风格**设计了枚举功能。声明枚举需要创建枚举值和枚举集合两个结构体，枚举值内嵌`*o.M_EnumValue`，枚举集合内嵌`*o.M_Enum`并指定泛型类型为枚举值的类型，再通过`o.NewEnum`函数创建枚举集合变量。

枚举集合的名称建议使用格式`_<枚举值名>`，因为实践中它基本不会暴露给包外。

枚举值和枚举集合都不需要构造函数，`o.NewEnum`函数已包装了统一逻辑。

判断枚举相等应比较`<枚举值>.ID() == <枚举值>.ID()`，不要直接比较`<枚举值> == <枚举值>`，switch语法也应使用`<枚举值>.ID()`作为条件。

建议使用单独文件声明枚举以区分普通Go代码，枚举值和枚举集合都在统一文件，文件名格式为`<枚举名>.enum.go`。

*o.M_EnumValue方法*

| 方法        | 说明                     |
|-----------|------------------------|
| Id        | 在枚举集合中的ID，值为在枚举集合的字段名。 |
| Undefined | 返回该枚举值是否存在。            |

*o.M_Enum方法*

| 方法             | 说明                                                                 |
|----------------|--------------------------------------------------------------------|
| Values         | 所有枚举值，一般用于自定义枚举查找。                                                 |
| Undefined      | 返回一个未定义的枚举值。                                                       |
| OfId           | 根据ID查找枚举值，ID为枚举字段名称。                                               |
| OfIdIgnoreCase | 根据ID查找枚举值，不区分大小写。                                                  |
| Is             | 判断目标枚举值是否等于源枚举值，多个目标枚举则只需满足其中一个，如果两个枚举值`Undefined`方法都返回true，则认为相等。 |
| Not            | 与`Is`方法相反。                                                         |

*示例代码*

```go
package main

import (
    "fmt"

    o "github.com/jishaocong0910/go-object"
)

// 枚举值
type DbType struct {
    *o.M_EnumValue // 内嵌*o.M_EnumValue
    Name           string
}

// 枚举集合
type _DbType struct {
    *o.M_Enum[DbType]                  // 内嵌*o.M_Enum
    MYSQL, ORACLE, SQLSERVER, POSTFRES DbType
}

// 自定义查找方法（接受者不需要指针，名称不需要为this）
func (d _DbType) OfName(name string) (result DbType) {
    for _, t := range d.Values() {
        if t.Name == name {
            result = t
            break
        }
    }
    return
}

// 创建枚举集合变量，并初始化每个枚举值
var DbTypes = o.NewEnum[DbType](_DbType{
    MYSQL:     DbType{Name: "MySQL"},
    ORACLE:    DbType{Name: "Oracle"},
    SQLSERVER: DbType{Name: "SQLserver"},
    POSTFRES:  DbType{Name: "PostgreSQL"},
})

func main() {
    d := DbTypes.MYSQL
    fmt.Println(d.ID())
    fmt.Println(DbTypes.OfIdIgnoreCase("sqlserver").ID())
    // Output:
    // MYSQL
    // SQLSERVER

    d2 := DbTypes.OfName("MySQL")
    fmt.Println(d2.Undefined())
    fmt.Println(d.ID() == d2.ID())
    // Output:
    // false
    // true

    d3 := DbTypes.OfName("abc")
    fmt.Println(d3.Undefined())
    // Output:
    // true

    switch d.ID() {
    case DbTypes.MYSQL.ID():
        fmt.Println("mysql")
    case DbTypes.ORACLE.ID():
        fmt.Println("oracle")
    case DbTypes.SQLSERVER.ID():
        fmt.Println("sqlserver")
    case DbTypes.POSTFRES.ID():
        fmt.Println("postgresql")
    }
}
```

> [!NOTE]
>
> 枚举功能虽然采用**go-object风格**实现，但是对一些规则做了优化：内嵌枚举值和枚举集合抽象类的结构体，不被视为“类”，它们以变量值形式存在而非指针，因此不推荐方法接受者使用指针、名称使用`this`。优化原因有以下几点：
>
> - 枚举信息应该是只读的，所以用值表示而非指针。
> - 枚举集合多用来访问其字段（枚举值）而非传值，值类型比指针访问字段更快。

# 判断NULL

**go-object风格**使用Golang的接口表示父类，而Golang有著名的`nil ≠ nil`问题（nil指针赋值给接口，接口≠nil），本项目封装了一个函数用于判断抽象类是否为nil。

*判断NULL示例*

```go
package main

import (
    "fmt"

    o "github.com/jishaocong0910/go-object"
)

func main() {
    var a any
    var i *int
    fmt.Println(a == nil)
    fmt.Println(o.IsNull(a))
    fmt.Println(i == nil)
    // Output:
    // true
    // true
    // true
    
    a = i
    fmt.Println(a == nil)
    fmt.Println(o.IsNull(a))
    // Output:
    // false
    // true
}
```