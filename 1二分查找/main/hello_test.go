package main

//单个测试
/*
func TestHello(t *testing.T) {
	res := getsum(5)
	if res != 15 {
		t.Fatalf("testing!!!,错误结果是%v", res)
	}
	t.Log("测试正确")
}
*/

/*
// 表组测试
const checkMark = "\u2713"
const ballotX = "\u2717"

func TestHello(t *testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4, 5}
	var res1 []int
	res1 = []int{1, 3, 6, 10, 15}

	t.Log("开始进行测试了：")
	{
		for i, num := range nums {
			t.Log("当前测试的用例为n=", num)
			{
				res := getsum(num)
				if res != res1[i] {
					t.Fatalf("测试出错", ballotX)
				}

				t.Log("测试正确", checkMark)
			}

		}

	}

}

func TestSearchRange(t *testing.T) {

	var nums []int
	nums = []int{5, 7, 7, 8, 8, 10}
	var target int
	target = 8
	var res []int
	res = SearchRange(nums, target)
	var res1 []int
	res1 = []int{3, 4}
	i := 0
	for ; i < len(res); i++ {
		if res[i] != res1[i] {
			t.Fatalf("结果错误：")
		}
	}
	if i == 2 {
		t.Log("测试成功")
	}
}
*/

// 基准测试
import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	num := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
