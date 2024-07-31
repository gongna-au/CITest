package main

import "fmt"

func main() {
	fmt.Println("Hello, World!") // 这行代码是有效的

	// 下面的代码将引发错误
	for i := 0; i < 10; i++ {
		i = 1

	}

	// 这行代码将引发一个潜在的错误
	// 这里的错误是因为没有处理错误
	var err error
	_, err = fmt.Println("This will not be checked") // 忽略错误
	fmt.Println(err)

}
