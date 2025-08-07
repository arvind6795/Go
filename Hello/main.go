package main

import (
    "fmt"
	"math"
	// "bufio"
	// "os"
    // "strconv"
)
//package level
// var x float32=32.1
// var (
// 	id int = 101
// 	fname string="Elliot"
// 	sname string="Smith"
// )
// var a int =12//package level declaration
// var A int=33//global level outside the package level can be access

func main(){
	// One line comment
	/* Multiline comment
	*/
	// fmt.Println("Hello Universe!")
	// var a int //declaration
	// a=12 //intitialization
    // var b int =13 // in one go both
    // c:=33.1
	// fmt.Println("c =",c)
	// fmt.Println(id,fname,sname)
	// fmt.Println(a)
	// var a int =21//local declaration
	// fmt.Println(a)
	//type converstion
	// var a int =20
	// fmt.Printf("%v,%T",a,a)
	// var b float32
	// b=float32(a)//int to float
	// fmt.Printf(" %v,%T",b,b)
	// s:=strconv.Itoa(a)//int to string
	// fmt.Printf(" %v,%T",s,s)
	// var s string="123"
	// i , err := strconv.Atoi(s)//string to int note:i will store if string is convertible to int otherwise err will store error
	// string to float
	// f, err := strconv.ParseFloat("3.14", 64)
	// if err!=nil{
	// 	fmt.Println("error",err)
	// }else{
	// 	fmt.Println(f)
	// }
	//Primitves datatype
	// var n bool=false//boolean
	// n:=1==1
	// m:=1==3
	// var n bool //by default false
	// fmt.Printf("%v,%T\n",n,n)
	// fmt.Printf("%v,%T\n",m,m)
	// var x uint8=22
	// fmt.Printf("%v,%T\n",x,x)
	// a:=10
	// b:=4
	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)
	// fmt.Println(a%b)
	// var a int=10
	// var b int8=3
	// fmt.Println(a+b)//Mismatch int type won't allow in go lang
	// bitwise operator & | ^
	// a:=10 //1010
	// b:=3  //0011
	// fmt.Println(a&b) //0010
	// fmt.Println(a|b) //1011
	// fmt.Println(a^b) //1001
	// fmt.Println(a&^b) //0100
	// a:=8
	// // a=a<<3 //2^3*2^3 for left-shift
	// // fmt.Println(a)
	// a=a>>3 //2^3 / 2^3 for right shift
	// fmt.Println(a)
	//floating points
	// var n float64=1.2e75
	// fmt.Println(n)
	//complex number
	// var x complex64=1+2i
	// var z complex64=complex(4,5)
	// fmt.Printf("%v,%T",x,x)
	// fmt.Printf("%v,%T",z,z)
	// s:="Hello Universe"
	// d:=[]byte(s)
	// fmt.Println(d)// prints bytes of each character in string as their respective ascii value
	// var s1 rune='E'//a rune is a built-in type that serves as an alias for int32.
	// fmt.Printf("%v,%T",string(s1),s1)
	// strings
	// fmt.Printf("Name: %q","Elliot")//double quoted strings. and %s for normal string manipulation.
	//buffio scanner for taking input from users
	// scanner:=bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type something: ")
	// scanner.Scan()
	// input:=scanner.Text()
	// fmt.Printf("We typed : %q",input)
	// var age int
	// fmt.Printf("Enter your age: ")
	// fmt.Scanf("%d",&age)
	// fmt.Println("I got your age: ",age)
	// scanner:=bufio.NewScanner(os.Stdin)
	// fmt.Printf("In which year you were born: ")
	// scanner.Scan()
	// input,_:=strconv.Atoi(scanner.Text())
	// fmt.Printf("So you are %d years old",2025-input)
	// If-Else conditions
	// age:=12
	// if age>=18{
	// 	fmt.Println("You can drive")
	// }else if age>=14{
	// 	fmt.Println("You can drive with parents")
	// }else{
	// fmt.Println("you can't sorry")
	// }
	//Loops
	// x:=3
	// for x<6{
	// 	fmt.Println(x)
	// 	x++
	// }
	// for x:=0;x<=100;x++{
	// 	if x!=0 && x%3==0 && x%7==0 && x%9==0{
	// 		fmt.Println(x)
	// 		break
	// 	}
	// }
	//switch statements
	// ans:=5
	// switch ans{
	// case 1:
	// 	fmt.Println("1. one")
	// 	fmt.Println("2. one")
	// case 2:
	// 	fmt.Println("Second")
	// default:
	// 	fmt.Println("Not a case we are looking for")
	// }
	// switch{
	// case ans<5:
	// 	fmt.Println("less than 5")
	// case ans>5:
	// 	fmt.Println("greater than 5")
	// default:
	// 	fmt.Println("Five")
	// }
	// arrays
	// var arr [5]int
	// arr[0]=101
	// fmt.Println(arr)
	// arr:=[6]int{2,3,4,5,6,7}
	// fmt.Println(arr)
	// fmt.Println(len(arr))
	//sum of elements in array
