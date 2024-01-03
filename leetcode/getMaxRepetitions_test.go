/**
 * @Author: evnxo
 * @Description:
 * @File:  getMaxRepetitions_test.go
 * @Version: 1.0.0
 * @Date: 2024/1/2 22:13
 */

package leetcode

import "testing"

func Test_GetMaxRepetitions(t *testing.T) {

	t.Run("example2", func(t *testing.T) {
		s1 := "abaacdbac"
		n1 := 100

		s2 := "adcbd"
		n2 := 4

		t.Log(_2(s1, s2, n1, n2))
	})

	t.Run("example1", func(t *testing.T) {
		s1 := "abdec"
		n1 := 3
		s1 = getRepeatS(s1, n1)
		s2 := "abc"
		n2 := 2
		s2 = getRepeatS(s2, n2)
		t.Log(_1(s1, s2, n1, n2))
	})
}
