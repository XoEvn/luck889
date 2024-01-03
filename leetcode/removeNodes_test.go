/**
 * @Author: evnxo
 * @Description:
 * @File:  removeNodes_test.go
 * @Version: 1.0.0
 * @Date: 2024/1/3 21:19
 */

package leetcode

import "testing"

func Test_RemoveNodes(t *testing.T) {
	head1 := &listNode{
		Val: 5,
		Next: &listNode{
			Val: 2,
			Next: &listNode{
				Val: 13,
				Next: &listNode{
					Val: 3,
					Next: &listNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	head2 := &listNode{
		Val: 1,
		Next: &listNode{
			Val: 1,
			Next: &listNode{
				Val: 1,
				Next: &listNode{
					Val: 1,
					Next: &listNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	t.Run("example1", func(t *testing.T) {
		removeNodes(head1)
		removeNodes(head2)
	})

	t.Run("example2", func(t *testing.T) {
		removeNodes2(head1)
		removeNodes2(head2)
	})

	t.Run("example3", func(t *testing.T) {
		removeNodes3(head1)
		removeNodes3(head2)
	})
}
