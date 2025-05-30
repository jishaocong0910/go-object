# go-object

本项目主要目的是介绍一种在Golang语法上实现面向对象代码风格，项目的代码并非重点。为了方便描述，本项目介绍的面向对象代码风格称为**go-object风格**。

> [!NOTE]
>
> **go-object风格的**经典实现例子：https://github.com/jishaocong0910/go-sql-parser

# 类

**go-object风格**中类只有两种：抽象类和非抽象类，抽象类只能被继承不能单独创建，非抽象类则相反。

## 抽象类

***抽象类***由***抽象接口***和***抽象结构***组成，***抽象接口***是一个接口，***抽象结构***是一个结构体，它们是一个整体相辅相成。

***抽象接口***是**抽象类**的代表，多态关系的父类表现形式，还用于声明子类必须实现的方法和可重写方法。***抽象接口***的声明规则：名称格式为`I_<类名>`，必须声明一个用于转化为***抽象结构***的方法，称为***父类成员方法***，名称格式为：`M_<类名>_`，无参数，返回***抽象结构***指针。

***抽象结构***用于存放**抽象类**的成员字段和方法，即使没有成员字段或方法，也需要定义它，方便以后扩展。***抽象结构***的声明规则：名称格式为`M_<类名>`（与***父类成员方法***相差一个下划线，原因是避免编译错误），必须声明一个***抽象接口***的字段`I`，称为***子类对象字段***，还需实现***父类成员方法***，实现逻辑为将本身指针返回。

*Biology抽象类*

```go
// biology.aclass.go

// 抽象接口
type I_Biology interface {
    // 父类成员方法
    M_Biology_() *M_Biology
}

// 抽象结构
type M_Biology struct {
    // 子类对象字段
    I I_Biology
    // 其他成员字段
    Name string
}

// 实现父类成员方法，返回本身指针
func (this *M_Biology) M_Biology_() *M_Biology {
    return this
}

// 其他成员方法
func (this *M_Biology) Breathe() {
    fmt.Println("I'm breathing.")
}
```

## 非抽象类

***非抽象类***非常简单，它形式仅是一个简单的结构体。

# 继承

**go-object风格**的继承遵循以下规则。

1. ***抽象类***继承***抽象类***，通过***抽象接口***内嵌另一个***抽象接口***方式实现，可多继承。
2. ***非抽象类***继承***抽象类***，通过内嵌***抽象结构***指针方式方式实现，可多继承。
3. <font color=Red>***非抽象类***继承的***抽象类***具有父类时，***非抽象类***必须内嵌这些父类的***抽象结构***。</font>

*Animal继承Biology（抽象类继承抽象类）*

```go
// animal.aclass.go

type I_Animal interface {
    // 内嵌Biology的抽象接口
    I_Biology
    // 父类成员方法
    M_Animal_() *M_Animal
    // 声明子类须实现的方法
    Eat() bool
}

type M_Animal struct {
    I I_Animal
}

func (this *M_Animal) M_Animal_() *M_Animal {
    return this
}
```

*Cat继承Biology（非抽象类继承顶级抽象类）*

```go
// cat.class.go

type Cat struct {
    *M_Biology // 内嵌Biology的抽象结构
}
```

*Bird继承Animal（非抽象类继承非顶级抽象类）*

```go
// bird.class.go

type Bird struct {
    *M_Biology // 内嵌Biology的抽象结构
    *M_Animal  // 内嵌Animal的抽象结构 
    canSwin bool
}

// 实现父类的方法
func (this *Bird) Eat() {
    fmt.Println("I'm eating")
}

func (this *Bird) Swim() {
    if this.canSwim {
        fmt.Println("I'm swimming.")
    } else {
        fmt.Println("I can't swim.")
    }
}
```

> [!CAUTION]
> 
> 以下方式是不合法的，Animal不是顶级父类，其继承的父类的抽象结构也需要内嵌。
> 
> ```go
> type Bird struct {
>   *M_Animal // 只内嵌Animal的抽象结构是不够的
> }
> ```
> 

# 构造器

所有类都必须具有构造器，构造器是一个函数，***抽象类***和***非抽象类***规则不同。

## 抽象类构造器

***抽象类***构造器声明规则：名称格式为`Extend<类名>`，创建并返回***抽象结构***指针，参数列表首个参数为***抽象接口***，用于初始化***子类对象***，其他参数根据实际需要自定。

*Biology类完整声明*

```go
// biology.aclass.go

type I_Biology interface {
    M_Biology_() *M_Biology
}

type M_Biology struct {
    I I_Biology
    Name string
}

func (this *M_Biology) M_Biology_() *M_Biology {
    return this
}

func (this *M_Biology) Breathe() {
    return "I'm breathing"
}

// 构造器，首个参数为抽象接口，其他参数根据实际需要补充
func ExtendBiology(i I_Biology, name string) *M_Biology {
    return &M_Biology{I: i, Name: name}
}
```

## 非抽象类构造器

