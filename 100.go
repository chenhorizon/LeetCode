// leetcode 100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    var ls, rs bool
    if p != nil && q != nil && p.Val == q.Val {
        ls = isSameTree(p.Left, q.Left)
        rs = isSameTree(p.Right, q.Right)
        return ls && rs
    }

    return false
}
