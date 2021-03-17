package BTree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序递归方法遍历二叉树
func preorderTraversalrecursion(root *TreeNode) {
	if root == nil {
		return
	} //首先判断根节点是否有值
	fmt.Println(root.Val)                  //访问根节点
	preorderTraversalrecursion(root.Left)  //访问左节点
	preorderTraversalrecursion(root.Right) //访问右节点
}

//前序非递归方法遍历二叉树
func preorderTraversalnormal(root *TreeNode) []int {
	if root == nil {
		return nil
	} //首先判断根节点是否有值
	result := make([]int, 0)      //result数组用于返回访问结果的顺序
	stack := make([]*TreeNode, 0) //stack用于访问时控制顺序
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val) //当访问到一个非空节点时，加入result，则在输出时它会优先处理
			stack = append(stack, root)       //每次往result加入非空节点的时候同时也加入stack
			root = root.Left                  //由于是前序，接下去将访问左节点（如果存在的话）
		}
		node := stack[len(stack)-1]  //在把某条左子树走到底之后，将会弹出节点
		stack = stack[:len(stack)-1] //开始检查是否有右节点
		root = node.Right
	}
	return result
}

//中序递归二叉树
func inorderTraversalrecursion(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalrecursion(root.Left)
	fmt.Println(root.Val)
	inorderTraversalrecursion(root.Right)
}

//中序非递归二叉树
func inorderTraversalnormal(root *TreeNode) []int {
	if root == nil {
		return nil
	} //首先判断根节点是否有值
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1] //把弹出的节点赋值给node
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

//后序递归遍历二叉树
func postorderTraversalrecursion(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversalrecursion(root.Left)
	postorderTraversalrecursion(root.Right)
	fmt.Println(root.Val)
}

//后序非递归遍历二叉树
func postorderTraversalnormal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVist *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVist {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVist = node
		} else {
			root = node.Right
		}
	}
	return result
}

//深度优先搜索-从上到下
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

//由于dfs的原理和前序一样，我们也可以借助dfs再次实现前序遍历
func preDfs(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

//BFS 层次遍历，需要借助队列
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0) //辅助队列，里面存放树节点
	queue = append(queue, root)   //首先往辅助队列里加入根节点
	for len(queue) > 0 {
		list := make([]int, 0) //辅助数组list，用于暂存节点值
		length := len(queue)   //length用于记录这一层有多少个元素，
		for i := 0; i < length; i++ {
			level := queue[0]              //首先将队列第一个元素出列，并且把出列元素给level
			queue = queue[1:]              //将队列更新
			list = append(list, level.Val) //先将出列元素值存入数组list中
			if level.Left != nil {
				queue = append(queue, level.Left) //如果这个元素有左节点，则加入下一层
			}
			if level.Right != nil {
				queue = append(queue, level.Right) //同理如果有右节点，也加入下一层
			}
		}
		result = append(result, list)
	}
	return result
}

//分治模板
func traversal(root *TreeNode) []int { //这里的[]int可以任意替换成目标类型ResultType
	result := make([]int, 0) //通常result用于存放节点值，具体是[]int根据ResultType而定
	if root == nil {
		return result //这句也可以任意替换，即当根节点为空时执行某些行为，比如返回或者操作
	}
	//分治开始
	left := traversal(root.Left)
	right := traversal(root.Right)

	//合并
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	//返回结果
	return result
}

//分治模板应:1：深度优先搜索-分治法
func divideAndConquer(root *TreeNode) []int { //例如此处二叉树节点的值为int，则使用[]int

	//标准操作，定义返回数组和节点为空时候的行为
	result := make([]int, 0)
	if root == nil {
		return result
	}
	//开始分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	//合并结果
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

//分治模板应用2：基于分治法我们也可以再次重写先序遍历
func preDiv(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

//分治模板应用3：归并排序
//详细链接参考：https://www.cnblogs.com/chengxiao/p/6194356.html 有图解更直观
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	//分治开始
	mid := len(nums) / 2           //首先将待排序的数组从中间截断分成两段
	left := mergeSort(nums[:mid])  //左边从0开始到mid继续分治
	right := mergeSort(nums[mid:]) //右边从mid开始到结束继续分治
	//合并结果
	result := merge(left, right) //当左右分到最小的时候开始合并，merge函数见下方
	return result
}
func merge(left, right []int) (result []int) {
	//两边数组合并的游标
	l := 0
	r := 0
	//注意判断有误越界
	//我们可以理解为理想状态下分治到最后的时候只剩下两个元素比较
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩余部分合并
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

//分治模板应用4：快排
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, start, end int) {
	if start < end {
		//分治开始
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int { //核心分区算法，将传入开始和结尾的位置
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 把中间的值换为用于比较的基准值
	swap(nums, i, end)
	return i
}
func swap(nums []int, i, j int) { //原地交换算法，需要用一个辅助变量t
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
