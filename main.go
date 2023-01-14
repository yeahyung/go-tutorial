package main

import (
	"fmt"
	"strings"
)

// 컴파일을 하면 main package 의 main func 을 찾아 실행하게 됨
// 따라서, 컴파일 시 반드시 main package & main func 이 있어야 한다.
func main() {
	const name string = "yea" // cannot change
	var name2 string = "yea"  // can change
	name3 := "yea"            // go 가 type 을 찾아서 설정, func 안에서만 사용 가능
	fmt.Println(name2)
	fmt.Println(name3)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("yea")
	fmt.Println(totalLength, upperName)
	fmt.Println(lenAndUpper2("yea"))

	repeatMe("yea1", "yea2", "yea3", "yea4")

	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println("superAdd: ", total)

	fmt.Println(canDrink(16))
	fmt.Println(canDrink(20))

	copyValue()
	sliceTest()
	mapTest()
	structTest()
}

//func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

// func can return multiple values
func lenAndUpper(name string) (int, string) {
	// defer --> func 이 끝나면 실행시킬 내용
	defer fmt.Println("I'm done")
	return len(name), strings.ToUpper(name)
}

// naked return 형태, return 문에 일일히 반환해야할 param 을 적어주지 않아도 됨
// 단, func 선언문에 return 에 필요한 변수 이름을 지정해줘야함
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	//for index, number := range numbers {
	for _, number := range numbers {
		total += number
	}
	return total
}

func canDrink(age int) bool {
	// variable expression, 아래와 같이 변수 생성 가능
	// 코드를 읽는 입장에선 koreanAge 변수가 if 문 내에서만 쓰일거다 라고 예상 가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canDrinkWithSwitch(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	//switch koreanAge := age + 2; koreanAge {
	//case 10:
	//	return false
	//case 18:
	//	return true
	//}
	return true
}

func copyValue() {
	a := 2
	// a 의 값을 b 에 넣어주는 것일 뿐, 메모리 주소는 다름
	b := a
	a = 5

	c := &a
	a = 10
	fmt.Println(a, b, *c)
	fmt.Println(&a, &b, c)
}

func sliceTest() {
	names := []string{"yea", "yea1", "yee2"}
	// append 는 names 를 수정하지 않고 새로운 slice 를 return
	names = append(names, "yea3")
	fmt.Printf("%v\n", names)
}

func mapTest() {
	//          key    value
	yea := map[string]string{"name": "yea", "age": "12"}

	for key, value := range yea {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func structTest() {
	favFood := []string{"kimchi", "kimchi2"}
	// yea := person{"yea", 12, favFood}
	yea := person{name: "yea", age: 12, favFood: favFood}
	fmt.Println(yea)
}
