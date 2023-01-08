package importTest

import "fmt"

// import 에서 사용할 수 있는 func 은 첫 글자가 반드시 대문자로 시작되어야 한다.
// 소문자로 실행된 것은 private method 로 외부에서 참조할 수 없음

func sayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}
