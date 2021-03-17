### 二叉树
本文的内容主要梳理介绍Go如何实现二叉树的各种遍历方式，以及在lc中如何变种
```

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


```
#### 1、 二叉树遍历
二叉树的遍历可以分为三种

1. 前序（根-左-右）
2. 中序（左-根-右）
3. 后续（左-右-根）

根据根访问的时间决定是什么遍历，并且左子树都要先于右子树

而从代码层面上又可以分为递归遍历和非递归遍历

##### 1.1、前序、中序、后序递归
```
//前序遍历
func preorderTraversal(root *TreeNode) {
if root==nil{
return
}
// 先访问根再访问左右
fmt.Println(root.Val)
preorderTraversal(root.Left)
preorderTraversal(root.Right)
}
```
三种遍历首先都是先判断根节点是不是空的。后三句直接决定了是何种遍历，如上图代码中先打印了根，即先序，如果改成
```
    inorderTraversalrecursion(root.Left)
    fmt.Println(root.Val)
    inorderTraversalrecursion(root.Right)
```
则为中序。

##### 1.2、前序、中序、后序非递归
无论是那种，由于是非递归，往往需要借助辅助结构
栈的弹出语句为：`stack = stack[:len(stack)-1]`
访问栈顶语句为：`node:= stack[len(stack)-1]`
先上代码：
```
func preorderTraversal(root *TreeNode) []int {
// 非递归
if root == nil{
return nil
}
result:=make([]int,0)
stack:=make([]*TreeNode,0)
for root!=nil || len(stack)!=0{
for root !=nil{
// 前序遍历，所以先保存结果
result=append(result,root.Val)
stack=append(stack,root)
root=root.Left
}
// pop
node:=stack[len(stack)-1]
stack=stack[:len(stack)-1]
root=node.Right
}
return result
}
```
此为前序非递归。首先照例要判空根节点，然后再后续操作。先声明一个辅助数组`result:=make([]int,0)`，用于存放每个节点的值。再声明`stack:=make([]*TreeNode,0)`用于暂存节点的位置。
由于是前序，所以每次访问到一个存在的根节点后就将其压入栈底，因为我们知道前序会一路沿着左路直接到底，因此越靠近底层的节点会越晚入栈，同时也会优先出栈。

我们利用for循环并且引入判空条件确保能一路走到底，每次访问到非空节点便`result=append(result,root.Val)`，因为是前序，所以可以直接把值输入到result中。但此时我们并不知道这个节点是否还有左右孩子，但暂时不管，所以立即调用`stack=append(stack,root)`暂存节点。最后使用`root=root.Left`继续往左。

当一路向左到底之后，我们建立一个辅助节点node，令`node:=stack[len(stack)-1]`，即弹出栈顶元素，然后检查其是否有右节点`root=node.Right`，注意弹出之后要更新stack，`stack=stack[:len(stack)-1]`。

那么中序非递归该如何实现呢？原理很简单，只是调整一下result的赋值位置，即需要先走到底，才开始赋值给result：
```
//向左语段
node:= stack[len(stack)-1]
stack = stack[:len(stack)-1]
result = append(result, node.Val)
root = val.Right
```
后续则稍微有些区别，如果直接在上述模板进行更改，其实是无法判断右节点是否已经弹出。
```
unc postorderTraversal(root *TreeNode) []int {
// 通过lastVisit标识右子节点是否已经弹出
if root == nil {
return nil
}
result := make([]int, 0)
stack := make([]*TreeNode, 0)
var lastVisit *TreeNode
for root != nil || len(stack) != 0 {
for root != nil {
stack = append(stack, root)
root = root.Left
}
// 这里先看看，先不弹出
node:= stack[len(stack)-1]
// 根节点必须在右节点弹出之后，再弹出
if node.Right == nil || node.Right == lastVisit {
stack = stack[:len(stack)-1] // pop
result = append(result, node.Val)
// 标记当前这个节点已经弹出过
lastVisit = node
} else {
root = node.Right
}
}
return result
}
```
##### 1.3、DFS（深度优先）
深度优先其实在二叉树中行为和先序是一样的，即先走到底，在往回看。因此有：
```
func preorderTraversal(root *TreeNode) []int {
result := make([]int, 0)
dfs(root, &result)
return result
}
// V1：深度遍历，结果指针作为参数传入到函数内部
func dfs(root *TreeNode, result *[]int) {
if root == nil {
return
}
*result = append(*result, root.Val)
dfs(root.Left, result)
dfs(root.Right, result)
}
```
##### 1.4、BFS
```
func levelOrder(root *TreeNode) [][]int {
// 通过上一层的长度确定下一层的元素
result := make([][]int, 0)
if root == nil {
return result
}
queue := make([]*TreeNode, 0)
queue = append(queue, root)
for len(queue) > 0 {
list := make([]int, 0)
// 为什么要取length？
// 记录当前层有多少元素（遍历当前层，再添加下一层）
l := len(queue)
for i := 0; i < l; i++ {
// 出队列
level := queue[0]
queue = queue[1:]
list = append(list, level.Val)
if level.Left != nil {
queue = append(queue, level.Left)
}
if level.Right != nil {
queue = append(queue, level.Right)
}
}
result = append(result, list)
}
return result
}
```

#### 2、分治法
分治思想即左右分开操作，然后一步一步合并回来，例如DFS也可以使用分治思想。

##### 2.1、分治法模板
```
func traversal(root *TreeNode) ResultType {
// nil or leaf
if root == nil {
// do something and return
}
// Divide
ResultType left = traversal(root.Left)
ResultType right = traversal(root.Right)
// Conquer
ResultType result = Merge from left and right
return result
}
```
##### 2.2、分治法应用：二叉树遍历
```
func preorderTraversal(root *TreeNode) []int {
result := divideAndConquer(root)
return result
}
func divideAndConquer(root *TreeNode) []int {
result := make([]int, 0)
// 返回条件(null & leaf)
if root == nil {
return result
}
// 分治(Divide)
left := divideAndConquer(root.Left)
right := divideAndConquer(root.Right)
// 合并结果(Conquer)
result = append(result, root.Val)
result = append(result, left...)
result = append(result, right...)
return result
}
```
##### 2.3、分治法应用：归并排序
```

func MergeSort(nums []int) []int {
return mergeSort(nums)
}
func mergeSort(nums []int) []int {
if len(nums) <= 1 {
return nums
}
// 分治法：divide 分为两段
mid := len(nums) / 2
left := mergeSort(nums[:mid])
right := mergeSort(nums[mid:])
// 合并两段数据
result := merge(left, right)
return result
}
func merge(left, right []int) (result []int) {
// 两边数组合并游标
l := 0
r := 0
// 注意不能越界
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

```

##### 2.4、分治法应用：快排
```

func QuickSort(nums []int) []int {
// 思路：把一个数组分为左右两段，左段小于右段，类似分治法没有合并过程
quickSort(nums, 0, len(nums)-1)
return nums
}
// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
if start < end {
// 分治法：divide
pivot := partition(nums, start, end)
quickSort(nums, 0, pivot-1)
quickSort(nums, pivot+1, end)
}
}
// 分区
func partition(nums []int, start, end int) int {
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
func swap(nums []int, i, j int) {
t := nums[i]
nums[i] = nums[j]
nums[j] = t
}

```