// leetcode19
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nl := make(map[int]*ListNode)
    var h *ListNode = head
    i := 0
    for h != nil {
        nl[i] = h
        i++
        h = h.Next
    }
    if n > i {
        return nil
    }
    if n == 1 {
        if i >= 2 {
            nl[i-2].Next = nil
        } else {
            return nil
        }
    } else if n == i {
        head = nl[1]
    } else {
        nl[i-n-1].Next = nl[i-n+1]
    }

    return head
}
