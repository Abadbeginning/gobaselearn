package stringofcharacters

import (
	"fmt"
	"unicode/utf8"
)

func strTest() {
	str := "去年 10 月，我们对外开源了 Quarkd（quark design 缩写）。开源之初，我们给自己定下 star 数量超过 70 就团建吃烤肉的目标，却意外得到了 1600 多 star，很受鼓舞～"
	fmt.Println("长度 --> ", utf8.RuneCountInString(str))
}

// func main() {
// 	strTest()
// }