// 	sum:=0
// 	for i:=0;i<len(arr);i++{
// 		sum+=arr[i]
// }
// 	fmt.Println(sum)
// 	//2d arr
// 	// arr2d:=[2][2]int{{2,3},{5,6}}
// 	// fmt.Println(arr2d[1][0])
// slice
// var s[] int=arr[1:3]
// fmt.Println(len(s))
// fmt.Println(cap(s))//capacity of slice
// creating new slice
// var a[]int=[]int{3,4,5,6,7,8,9,12,34,56,3,5}
// fmt.Println(cap(a))
// b:=append(a,10)//creating new slice by appending previous slice
// fmt.Println(b)
// make() for slice
// a:=make([]int,5)
// fmt.Println(a)
// fmt.Printf("%T",a)
// for i:=0;i<len(a);i++{
// 	fmt.Println(a[i])
// }
// range and slice/Array examples
// for i,ele :=range a{
// 	fmt.Printf("%d: %d\n",i,ele)
// }
// find duplicate in array
// for i,ele1 :=range a{
	// for j,ele2 := range a{
	// 	if ele1==ele2 && j>i{
	// 		fmt.Println(ele1)
	// 	}
	// for j:=i+1;j<len(a);j++{
	// 	ele2:=a[j]
	// 	if ele1==ele2{
	// 		fmt.Println(ele1)
	// 	}
	// }
	// }
	//maps(key:value)
	// var mp map[string]int=map[string]int{
	// 	"Apple":5,
	// 	"orange":3,
	// 	"pear":9,
	// }
	// // fmt.Println(mp["Apple"])
	// // mp["Apple"]=300//update value
	// // mp["elliot"]=99//add new item to map
	// val,ok:=mp["pear"]
	// fmt.Println(val,ok)
	// fmt.Println(len(mp))
	//functions
	// ans1,ans2:=add(21,7)
	// fmt.Println(ans1,ans2)
	// x:=test
	// x(9)
	// test:=func(x int)int{
	// 	// fmt.Println(x)
	// 	return x * -1
	// }(7)
	// fmt.Println(test)
	// test:=func(x int)int{
	// 	// fmt.Println(x)
	// 	return x * -1
	// }
	// test3:=func(x int)int{
	// 	return x*7
	// }
	// test2(test3)
	// x:=returnfunc("hello universe")
	// returnfunc("I'm genius")()
	// x()
	// x:=[]int{3,4,5}
	// fmt.Println(x)
	// changeFirst(x)
	// fmt.Println(x)
	//pointers and derefrence
	// x:=7
	// y:=&x
	// fmt.Println(x,y)//0xc00008c0a8 where the value 7 is stored
	// *y=9//derefrence the value of x.
	// fmt.Println(x,y)
	// toChange:="Hello"
	// // fmt.Println(toChange)
	// // changeValue(&toChange)
	// // fmt.Println(toChange)
	// var ptr *string= &toChange
	// fmt.Println(*ptr,ptr,&ptr)//Value:Hello pointer:0xc00009c060 ptr(ptr):0xc00008e058//pointer of ptr
	//Structs and Custom Types
	// var p1 Point=Point{2,3}
	// p2:=Point{-2,-3}
	// fmt.Println(p1.x,p2.x)
	// p1:=&Point{y:3}
	// fmt.Println(p1)
	// changeX(p1)
	// fmt.Println(p1)
	// c1:=Circle{4.5,&Point{2,3}}
	// fmt.Println(c1.x)
	// s1:=Student{"Elliot",[]int{70,85,90,50},21}
	// fmt.Println(s1)
	// average:=s1.getAverageGrade()
	// fmt.Println(s1.getMaxGrade())
	// fmt.Println(average)
	c1:=circle{4.5}
	r1:=rectangle{2,6}
	s1:=square{12.5}
	shapes:=[]Shape{c1,r1,s1}
	for _,v:=range shapes{
		fmt.Println(v.area())
	}
}
//Interface
type Shape interface{
	area() float64
}
type circle struct{
	radius float64
}
type square struct{
	side float64
}
type rectangle struct{
	width float64
	height float64
}
func (r rectangle) area() float64{
	return r.width*r.height
}
func (c circle) area()float64{
	return math.Pi*c.radius*c.radius
}
func(s square) area() float64{
	return s.side*s.side
}
//struct  methods
// type Student struct{
// 	name string
// 	grade []int
// 	age int
// }
// func (s *Student) getAverageGrade()float32{
// 	sum:=0
// 	for _ , v:=range s.grade{
// 		sum+=v
// 	}
// 	return float32(sum)/float32(len(s.grade))
// }
// func (s *Student) getMaxGrade()int{
// 	curr_max:=0
// 	for _,v:=range s.grade{
// 		if curr_max<v{
// 			curr_max=v
// 		}
// 	}
// 	return curr_max
// }
//embedded struct
// type Circle struct{
// 	radius float64
// 	*Point
// }
//struct
type Point struct{
	x int32
	y int32
}
func changeX(pt *Point){
	pt.x=1000
}
// func changeValue(str *string){
// 	*str="changed!"
// }
// func changeValue2(str string){
// 	//this function won't change anything because we are passing value that will get str="hello"
// 	str="changed!"
// }
// func test(x int){
// 	fmt.Println("test",x)
// }
// func add(x ,y int)(z1 ,z2 int){
// 	defer fmt.Println("processing")//delay the execution of function.
// 	z1=x+y
// 	z2=x-y
// 	fmt.Println("before return")
// 	return
// }
//func with parameter as a function
// func test2(myfunc func(int)int){
// 	fmt.Println(myfunc(7))
// }
//return functions
// func returnfunc(x string)func(){
// 	return func() {fmt.Println(x)}
// }
//mutable function example
// func changeFirst(slice[]int){
// 	slice[0]=1000
// }