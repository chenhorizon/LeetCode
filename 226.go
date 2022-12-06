// leetcode 226

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    if root.Left == nil && root.Right == nil {
        return root
    }

    ln := root.Left
    rn := root.Right

    root.Left = rn
    root.Right = ln

    if rn != nil {
        invertTree(root.Left)
    }
    if ln != nil {
        invertTree(root.Right)
    }

    return root
}
