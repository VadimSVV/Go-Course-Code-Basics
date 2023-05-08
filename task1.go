// // Определение пакета main
// package main
// // Импорт пакета fmt
// import "fmt"
// // Определение функции main
// func main() {
//   // Вызов функции Print из пакета fmt
//   // Отступ 1 таб
//   fmt.Print("Hello, Friend!") // В конце не нужна точка с запятой
// }

wg := sync.WaitGroup{}

for i := 0; i < 3; i++ {
	wg.Add(1)
	go func() {
		fmt.Println("Go!")
		wg.Done()
	}()
}

wg.Wait()