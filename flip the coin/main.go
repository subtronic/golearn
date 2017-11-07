package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
)

func printResult(result map[string]int){
	fmt.Println("Summary you drop ", result["head"] + result["tail"], " with head -", result["head"], " and with tail - ", result["tail"])
}

func main()  {
	isRunning := true
	fmt.Println("Emulate drop the coin")
	fmt.Println("For yes just enter")
	fmt.Println("If you want to end, just print `stop`")
	result := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	for isRunning {
		fmt.Println("Drop the coin?")
		decision,_ := reader.ReadString('\n')
		decision = strings.TrimRight(decision, "\n")

		switch decision{
			case "stop":
				isRunning = false
			case "result":
				printResult(result)
			default:
				drop := rand.Intn(100)
				if drop > 50 {
					result["head"]++
				} else {
					result["tail"]++
				}
		}
	}

	printResult(result)
}
