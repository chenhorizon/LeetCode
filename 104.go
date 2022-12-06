// leetcode 104

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }

    lh, rh := 1, 1
    if root.Left != nil {
        lh = lh + maxDepth(root.Left)
    }
    if root.Right != nil {
        rh = rh + maxDepth(root.Right)
    }

    if lh > rh {
        return lh
    }
    return rh
}
