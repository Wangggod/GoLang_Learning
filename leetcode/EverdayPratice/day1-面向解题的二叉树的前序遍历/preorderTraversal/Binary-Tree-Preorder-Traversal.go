package preorderTraversal

//题目 144. 二叉树的前序遍历
/**
本文件将尽可能列举出二叉树的前序遍历模板，建议背诵。
在复习数据结构的过程中，我们所写的前序遍历都是直接输出结果的类型
但实际上在LeetCode写题的过程中我们发现都是要你返回数组。
借此抛转引玉，在复习二叉树代码熟练度的同时顺便规范如何正确返回值给测试单元

*/
//定义二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一：普通递归
/**
首先我们看一段代码：
func preorderTraversal(root *TreeNode)  {
    if root==nil{
        return
    }
    // 先访问根再访问左右
    fmt.Println(root.Val)
    preorderTraversal(root.Left)
    preorderTraversal(root.Right)
}
这段代码就是上述的直接输出节点型，但题目要求我们要[]int，因此我们稍加修改
*/
func preorderTraversal1(root *TreeNode) []int {
	/**
	函数末端声明了[]int，因此函数中必然要return ResultType，但如果是直接声明：
	func preorderTraversal(root *TreeNode) (vals []int){ }
	这就表明函数会返回vals，vals已经为[]int，所以我们只需要写return而无需声明类型
	*/
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	//其实这个方法有一丝丝分治法的意思，即左右各自操作结束后合并。
	res = append(res, preorderTraversal1(root.Left)...) //即每次都把res所有元素都添加进末尾
	res = append(res, preorderTraversal1(root.Right)...)
	return res
}

//方法二：基于DFS的递归
/**
普通递归是基于最经典的教科书递归方法构造的，但我们可以稍微进行延伸，利用DFS的思想来完成这个过程
DFS和分治法的区别是，DFS一般将最终结果作为指针参数传入，分治法（方法一的思想）则是在递归中作为返回结果最后合并
*/
func preorderTraversal2(root *TreeNode) []int {
	var res []int   //由于存放返回的结果
	dfs(root, &res) //这边体现了DFS的”将结果作为指针参数传入“
	return res
}
func dfs(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		dfs(root.Left, output)
		dfs(root.Right, output)
	}
}

/**
方法二其实还有一种写法：
var ans []int
func preorderTraversal(root *TreeNode) []int {
	ans = ans[:0]
	handler(root)
	return ans

}
func handler(root *TreeNode) {
	if root != nil {
		ans = append(ans, root.Val)
		handler(root.Left)
		handler(root.Right)
	}
}
*/

//方法三：还是递归
/**
这个方法其实是第方法一和方法二的结合体，第一个方法是不在函数中声明返回值
而这个方法是直接在函数中声明了要返回val
*/
func preorderTraversal3(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)      //将preorder作为函数变量
	preorder = func(node *TreeNode) { //定义了函数变量的做法
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

//方法四：迭代法
/**
该方法即通过栈的方式用来暂存节点，在逐级返回。
*/
func preorderTraversal4(root *TreeNode) []int {
	res := make([]int, 0)         //可以用var res []int 声明
	stack := make([]*TreeNode, 0) //这句也可以直接用var stack []*TreeNode声明
	if root != nil {
		return res
	}
	/**接下来我们要把节点暂存入stack，把得到的值存入res，因此需要借助循环
	因此制约循环结束的条件是stack中没有值
	*/
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	//最终返回结果
	return res
}

//方法五：Morris遍历
func preorderTraversal5(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
