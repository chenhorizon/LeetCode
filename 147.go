// leetcode147
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    h := head
    ph := head
    p := head.Next
    pre := head

    for p != nil {

        for h != p && p.Val > h.Val {
            ph = h
            h = h.Next
        }

        if h == head {
            pre.Next = p.Next
            p.Next = h
            head = p
            h = head
            p = pre.Next
        } else if h == p {
            pre = p
            p = p.Next
            h = head
            ph = head
        } else {
            pre.Next = p.Next
            p.Next = h
            ph.Next = p
            p = pre.Next
            h = head
        }
    }

    return head
}
