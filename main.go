package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var totalCases int
	fmt.Scanf("%d\n", &totalCases)

	results := make([]int, totalCases)

	for i := 0; i < totalCases; i++ {
		var x int
		fmt.Scanf("%d\n", &x)

		inputLine, _ := reader.ReadString('\n')
		numbers := strings.Fields(inputLine)

		sum := 0
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			if num > 0 {
				sum += num * num
			}
		}
		results[i] = sum
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
