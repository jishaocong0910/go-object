# go-object

本项目主要目的是介绍一种在Golang语法上实现面向对象代码风格，简称为**go-object风格**。

> [!NOTE]
>
> **go-object风格的**经典实现例子：https://github.com/jishaocong0910/go-sql-parser

# 类

**go-object风格**只有两种类：抽象类和非抽象类，抽象类只能被继承不能单独创建，非抽象类则相反。

## 抽象类

***抽象类***由***抽象接口***和***抽象结构***组成，***抽象接口***是一个接口，***抽象结构***是一个结构体，它们是一个整体相辅相成。

***抽象接口***是**抽象类**的代表，多态关系的父类表现形式，还用于声明子类必须实现的方法和可重写方法。***抽象接口***的声明规则：名称格式为`<类名>_`，必须声明一个用于转化为***抽象结构***的方法，称为***抽象结构转换方法***，名称与***抽象接口***名称相同，无参数，返回***抽象结构***指针。

***抽象结构***用于存放**抽象类**的成员字段和方法，即使没有成员字段或方法，也需要定义它，因为继承规则需要，也方便以后扩展。***抽象结构***的声明规则：名称格式为`<类名>__`（双下划线），必须声明一个***抽象接口***的字段`I`，称为***子类对象字段***，还需实现***抽象结构转换方法***，实现逻辑为将本身指针返回。

*Biology抽象类*

```go
// biology.aclass.go

// 抽象接口
type Biology_ interface {
    // 抽象结构转换方法
    Biology_() *Biology__
}

// 抽象结构
type Biology__ struct {
    // 子类对象字段
    I Biology_
    // 其他成员字段
    Name string
}

// 实现抽象结构转换方法，返回本身指针
func (this *Biology__) Biology_() *Biology__ {
    return this
}

// 其他成员方法
func (this *Biology__) Breathe() {
    fmt.Println("I'm breathing.")
}
```

## 非抽象类

***非抽象类***非常简单，它形式仅是一个简单的结构体。

# 继承

**go-object风格**的继承遵循以下规则。

1. ***抽象类***继承***抽象类***，通过***抽象接口***内嵌另一个***抽象接口***方式实现，可多继承。
2. ***非抽象类***继承***抽象类***，通过内嵌***抽象结构***指针方式方式实现，可多继承。
3. <font color=Red>***非抽象类***继承<u>非顶级</u>***抽象类***时，必须内嵌所有父类的***抽象结构***。</font>

*Animal继承Biology（抽象类继承抽象类）*

```go
// animal.aclass.go

type Animal_ interface {
    // 抽象结构转换方法
    Animal_() *Animal__
    // 内嵌Biology的抽象接口
    Biology_
    // 声明子类须实现的方法
    Eat() bool
}

type Animal__ struct {
    I Animal_
}

func (this *Animal__) Animal_() *Animal__ {
    return this
}
```

*Cat继承Biology（非抽象类继承顶级抽象类）*

```go
// cat.class.go

type Cat struct {
    *Biology__ // 内嵌Biology的抽象结构
}
```

*Bird继承Animal（非抽象类继承非顶级抽象类）*

