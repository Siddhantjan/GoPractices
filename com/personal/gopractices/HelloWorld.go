package main

import (
	"fmt"
	"math/rand"
)

//type Vertex struct {
//	X int
//	Y int
//}

func hello(arr []int, ch chan int) {
	for _, val := range arr {
		ch <- val * 3
	}
}

func main() {
	//fmt.Println("hey Siddhant!")
	//var i int
	//i = 12
	//i:=12
	//var i int
	//fmt.Scanln(&i)
	//fmt.Println(i)
	//var s string
	//s = "hello"
	//fmt.Println(s)
	//fmt.Println("%q\n", s)
	//i := 5
	//switch i {
	//case 1:
	//	fmt.Println("hey")
	//	break
	//case 2:
	//	fmt.Println("Great")
	//	break
	//default:
	//	fmt.Println("No problem")
	//
	//}
	//	sum := 0

	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//i := 0
	//for i < 10 {
	//	sum += i
	//	i++
	//}
	//fmt.Println(sum)
	//i := 0
	//for {
	//	if i == 10 {
	//		break
	//	}
	//	sum += i
	//	i++
	//}
	//fmt.Println(sum)
	//fmt.Println(Vertex{1, 2})
	//var a [10]int
	//for i := 0; i < 10; i++ {
	//	a[i] = i + 1
	//}
	//fmt.Println(a)
	//var s []int = a[1:4]
	//fmt.Println(s)

	//fmt.Println("main starts")
	//ch := make(chan int)
	//arr := []int{1, 2, 4}
	//go hello(arr, ch)
	//res := <-ch

	//time.Sleep(time.Second)
	//fmt.Println(res)
	//res = <-ch
	//fmt.Println(res)
	//res = <-ch
	//fmt.Println(res)

	//var isTrue bool = true
	//var isFalse bool = false
	//// AND
	//if isTrue && isFalse {
	//	fmt.Println("Both Conditions need to be True")
	//}
	//// OR - not exclusive
	//if isTrue || isFalse {
	//	fmt.Println("Only one condition needs to be True")
	//}

	//var bus interface{}
	//bus = "1"
	//s, _ := strconv.ParseInt(bus.(string), 10, 16)
	//fmt.Println(s)
	//ch := make(chan int)
	//ch1 := make(chan int)
	//arr := []int{2, 3, 4}
	//go hello(arr, ch)
	//go hello1(ch, ch1)
	//res := <-ch1
	//fmt.Println(res)
	//values := mak//e(// int, 2)
	//go CalculateValues(values)
	//
	//for i := 0; i <= 10; i++ {
	//	time.Sleep(1 * time.Second)
	//	value := <-values
	//	fmt.Println(value)
	//}
	fmt.Println("main ends")

}

func CalculateValues(values chan int) {
	for i := 0; i <= 10; i++ {
		value := rand.Intn(10)
		fmt.Printf("Value Calculated: %d\n", value)
		values <- value
	}
}
func hello1(ch, ch1 chan int) {
	for i := 0; i < 3; i++ {
		val := <-ch
		ch1 <- val * 3
	}
}
