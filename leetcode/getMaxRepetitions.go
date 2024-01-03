/**
 * @Author: evnxo
 * @Description:
 * @File:  getMaxRepetitions.go
 * @Version: 1.0.0
 * @Date: 2024/1/2 22:11
 */

package leetcode

// 复杂度
// 时间: s1*s2
// 空间: s2
// 找到循环节
func _2(s1, s2 string, n1, n2 int) int {
	if n1 == 0 {
		return 0
	}
	var (
		s1cnt, index, s2cnt int
		recall              map[int][]int
		preloop, inLoop     []int
	)
	s1cnt, index, s2cnt = 0, 0, 0
	recall = make(map[int][]int)
	preloop, inLoop = make([]int, 0), make([]int, 0)
	for {
		s1cnt++
		for i := range s1 {
			if s1[i] == s2[index] {
				index += 1
			}
			if index == len(s2) {
				s2cnt++
				index = 0
			}
		}
		if s1cnt == n1 {
			return s2cnt / n2
		}

		if val, ok := recall[index]; ok {
			preloop = []int{val[0], val[1]}
			inLoop = []int{s1cnt - val[0], s2cnt - val[1]}
			break
		} else {
			recall[index] = []int{s1cnt, s2cnt}
		}

	}
	ans := preloop[1] + (n1-preloop[0])/inLoop[0]*inLoop[1]
	rest := (n1 - preloop[0]) % inLoop[0]
	for i := 0; i < rest; i++ {
		for j := range s1 {
			ch := s1[j]
			if ch == s2[index] {
				index++
			}
			if index == len(s2) {
				ans++
				index = 0
			}
		}
	}
	return ans / n2
}

// 复杂度
// 时间:O(n1*s1)
// 空间:O(1)
func _1(s1, s2 string, n1, n2 int) int {
	// 暴力破解法,贪心算法
	index := 0
	repeat_count := 0
	s1_size, s2_size := len(s1), len(s2)
	for i := 0; i < n1; i++ {
		for j := 0; j < s1_size; j++ {
			if s1[j] == s2[index] {
				index++
			}
			if index == s2_size {
				index = 0
				repeat_count++
			}
		}
	}
	return repeat_count / n2
}

func getRepeatS(s string, n int) string {
	a := ""
	for i := 0; i < n; i++ {
		a += s
	}
	return a
}
