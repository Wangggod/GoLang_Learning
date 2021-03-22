package PostorderTraversal

/**
题目：145
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一：普通递归
/**
这里有个小细节是res的声明应该在函数外还是函数内？
应该是在函数内，实际上递归的精髓就在于每次都建一个小块，然后一块一块合并
*/
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

//方法二：拆开的递归
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}

//方法三：非递归

func postorderTraversal3(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	if root == nil {
		return res
	}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return res
}
