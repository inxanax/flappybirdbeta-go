package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	fmt.Println("Get ready")
	score := 0
	fmt.Println("Ваш счет", score)
	time.Sleep(2800 * time.Millisecond)

number := rand.Intn(20 + 1)
fmt.Println("Выпало:",number)


for i := 1; i <= number; i++ {
	
	fmt.Println("=============")
	fmt.Println("Я подлетаю к трубе", i)
	fmt.Println("🦆||")
	time.Sleep(1600 * time.Millisecond)

	fmt.Println("Я прохожу через трубу")
	fmt.Println("🦆..")
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("Я прошел через трубу")
	fmt.Println("||🦆")
	time.Sleep(1700 * time.Millisecond)
	score++
}
	fmt.Println("Итог", score)
	time.Sleep(1800 * time.Millisecond)

	fmt.Println("")
	if score > 10 {
		fmt.Println("\033[33mYou win!\033[0m")
		fmt.Println("")
	} else {
		fmt.Println("Ты не дошел до конца :(")
		fmt.Println("")
	}


}
