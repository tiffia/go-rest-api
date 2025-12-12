## func
Functions in Go are first-class citizens. They can be passed as arguments, returned from functions, and assigned to variables. Functions are not automatically invoked - they must be explicitly called.

- func() gets invoked only when called
- Functions can take multiple parameters and return multiple values
- Named return values allow implicit returns
```go
    func volume (length,width,height int) int {
      return length * width * height
    }

    func dimensions(length,width,height int) (area int, volume int) {
      area = length * width
      volume = length * width * height

      return
    }

    func main(){
      area, volume := dimensions(5,6,7)
      fmt.Println("area: ", area, "volume: ", volume)
    }
```

- No classes in Go - use structs and methods instead

## struct
Structs are collections of fields that group related data together. Methods are functions associated with a struct type, similar to class methods. Go uses value receivers (pass by value) or pointer receivers (pass by reference).

Description: A struct is a composite data type that allows you to group multiple fields of different types into a single type. Methods defined on structs enable object-oriented programming patterns.
```go
    type Dimension struct{
      length int
      width int
      height int
    }

    func(d Dimension) Area() int {
      return d.length * d.width ///struct pointer
    }
    func(d Dimension) Volume() int {
      return d.length * d.width * d.height ///struct pointer
    }
    func main(){
      d := Dimension{9,8,7}
      fmt.Println(d.Area())
      fmt.Println(d.Volume())
    }
```

## type
Type definitions allow you to create new types based on existing types. This is useful for semantic clarity and for defining methods on primitive types.

Description: You can define new types using the `type` keyword. For example, `type MyInt int` creates a new type MyInt that has the underlying type int. This enables you to define methods on primitive types.

```go
  type MyInt int
  
  func (m MyInt) Double() MyInt {
    return m * 2
  }
```

## pointer
Pointers hold memory addresses of variables. They allow you to pass variables by reference and modify them in functions.

Description: A pointer stores the memory address of a variable. The `&` operator gets the address of a variable, while the `*` operator dereferences a pointer to access its value.

- `&x` is the memory address of variable x
- `*p` is the value at the address stored in pointer p
- Pointers allow you to modify the original variable
```go
  func main(){
    x,y := 5,10
    n := &x

    fmt.Println(*n)

    *n = 50
    fmt.Println(x)

    t := &y
    *t = *t +30
    fmt.Println(t)
  }

```

## value vs reference receivers
This section demonstrates the difference between value receivers (pass by value) and pointer receivers (pass by reference). Value receivers receive a copy of the struct, while pointer receivers receive a reference to the original struct.

Description: 
- **Value receiver** (pass by value): Changes made inside the method do not affect the original struct
- **Pointer receiver** (pass by reference): Changes made inside the method affect the original struct

```go
  func(d *Dimension) Area() int {
    d.height =8
      return d.length * d.width ///struct pointer
    }
  func(d Dimension) Volume() int {
    d.height = 6
      return d.length * d.width * d.height ///struct pointer
    }
  func main(){
    d := Dimension (10,5,6)
    fmt.Println(d.Area())
    fmt.Println(d)
    fmt.Println(d.Volume())
    fmt.Println(d)
  }
```

Output:
```
50
{10 5 8} // Dimension.height changed to 8 (pointer receiver modified original)
300
{10 5 8} // Dimension.height remains 8 (value receiver created a copy)
```

## databases
Database section - coming soon with examples of connecting to databases from Go applications.
