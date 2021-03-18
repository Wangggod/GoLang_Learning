package BTree

/**
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ResultType struct {
	SinglePath int
	MaxPath    int
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}
	left := helper(root.Left)
	right := helper(root.Right)
	result := ResultType{}
	if left.SinglePath > right.SinglePath {
		//把左边路的值加上去
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		//把右边路的值加上去
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
