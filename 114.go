//leetcode114

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    var p *TreeNode = root
    for p != nil {
        if p.Left != nil {
            var pr *TreeNode = p.Left
            for pr.Right != nil {
                pr = pr.Right
            }
            pr.Right = p.Right
            p.Right = p.Left
            p.Left = nil
        }
        p = p.Right
    }
}



