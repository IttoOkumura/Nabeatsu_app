package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	count := 0

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("٩( ᐛ ) さん%d!\n", i)
			count++
		} else {
			var num string
			num = strconv.Itoa(i)
			if strings.Contains(num, "3") == true {
				fmt.Printf("٩( ᐛ ) さん%d!\n", i)
				count++
			}

		}
	}

	fmt.Printf("%d回アホになった", count)
}