***非抽象类***构造器声明规则：名称格式为`New<类名>`，无必要构造参数，根据实际需要自定，返回类的结构体指针。构造基本逻辑是，先创建本类的对象（本结构体指针），再调用抽象类的构造函数，首个参数传入本类对象，其他参数根据实际需要传值，创建的变量用来初始化自身内嵌的***抽象结构***。

*Bird类完整声明*

```go
// bird.class.go

type Bird struct {
    *M_Biology
    *M_Animal
    canSwin bool
}

func (this *Bird) Eat() {
    fmt.Println("I'm eating")
}

func (this *Bird) Swim() {
    if this.canSwim {
        fmt.Println("I'm swimming.")
    } else {
        fmt.Println("I can't swim.")
    }
}

// 构造器，无必要参数，根据实际需要自定
func NewBird(name string, canSwim bool) *Bird {
    // 先创建本类对象
    b := &Bird{canSwim: canSwim}
    // 初始化所有内嵌抽象结构。
    b.M_Biology = ExtendBiology(b, name)
    b.M_Animal = ExtendAnimal(b)
    return b
}
```

# 类声明的其他规则

- 类的名称、字段、方法和构造器等等，都没有强制是否暴露，可根据实际需要选择首字母的大小写。例如抽象类只在包内使用，抽象接口和抽象结构可命名为`i_<类名>`和`m_<类名>`。
- 方法接受者使用指针，统一名称为`this`。
- 使用单独文件声明类，以区分普通Go代码，文件名格式为，抽象类：`<类名>.aclass.go`，非抽象类：`<类名>.class.go`。

# 多继承二义性

多继承一般都会产生二义性问题，常见的场景：子类继承的多个父类又继承了相同父类，导致出现二义性问题。例如B和C继承了A，D继承B和C，D调用A的方法时出现二义性问题。

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

**go-object风格**消除了这种场景的二义性问题，因为在它的规则中，***抽象类***之间的继承只内嵌***抽象接口***，没有内嵌***抽象结构***，也就不会拷贝父类的成员字段和方法，所以不会产生二义性。***抽象结构***是在子类中进行内嵌的，根据继承规则，子类内嵌了每一级父类，所以保证子类会拥有所有父类的成员字段和方法。

现在我们用**go-object风格**来改写上面的C++例子。

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

# 中间类

即可以被继承，又可以实例化的类，在**go-object**风格中称为***中间类***。尽管**go-object**风格没有***中间类***的声明方式，但还是能变通实现。

## 使抽象类实例化

为***抽象类***增加一个相同类名的子类，并且没有其他成员字段和方法，那么这个子类相当于可实例化的***抽象类***。从类名上看，***抽象类***为`I_<类名>`和`M_<类名>`，子类为`<类名>`，从文件名上看，***抽象类***为`<类名>.aclass.go`，子类为`<类名>.class.go`，命名都不冲突，并且容易看出两者之间的关系。

*Biology类改为可实例化*
```go
// biology.class.go

type Biology struct {
    *M_Biology
}

func NewBiology(name string) *Biology {
    b := &Biology{}
    b.M_Biology = ExtendBiology(b, name)
}
```

## 使非抽象类可继承

和抽象类实例化原理相同，将***非抽象类***改造为***抽象类***，再增加可实例化它的子类。

*Bird类改为可继承*

```go
// bird.aclass.go

type I_Bird interface {
    I_Animal
    M_Bird_() *M_Bird
}

type M_Bird struct {
    I I_Animal
    canSwim bool
}

func (this *M_Bird) M_Bird_() *M_Bird {
    return this
}

func (this *Bird) Eat() {
    fmt.Println("I'm eating")
}

func (this *Bird) Swim() {
    if this.canSwim {
        fmt.Println("I'm swimming.")
    } else {
        fmt.Println("I can't swim.")
    }
}

func ExtendBird(i I_Bird, canSwim bool) *M_Bird {
    return &M_Bird{I: i, canSwim: canSwim}
}
```

```go
// bird.class.go

type Bird struct {
    *M_Biology
    *M_Animal
    *M_Bird
}

func NewBird(name string, canSwim bool) *Bird {
    b := &Bird{}
    b.M_Biology = ExtendBiology(b, name)
    b.M_Animal = ExtendAnimal(b)
    b.M_Bird = ExtendBird(b, canSwim)
    return b
}
```

# 编译器的支持

**go-object风格**是一种自我约定的代码风格，编译器无法检查每一条规则，但一些重要规则还是能得到编译器的支持。

## 子类必须内嵌每一级父类的抽象结构

即继承规则的最后一条。

*Dog只继承Animal将编译错误*

```go
// dog.class.go

type Dog struct {
    *M_Animal
}

func (this *Dog) Eat() {
    fmt.Println("I'm eating.")
}

func NewDog() *Dog {
    d := &Dog{}
    d.M_Animal = ExtendAnimal(d) // 编译错误：some methods are missing: M_Biology_() *M_Biology
    return d
}
```

## 子类必须实现抽象方法

*Rabbit未实现Animal的方法将编译错误*

