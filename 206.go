// leetcode 206
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var cur *ListNode = head.Next
    var pre *ListNode = head
    for cur != nil {
        pre.Next = cur.Next
        cur.Next = head
        head = cur
        cur = pre.Next
    }

    return head
}
