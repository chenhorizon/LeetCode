// leetcode61
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    nm := make(map[int]*ListNode)
    i := 0
    h := head
    for h != nil {
        i++
        nm[i] = h
        h = h.Next
    }

    k = k % i
    if k == 0 {
        return head
    }

    nm[i-k].Next = nil
    nm[i].Next = nm[1]
    return nm[i-k+1]

}

