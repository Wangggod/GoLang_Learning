package recursion

/**
在day1和day2中我们接触了不少关于递归的思路，但毕竟年轻，总是会不自觉陷入递归旋涡中
因此今天我们试图通过一些巩固练习，加深对递归的理解。
递归说到底就是大问题转换成小问题，我们不再从全局的角度思考，而是拆分成小问题
*/

/**
首先我们来写一段经典的前序递归
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

/**
递归，就是反复调用自己。它有连个最重要的组成：1、终止条件 2、做什么 3、等价关系式
递归调用总是去尝试解决一个规模更小的问题。
上述代码看起来比较不直观，我们可以稍加修改让它长得更像递归
*/

func preorderTraversal2(root *TreeNode) []int {
	var result []int
	preOrder(root, &result)
	return result
}
func preOrder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preOrder(root.Left, output)
		preOrder(root.Right, output)
	}
}

/**
可以看到递归就是把大问题拆成小问题，基于上述二叉树的遍历方式，我们稍加修改得到新的题目：
104：二叉树的最大深度
如果我们从根出发，要知道一颗二叉树有多高等价于根节点往下还有几层，再加上根节点本身的高度即可。
例如我们知道二叉树不为空，则至少有1层，根节点往下还有几层？不知道，但我们可以假设为有N层
那么是不是总层数=N+1？按照这个思维往下推，实际上我们最终就是要从最底层开始，一点一点往上加
*/

func maxDepth(root *TreeNode) int {
	//我们首先要思考递归退出的条件是什么
	//在本题中就是遍历到了叶子节点就退出
	if root == nil {
		return 0
		//这句话的理解方式有两个层面，第一个层面是如果根不存在，那么高度就是0.
		//进一步，如果一个节点是空，那么它这层就不存在，所以为0.
		//递归的核心思路就是从小问题逐级累积起来
	}
	//接着我们要考虑如果不是空层的时候做什么
	//我们可以按照前序遍历的方式先走到底
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	//left和right在小问题中就是一个节点的左右高度，假设树只有根和左孩子。那么左孩子的left和right返回值都是0
	//因此此时左孩子已经是最后一层了。但左孩子是存在的，因此其高度为1.
	//如何才能返回得到这个1呢？答案是left和right中较大的那个+1，因为在递归返回中这个值就会变成左孩子的高度
	//失去了左右属性。
	//于是我们加入判断
	if left > right {
		return left + 1
	}
	return right + 1
}

/**
上题中我们把大问题拆分成了小问题，具体到每个小子树都高度上，再逐级返回得到了结果
110：平衡二叉树
平衡二叉树即所有子树的左右子树高度差都不超过1，现在我们基于上一题的思维可以很快想到
可以通过从最小的树开始累积值，如果最终汇集到根之后左右两边高度差小于等于1则成立。
*/
func isBalanced(root *TreeNode) bool {
	//由于给出的函数是返回bool，所以我们可以再建立一个函数负责判断高度
	if judge(root) == -1 {
		return false
	}
	return true
}
func judge(root *TreeNode) int {
	//这个函数的返回值为int，因为我们将会返回每个小问题的高度差
	//首先考虑退出机制，和上一题一样，当节点为空的时候，高度差就不存在了
	if root == nil {
		return 0
	}
	//接着我们要求两边的高度差，其实是借助两边的高度比较，因此实际上是求两边的高度
	left := judge(root.Left)
	right := judge(root.Right)
	if left-right > 1 || right-left > 1 || left == -1 || right == -1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

/**
有了上述基础之后，我们稍微来做一题难一些的】
98:验证搜索二叉树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
最关键的是那句“左右子树也是二叉搜索树”
那么在中序条件下，我们得到的就是递增的数列
所以我们可以直接到最底部的子树，然后看是否满足，满足的话就继续，不满足就结束
*/
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	helper(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}
func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root == nil {
		return
	}
	helper(root.Left, result)
	*result = append(*result, root.Val)
	helper(root.Right, result)
}
