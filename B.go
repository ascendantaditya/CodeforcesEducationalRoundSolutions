package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(scanner *bufio.Scanner) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(input[i])
	}

	if n > 2 {
		fmt.Println("NO")
		return
	}

	if a[1] == a[0]+1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for t > 0 {
		solve(scanner)
		t--
	}
}
