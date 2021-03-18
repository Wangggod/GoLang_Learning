package BTree

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//利用分治法解决问题
//强烈建议看完这题直接看balanced-binary-tree，是一套思路的
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//maxDepth函数返回值会int（题目要求）
func maxDepth(root *TreeNode) int {
	//首先还是要判空
	if root == nil {
		return 0 //这句很重要，即返回0可以看做定义了变量count计数
	}
	left := maxDepth(root.Left)   //沿着左边继续
	right := maxDepth(root.Right) //沿着右边继续
	//这边强烈建议画图
	if left > right {
		return left + 1
	}
	return right + 1
}
