package reptile

import (
	"testing"
)

/*
* @Author:hanyajun
* @Date:2019/9/23 21:08
* @Name:reptile
* @Function:
 */

func Test_GetNewsContent(t *testing.T) {

}

func Benchmark_GetNewsContent(b *testing.B) {
	for i := 0; i < 4; i++ {
		b.Log("", i)
	}
}
func Benchmark_GetNewsContent2(b *testing.B) {

}
