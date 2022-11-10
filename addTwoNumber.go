package main 

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode
    var h3 *ListNode
    ten := 0

    for l1 != nil || l2 != nil {
        nVal := 0
        if l1 != nil {
            nVal += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            nVal += l2.Val
            l2 = l2.Next
        }
        nVal += ten
        if nVal > 10 {
            nVal = nVal % 10
            ten = 1
        }else {
            ten = 0
        }
        nNode := ListNode{nVal, nil}

        if h3 == nil {
            h3 = &nNode
            l3 = &nNode
        }else {
            l3.Next = &nNode
            l3 = l3.Next
        }
    }

    if ten > 0 {
        nNode := ListNode{1, nil}
        l3.Next = &nNode
        l3 = l3.Next
    }

    return h3
}

