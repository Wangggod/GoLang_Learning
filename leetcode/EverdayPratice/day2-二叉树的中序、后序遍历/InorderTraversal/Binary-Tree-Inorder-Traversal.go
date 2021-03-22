package InorderTraversal

/**
题目：94
主要是回顾中序遍历的写法，形成肌肉记忆并且背诵
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一：递归法
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)
	return res
}

//方法二：拆开的递归法
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	inorder2(root, &res)
	return res
}

func inorder2(root *TreeNode, output *[]int) {
	if root != nil {
		inorder2(root.Left, output)
		*output = append(*output, root.Val)
		inorder2(root.Right, output)
	}
}

//方法三：非递归模拟
func inorderTraversal3(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}
