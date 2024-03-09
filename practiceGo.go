package main

import(
	"fmt" 
	"time"
)


func main() {
    fmt.Println("Hello,Im First Go program writer")
 fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    fmt.Println("Substraction -----------",5-2)

        var b, c string = "1","2"
    fmt.Println("Variable String",b,"", c)

    var d bool
    fmt.Println("Boolean value of d ",d)

 for i := range 3 {
        fmt.Println("range", i)
    }

      for {
        fmt.Println("loop")
        break
    }

     for n := range 6 {
        if n%2 != 0 {
            continue
        }
        fmt.Println("For loop with range 6 ",n)
    }

    for i := range 10{

	    fmt.Println("printing 0 to 10",i)
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    i:=2
    fmt.Println("Write ",i, " as")
    switch i{
    case 1:
	    fmt.Println("one")
    case 2:
	    fmt.Println("Two")
    case 3:
	    fmt.Println("Three")
    }
fmt.Println("Current time ",time.Now())


    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
   
    fmt.Println("Hour ",time.Now().Hour())


    whatAmI := func(i interface{}) {
        switch t := i.(type){
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type ", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")


}

