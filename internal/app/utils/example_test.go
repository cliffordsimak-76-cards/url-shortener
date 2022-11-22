package utils

import "fmt"

func ExampleStringToMD5() {
	out := StringToMD5("https://yandex.ru")
	fmt.Println(out)

	// Output:
	// e9db20b246fb7d3f
}