```go
// rabbit.class.go

type Rabbit struct {
    *M_Biology
    *M_Animal
}

func NewRabbit(name string) *Rabbit {
    r := &Rabbit{}
    r.M_Biology = ExtendBiology(r, name)
    r.M_Animal = ExtendAnimal(r) // 编译错误：some methods are missing: Eat()
    return r
}
```

# 多态

**go-object风格**的多态规则：

- 子类转化为父类表现形式为子类转化为***抽象接口***。
- 父类（***抽象接口***）使用Golang的类型断言转化为子类。

*多态转换*

```go
bd := NewBird()
var b I_Biology = bd  // Bird转Biology
var a I_Animal = bd   // Bird转Animal
a = b.(I_Animal)      // Biology转Animal
bd = b.(*Bird)        // Biology转Bird
b = a                 // Animal转Biology
bd = a.(*Bird)        // Animal转Bird
```

# 方法重写

方法的重写规则：***抽象类***可重写的方法在***抽象接口***和***抽象结构***都需要声明。

*Plant类声明可重写方法*

```go
// plant.aclass.go

type I_Plant interface {
    M_Plant_() *Plant
    Say() // 声明抽象结构的方法，使其可重写
}

type M_Plant struct {
    I I_Plant
}

func (this *M_Plant) M_Plant_() *M_Plant {
    return this
}

// 在抽象接口也声明此方法，使其可重写
func (this *M_Plant) Say() {
    fmt.Println("I'm a plant.")
}

func ExtendPlant(i I_Plant) *M_Plant {
    return &M_Bird{I: i}
}
```

*Tree类*

```go
// tree.class.go

type Tree struct {
    *M_Plant
}

func NewTree() *Tree {
    t := &Tree{}
    t.M_Plant = ExtendPlant(t)
    return t
} 

```

*Flower类*

```go
// flower.class.go

type Flower struct {
    *M_Plant
}

func (this *M_Plant) Say() {
    fmt.Println("I'm a flower.")
}

func NewFlower() *Flower {
    f := &Flower{}
    f.M_Plant = ExtendPlant(f)
    return f
} 

```

*重写方法的调用*

```go
func main() {
    t := NewTree()
    f := NewFlower()
    ps := []I_Plant{t, f}

    // 通过抽象接口调用，将调用子类重写的方法，如果子类没有重写则调用父类方法。
    for _, p := range ps {
        fmt.Println(b.Say())
    }
    // Output:
    // I'm a plant.
    // I'm a flower.

    // 子类可通过抽象结构调用父类实现的方法。
	for _, p := range ps {
        fmt.Println(b.M_Plant_().Say())
    }
    // Output:
    // I'm plant.
    // I'm plant.
}
```

# 枚举

本项目用**go-object风格**设计了枚举功能。声明枚举需要创建枚举值和枚举集合两个结构体，枚举值内嵌`*o.M_EnumElem`，枚举集合内嵌`*o.M_Enum`并指定泛型类型为枚举值的类型，再通过`o.NewEnum`函数创建枚举集合变量。

枚举值名称建议为`<枚举名>_`，枚举集合名称建议为`_<枚举名>`。

枚举值和枚举集合都不需要构造函数，`o.NewEnum`函数已包装了统一逻辑。

判断枚举相等应比较`<枚举值>.ID() == <枚举值>.ID()`，不要直接比较`<枚举值> == <枚举值>`，switch语法也应使用`<枚举值>.ID()`作为条件。

建议使用单独文件声明枚举以区分普通Go代码，枚举值和枚举集合都在统一文件，文件名格式为`<枚举名>.enum.go`。

*o.M_EnumElem方法*

| 方法        | 说明                     |
|-----------|------------------------|
| ID        | 在枚举集合中的ID，值为在枚举集合的字段名。 |
| Undefined | 返回该枚举值是否未定义。           |

*o.M_Enum方法*

| 方法             | 说明                                                                 |
|----------------|--------------------------------------------------------------------|
| Elems           | 所有枚举值，一般用于自定义枚举查找。                                                 |
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
type DbType_ struct {
    *o.M_EnumElem // 内嵌*o.M_EnumElem
    Name         string
}

// 枚举集合
type _DbType struct {
    *o.M_Enum[DbType_]                 // 内嵌*o.M_Enum
    MYSQL, ORACLE, SQLSERVER, POSTFRES DbType_
}

// 自定义查找方法（接受者不需要指针，名称不需要为this）
func (d _DbType) OfName(name string) (result DbType_) {
    for _, t := range d.Elems() {
        if t.Name == name {
            result = t
            break
        }
    }
    return
}

// 创建枚举集合变量，并初始化每个枚举值
var DbTypes = o.NewEnum[DbType_](_DbType{
    MYSQL:     DbType_{Name: "MySQL"},
    ORACLE:    DbType_{Name: "Oracle"},
    SQLSERVER: DbType_{Name: "SQLserver"},
    POSTFRES:  DbType_{Name: "PostgreSQL"},
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
> - 枚举应该是只读的，所以用值表示而非指针。
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