// leetcode23
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

    l := len(lists)
    if l == 0 {
        return nil
    }
    if l == 1 {
        return lists[0]
    }

    var ret *ListNode = &ListNode{}
    for i:=0; i<l; i++ {

        h := lists[i]
        for h != nil {
            bn := h.Next
            insertNode(ret, h)
            h = bn
        }
    }

    return ret.Next
}

func insertNode(ret *ListNode, n *ListNode) *ListNode {
    if ret.Next == nil {
        ret.Next = n
        n.Next = nil
        return ret
    }

    var p, pre *ListNode
    p = ret.Next
    pre = ret
    for p != nil && n.Val >= p.Val {
        pre = p
        p = p.Next
    }
    if p == nil {
        pre.Next = n
        n.Next = nil
    } else if p == ret.Next {
        ret.Next = n
        n.Next = p
    } else {
        pre.Next = n
        n.Next = p
    }

    return ret
}
