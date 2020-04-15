package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Speak(string) string
}

type Man struct{}

func (m *Man) Speak(thing string) (talk string) {
	if thing == "sb" {
		talk = "你太给力了"
	} else {
		talk = "你不行"
	}
	return
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v - kind:%v\n", v.Name(), v.Kind())
}

func main() {
	fmt.Println("hello")

	s3 := []int{1, 2, 3}
	s4 := s3
	s4[0] = 100

	s4 = append(s4, 120, 11, 444)
	s5 := make([]int, 3, 5)
	copy(s5, s4) //copy 不返回数据
	s6 := removeItem(s4, 2)
	fmt.Println(s3, s4, s5, s6)

	pin := 11

	fmt.Println(&pin)
	fmt.Printf("%T\n", &pin)

	var q1 = new(int)
	*q1 = 100
	fmt.Printf("%d\n", *q1)

	map1 := make(map[string]int, 8)
	map1["q"] = 100
	map1["w"] = 200
	fmt.Println(map1)
	fmt.Println((map1["q"]))

	v, ok := map1["a"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("a 不存在")
	}

	for k, v := range map1 {
		fmt.Printf("%s %d \n", k, v)
	}

	delete(map1, "q")

	for k, v := range map1 {
		fmt.Printf("%s %d \n", k, v)
	}

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	// var mm People = Man{}
	// think := "bitch"
	// fmt.Println(mm.Speak(think))

	var a float32 = 3.14
	reflectType((a))
}

func removeItem(arr []int, index int) []int {
	if index > len(arr) {
		return nil
	}
	//会改变自己，所以要用copy
	tmp := make([]int, len(arr), len(arr))
	copy(tmp, arr)
	return append(tmp[:index], tmp[(index+1):]...)
}

//使用go build 生成可执行文件，再用terminal可以直接运行
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
