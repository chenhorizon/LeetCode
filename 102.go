// leetcode 102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

    if root == nil {
        return [][]int{}
    }

    ret := make(map[int][]*TreeNode)
    if root != nil {
        // ret = append(ret, []*TreeNode{root})
        ret[0] = []*TreeNode{root}
    }
    i := 1
    for len(ret[i-1]) > 0 {
        for _, node := range(ret[i-1]) {
            if node.Left != nil {
                ret[i] = append(ret[i], node.Left)
            }
            if node.Right != nil {
                ret[i] = append(ret[i], node.Right)
            }
        }
        i++
    }

    r := [][]int{}
    for i:=0; i<len(ret); i++ {
        item := make([]int, len(ret[i]))
        for j:=0; j<len(ret[i]); j++ {
            item[j] = ret[i][j].Val
        }
        r = append(r, item)
    }

    return r
}