```go
// bird.class.go

type Bird struct {
    *Biology__ // 内嵌Biology的抽象结构
    *Animal__  // 内嵌Animal的抽象结构 
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
> 以下方式是不合法的，Animal__不是顶级父类，其继承的父类的抽象结构也需要内嵌。
> 
> ```go
> type Bird struct {
>   *Animal__ // 只内嵌Animal的抽象结构是不够的
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

type Biology_ interface {
    Biology_() *Biology__
}

type Biology__ struct {
    I Biology_
    Name string
}

func (this *Biology__) Biology_() *Biology__ {
    return this
}

func (this *Biology__) Breathe() {
    return "I'm breathing"
}

// 构造器，首个参数为抽象接口，其他参数根据实际需要补充
func ExtendBiology(i Biology_, name string) *Biology__ {
    return &Biology__{I: i, Name: name}
}
```

## 非抽象类构造器

***非抽象类***构造器声明规则：名称格式为`New<类名>`，无必要构造参数，根据实际需要自定，返回类的结构体指针。构造基本逻辑是，先创建本类的对象（本结构体指针），再调用抽象类的构造函数，首个参数传入本类对象，其他参数根据实际需要传值，创建的变量用来初始化***非抽象类***内嵌的***抽象结构***。

*Bird类完整声明*

```go
// bird.class.go

type Bird struct {
    *Biology__
    *Animal__
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
    b.Biology__ = ExtendBiology(b, name)
    b.Animal__ = ExtendAnimal(b)
    return b
}
```

# 类声明的其他规则

- 方法接受者使用指针，统一名称为`this`。
- 类的名称、字段、方法和构造器等等，都没有强制是否导出，可根据实际需要选择首字母的大小写。
- 使用单独文件声明类，以区分普通Go代码，文件名格式为，抽象类：`<类名>.aclass.go`，非抽象类：`<类名>.class.go`。

# 多继承二义性

多继承一般都会产生二义性问题，常见的场景：子类继承多个父类，这些父类又继承了相同父类，导致出现二义性问题。例如B和C继承了A，D继承B和C，D调用A的方法时出现二义性问题。

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

**go-object风格**消除了这种场景的二义性问题。 我们用**go-object风格**来改写上面的C++例子。

*go-object风格改写*

```go
package main

import "fmt"

// 类A
type A_ interface {
    A_() *A__
}

type A__ struct {
    I A_
}

func (this *A__) A_() *A__ {
    return this
}

func (this *A__) Foo() {
    fmt.Println("A::foo()")
}

func ExtendA(i A_) *A__ {
  return &A__{I: i}
}

// 类B
type B_ interface {
    B_() *B__
    A_ // 继承类A
}

type B__ struct {
    I B_
}

func (this *B__) B_() *B__ {
    return this
}

func (this *B__) Bar() {
    fmt.Println("B::bar()")
}

func ExtendB(i B_) *B__ {
  return &B__{I: i}
}

// 类C
type C_ interface {
    C_() *C__
    A_ // 继承类A
}

type C__ struct {
    I C_
}

func (this *C__) C_() *C__ {
    return this
}

func (this *C__) Baz() {
    fmt.Println("C::baz()")
}

func ExtendC(i C_) *C__ {
  return &C__{I: i}
}

// 类D
type D struct {
    *A__ // 继承类A
    *B__ // 继承类B
    *C__ // 继承类C
}

func (this *D) Qux() {
    fmt.Println("D::qux()")
}

func NewD() *D {
  d := &D{}
  d.A__ = ExtendA(d)
  d.B__ = ExtendB(d)
  d.C__ = ExtendC(d)
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

为***抽象类***增加一个相同类名的子类，并且没有其他成员字段和方法，那么这个子类相当于可实例化的***抽象类***

*Biology类改为可实例化*
```go
// biology.class.go

type Biology struct {
    *Biology__
}

func NewBiology(name string) *Biology {
    b := &Biology{}
    b.Biology__ = ExtendBiology(b, name)
}
```

## 使非抽象类可继承

和抽象类实例化原理相同，将***非抽象类***改造为***抽象类***，再增加可实例化它的子类。

*Bird类改为可继承*

```go
// bird.aclass.go

type Bird_ interface {
    Bird_() *Bird__
    Animal_
}

type Bird__ struct {
    canSwim bool
}

func (this *Bird__) Bird_() *Bird__ {
    return this
}

func (this *Bird__) Eat() {
    fmt.Println("I'm eating")
}

func (this *Bird__) Swim() {
    if this.canSwim {
        fmt.Println("I'm swimming.")
    } else {
        fmt.Println("I can't swim.")
    }
}

func ExtendBird(i Bird_, canSwim bool) *Bird__ {
    return &Bird__{I: i, canSwim: canSwim}
}
```

```go
// bird.class.go

type Bird struct {
    *Biology__
    *Animal__
    *Bird__
}

func NewBird(name string, canSwim bool) *Bird {
    b := &Bird{}
    b.Biology__ = ExtendBiology(b, name)
    b.Animal__ = ExtendAnimal(b)
    b.Bird__ = ExtendBird(b, canSwim)
    return b
}
```

# 编译器的支持

**go-object风格**是一种自我约定的代码风格，编译器无法检查每一条规则，但一些重要规则还是能得到编译器的支持。

## 子类必须内嵌所有父类的抽象结构

即继承规则的最后一条。

*Dog只继承Animal将编译错误*

```go
// dog.class.go

type Dog struct {
    *Animal__
}

func (this *Dog) Eat() {
    fmt.Println("I'm eating.")
}

func NewDog() *Dog {
    d := &Dog{}
    d.Animal__ = ExtendAnimal(d) // 编译错误：some methods are missing: Biology_() *Biology__
    return d
}
```

## 子类必须实现抽象方法

*Rabbit未实现Animal的方法将编译错误*

```go
// rabbit.class.go

type Rabbit struct {
    *Biology__
    *Animal__
}

func NewRabbit(name string) *Rabbit {
    r := &Rabbit{}
    r.Biology__ = ExtendBiology(r, name)
    r.Animal__ = ExtendAnimal(r) // 编译错误：some methods are missing: Eat()
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
var b Biology_ = bd  // Bird转Biology
var a Animal_ = bd   // Bird转Animal
a = b.(Animal_)      // Biology转Animal
bd = b.(*Bird)       // Biology转Bird
b = a                // Animal转Biology
bd = a.(*Bird)       // Animal转Bird
```

# 方法重写

方法的重写规则：***抽象类***可重写的方法在***抽象接口***和***抽象结构***都需要声明。

*Plant类声明可重写方法*

```go
// plant.aclass.go

type Plant_ interface {
    Plant_() *Plant__
    Say() // 声明抽象结构的方法，使其可重写
}

type Plant__ struct {
    I Plant_
}

func (this *Plant__) Plant_() *Plant__ {
    return this
}

// 在抽象接口也声明此方法，使其可重写
func (this *Plant__) Say() {
    fmt.Println("I'm a plant.")
}

func ExtendPlant(i Plant_) *Plant__ {
    return &Plant__{I: i}
}
```

*Tree类*

```go
// tree.class.go

type Tree struct {
    *Plant__
}

func NewTree() *Tree {
    t := &Tree{}
    t.Plant__ = ExtendPlant(t)
    return t
} 

```

*Flower类*

```go
// flower.class.go

type Flower struct {
    *Plant__
}

func (this *Flower) Say() {
    fmt.Println("I'm a flower.")
}

func NewFlower() *Flower {
    f := &Flower{}
    f.Plant__ = ExtendPlant(f)
    return f
} 

```

*重写方法的调用*

```go
func main() {
    t := NewTree()
    f := NewFlower()
    ps := []Plant_{t, f}

    // 通过抽象接口调用，将调用子类重写的方法，如果子类没有重写则调用父类方法。
    for _, p := range ps {
        fmt.Println(b.Say())
    }
    // Output:
    // I'm a plant.
    // I'm a flower.

    // 子类可通过抽象结构调用父类实现的方法。
    for _, p := range ps {
        fmt.Println(b.Plant__().Say())
    }
    // Output:
    // I'm plant.
    // I'm plant.
}
```

# 判断NULL

**go-object风格**使用Golang的接口表示父类，而Golang有著名的`nil ≠ nil`问题（nil指针赋值给接口，接口≠nil），判断抽象类的是否为nil时，目前最好通过反射。