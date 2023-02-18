// 【問題】
// 1-100までの数字で、3がつく数字と、3の倍数のときに（例）「٩( ᐛ ) さん(ここは数字)!」を表示させる
// そして、何回アホ(٩( ᐛ ) これ)になったかを出力する「〜回アホになった」
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
