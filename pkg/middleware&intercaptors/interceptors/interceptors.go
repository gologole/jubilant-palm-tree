package interceptors

import (
	"fmt"
	"time"
)

type Interceptor func(func())

func LogInterceptor(fn func()) {
	start := time.Now()
	fmt.Println("Starting function")

	fn() // вызов оригинальной функции

	duration := time.Since(start)
	fmt.Printf("Function completed in %s\n", duration)
}

//
//func main() {
//	myFunction := func() {
//		time.Sleep(2 * time.Second) // имитация работы функции
//		fmt.Println("Function is executing")
//	}
//
//	// Использование логгера-интерцептора
//	LogInterceptor(myFunction)
//}
