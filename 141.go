// leetcode141
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

    if head == nil || head.Next == nil {
        return false
    }

    var fast, slow *ListNode = head, head
    for slow != nil && fast != nil {
        if slow.Next != nil {
            slow = slow.Next
        } else {
            return false
        }

        if fast.Next != nil && fast.Next.Next != nil {
            fast = fast.Next.Next
        } else {
            return false
        }

        if slow == fast {
            return true
        }
    }

    return false
}
