package main
import "fmt"
import s "strings"
var p = fmt.Println
var gVariable int = 500 //GLOBAL VARIABLE

func main() {
	/* Variable */
	// Age := 22
	// something, something2 := true, false 
	// fmt.Println("Age :", Age)
	// fmt.Println("Something :", something)
	// fmt.Println("Something :", something2)

	/* Number and String */
	// fNumber := 50.00
	// fSecond := 12.6
	// fmt.Println(fNumber / fSecond)
	// msg1 := "Beer"
	// msg2 := "Zezy"
	// msg3 := msg1 + " " + msg2 //Concatenation
	// fmt.Println(msg3)
	// fmt.Println(msg3[5:7]) // Select String 
	// fmt.Println(msg3[0]) // ASCII CODE

	/* Boolean */
	// isEmpty := true
	// isEmpty2 := false
	// fmt.Println(isEmpty)
	// fmt.Println(isEmpty2)
	// someboolen := 5 == 3
	// fmt.Println(someboolen)

	/* Scope */
	//anotherFunc()

	/* Input Scanf */
	// fmt.Print("Input Your Number : ")
	// var input float64
	// fmt.Scanf("%f", &input)
	// output := input * 2
	// fmt.Println("Number is :", output)

	/* IF */
	// fmt.Print("Input Your Number : ")
	// var input float64
	// fmt.Scanf("%f", &input)
	// condition := input > 2
	// if condition {
	// 	fmt.Println("INPUT > 2")
	// } else {
	// 	fmt.Println("INPUT < 2")
	// }

	/* Switch */
	// fmt.Print("Enter Your Number : ")
	// var number int
	// fmt.Scanf("%d", &number)
	// switch number {
	// case 0:fmt.Println("Zero")
	// case 1:fmt.Println("One")
	// case 2:fmt.Println("Two")
	// default: fmt.Println("Unknown")
	// }

	/* For */
	// for i:=1; i<=10; i++ {
	// 	fmt.Println("BEER", i)
	// }

	// i:=1
	// for i<=10 {
	// 	fmt.Println("BEER2", i)
	// 	//i++
	// 	i = i + 1
	// }

	// for i:=1; i<=10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "is Even Number")
	// 	} else {
	// 		fmt.Println(i, "is Odd Number")
	// 	}
	// }

	/* Function */
	// fmt.Print("Enter Your Text : ")
	// var str string
	// fmt.Scanf("%s", &str)
	// dosomething(str)
	
	// fmt.Print("Enter Frist Number : ")
	// var a int
	// fmt.Scanln(&a)
	
	// fmt.Print("Enter Second Number : ")
	// var b int
	// fmt.Scanln(&b)
	// addition(a, b)

	/* Variadic */
	// sumation(1, 1, 1)

	/* Recursive Function */
	// fmt.Println(factorial(5))

	/* Array */
	//var x [5]int
	// x := [5]int{}
	//x := [5]float64{2, 2, 3, 4, 5}
	//  for i:=0; i<len(x); i++ {
	// 	fmt.Println("Array index[", i, "] :", x[i])
	//  }
	// var total float64
	// for _ , value:=range x {
	// 	total += value
	// }
	// fmt.Println(total)
	// fmt.Println("Average :", total/float64(len(x)))

	/* Slice */
	// slice1 := []int{1, 2, 3, 4}
	// slice2 := append(slice1, 5, 6)
	// fmt.Println(slice2)
	
	// arr := [3]float64{1, 2, 3}
	// x := arr[0:3]
	// y:= append(x, 5, 6)
	// fmt.Println(y)

	// slice1 := []int{1, 2, 3}
	// slice2 := make([]int, 2)
	// copy(slice2, slice1) //copy(des, source)
	// fmt.Println(slice2)

	/* Map */
	// x := make(map[string] string)
	// x["TH"] = "THAILAND"
	// x["JP"] = "JAOAN"
	// x["ENG"] = "ENGLAND"
	// x := map[string] string {
	// 	"TH": "THAILAND",
	// 	"JP": "JAOAN",
	// 	"ENG": "ENGLAND",
	// }

	// for keys, val := range x {
	// 	fmt.Println("[", keys, "] =>", val)
	// }

	/* Closure */
	// add := func (x, y int) int {
	// 	result := x + y
	// 	return result
	// }
	// x := 0
	// increment := func () int {
	// 	x++
	// 	return x
	// }
	// fmt.Println(add(1, 2))
	// fmt.Println(increment())
	// fmt.Println(increment())

	/* Pointer */
	// x := 10
	// var p *int
	// p = &x // ชี้ไปยัง Address ที่ x อยู่
	// *p = 20
	// fmt.Println("value is", x)
	// fmt.Println("Address x variable", p)

	/* Structure */
	// var Book1 Books 
	// Book1.Username = "Beer"
	// Book1.Fistname = "Phanudet"
	// Book1.Lastname = "Khawilai"
	// Book1 := Books{ Username: "Beer", Fistname:" Phanudet", Lastname: "Khawilai" }
	// fmt.Println(Book1.Username)
	// fmt.Println(Book1.Fistname)
	// fmt.Println(Book1.Lastname)

	/* Go Routine */
	// go f(0)
	// var input string
	// fmt.Scanln(&input)

	/* String Function */
	//Contains, Count, HasPrefix, HasSuffix, Index, Join, Repeat, Replace, Split, ToLower, ToUpper
	p("Contains:", s.Contains("test", "est")) //ทำการเช็คว่ามี est ของชุดข้อความนั้นหรือไม่ แล้ว return bool
	p("Contains:", s.Count("test", "t")) //ทำการนับตัวอักษร t ว่ามีกี่ตัวของชุดข้อความนั้นๆ
	p("HasPrefix:", s.Count("test", "te")) //เช็คตัวขึ้นต้น
	p("HasSuffix:", s.Count("test", "st")) //เช็คตัวลงท้าย
	p("Index:", s.Index("test", "e")) //return index ของ string
	p("Join:", s.Join([]string{"a", "b"}, "...")) //ต่อ string
	p("Index:", s.Repeat("a", 5)) //ทำซ้ำ
	p("Replace:", s.Replace("foo", "o", "0", -1)) 
	p("Replace:", s.Replace("fooo", "o", "0", 2)) 
	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("Len:", len("Hello"))
}
/* Go Routine */
// func f(n int) {
// 	for i:=0; i<10; i++ {
// 		fmt.Println(n, ":", i)
// 	}
// }

/* Structure */
type Books struct {
	Username string
	Fistname string
	Lastname string
	Age int
}

// func anotherFunc() {
// 	localVariable := 40
// 	fmt.Println("GLOBAL VARIABLE :", gVariable)
// 	fmt.Println("LOCAL VARIABLE :", localVariable)
// }

/* Function */
// func dosomething(str string) {
// 	fmt.Println("YOUR TEXT :", str)
// }
// func addition(a int, b int) {
// 	output := a + b
// 	fmt.Println("Sum of Number You Enter :", output)
// }

/* Variadic */
// func sumation(args...int) {
// 	var total int
// 	for _ , n := range args {
// 		total +=n
// 	}
// 	fmt.Println("TOTAL IS :", total)
// }

/* Recursive Function */
// func factorial(num int) int {
// 	//if num is 5 Fac => 5*4*3*2*1
// 	if num == 0 {
// 		return 1
// 	} else {
// 		return num * factorial(num-1)
// 	}
// }
