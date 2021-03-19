### 链表
核心考点

* nil的处理
* dummy node（哑巴节点，用来控制头结点可能变化的题目）
* 快慢指针（可以找到链表的中间or末尾）
* 插入节点到排序链表
* 从链表中移除一个节点
* 翻转链表
* 合并链表
* 找到链表中间节点（就是快慢节点的用法）

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
```

#### 1、nil判定和dummy node
这一步操作几乎要放在所有代码段的最开头。

首先我们要明确概念是，一般的链表题head就代表了链表的第一个元素。因此我们要特别考虑第一个节点的可删除性。

比如删除多于的重复元素，只有一个元素的话或者是空的，直接返回就行了，因为第一个一定不会被删掉。

但比如删除所有含有重复元素的节点，那么head是可能被删掉的，因此我们要设置一个dummy node放在head前面，防止出现意外。

例如，普通链表 1(head）-> 2 -> 3 -> nil
我们可以一顿操作：
```
dummy := &ListNode{Val:0} // dummy（空节点）
dummy.Next = head // dummy -> head(1) -> 2 -> 3 -> nil
head = dummy // head -> 1 -> 2 -> 3 -> nil
```
这样一来head在初始状态下就不代表任何值，我们便可以通过head.Next访问到第一个节点的值。

例如删除重复的元素（至少保留一个）
```
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil {
        // 全部删除完再移动到下一个元素
        for current.Next != nil && current.Val == current.Next.Val {
            current.Next = current.Next.Next
        }
        current = current.Next
    }
    return head
}
```

例如删除所有含重复元素的节点（一个都不剩）
```
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        retrun head
    }
    dummy := &ListNode{Val:0}
    dummy.Next = head
    head = dummy
    var rmVal int
    for head.Next != nil && head.Next.Next != nil {
        if head.Next.Val ==  head.Next.Next.Val {
            rmVal = head.Next.Val           
           for head.Next != nil && head.Next.Val == rmVal {
                head.Next = head.Next.Next
           }
        } else {
            head = head.Next
        }
    }
    return dummy.Next
}
```

#### 2、翻转链表
