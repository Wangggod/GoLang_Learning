package BTree

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/

/**
在104中我们已经学会了用分治法寻找高度
这个过程中我们发现，实际上求深度是通过分治累积left和right
然后逐层合并起来的过程
其中递归深度最大时其实就是比较当前的子树左右大小，无非是左大、右大、等大
因此要求二叉平衡树，我们可以在递归最深处时进行一个比较
如果当前累积值以及不平衡了，那我们便可以终止
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//我们首先定义一个判断函数通过深度来判断是否为平衡树
func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

//在isBalanced中我们定义如果深度为-1则不是平衡树，因此maxDepth中我们要设置当不符合条件时深度自动为-1
func maxDepth(root *TreeNode) int {
	//首先进行正常的分治操作
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	//在返回值前，我们要判断是否出现了意料之外的情况
	/**
	当目前的左子树和右子树积累值（对应left和right）相减>1，已经不平衡了，则直接返回-1
	此外我们也可以想，如果left和right如果已经有一个已经是-1了，完全没有必要继续比，直接继续-1
	*/
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
