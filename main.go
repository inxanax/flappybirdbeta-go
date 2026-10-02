package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("========")
	fmt.Println("")
	fmt.Println("Приготовтесь:")
	fmt.Println("")

	score := 0
	fmt.Println("Ваш счет:", score)
	fmt.Println("")
	fmt.Println("\033[32mСделай 10 очков и выиграй!\033[0m  ")
	fmt.Println("========")
	time.Sleep(2800 * time.Millisecond)
fmt.Println("")
	

	for i := 1; i <= 20; i++ {

		fmt.Println("=============")
		fmt.Println("Я подлетаю к трубе", i)
		fmt.Println("🦆||")
		time.Sleep(1600 * time.Millisecond)

		fmt.Println("Я прохожу через трубу")
		fmt.Println("🦆..")
		time.Sleep(1600 * time.Millisecond)

		if rand.Intn(6) == 5 {
			fmt.Println("Я врезался в трубу :( ")
			fmt.Println("🦆❌ ||")
			break
		}

		fmt.Println("Я прошел через трубу")
		fmt.Println("||🦆")
		time.Sleep(1700 * time.Millisecond)
		score++
	}
	fmt.Println("")
	fmt.Println("Итог:", score)
	time.Sleep(1800 * time.Millisecond)

	fmt.Println("")
	if score > 10 {
		fmt.Println("\033[33mYou win!\033[0m")
		fmt.Println("")
	} else {
		fmt.Println("\033[31mТы не дошел до конца :( \033[0m  ")
		fmt.Println("")
		fmt.Println("\033[36mПопробуй еще раз!\033[0m")
		fmt.Println("")
	}

}
