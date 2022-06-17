package list

import "fmt"

//链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Last *DoubleNode
}

//makeListNormal ...
func makeListNormal() *ListNode {
	head := &ListNode{Val: 1} //形式1
	ln2 := &ListNode{Val: 2}
	ln3 := &ListNode{Val: 3}

	ln4 := &ListNode{ //形式2
		Val: 4,
	}

	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4

	return head

}

//makeDoubleNormal ...
func makeDoubleNormal() *DoubleNode {
	head := &DoubleNode{Val: 1} //形式1
	ln2 := &DoubleNode{Val: 2}
	ln3 := &DoubleNode{Val: 3}

	ln4 := &DoubleNode{ //形式2
		Val: 4,
	}

	head.Next = ln2
	head.Last = nil
	ln2.Next = ln3
	ln2.Last = head
	ln3.Next = ln4
	ln3.Last = ln2
	ln4.Last = ln3

	return head

}

//MakeListNode  根据数组生成链表
func MakeListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{Val: nums[0]}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return res
}

//Traversal  遍历链表
func Traversal(l1 *ListNode) {
	//这样遍历，l1指向nil
	for l1 != nil {
		fmt.Print(l1.Val, "  ")
		l1 = l1.Next
	}
	fmt.Println("\n--------------over--------------")
}

func Traverse(l1 *ListNode) {
	//var p *ListNode
	for p := l1; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
	fmt.Println("\n--------------over--------------")
}

//ReverseList 反转链表的实现
/**************************************************
*No206:给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
***************************************************/
func ReverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

/****************************************************
*No.19删除链表的倒数第N个结点
***************************************************/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	first := head
	second := dummy

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	ans := dummy.Next
	return ans
}

/**********************************************
*No24. 两两交换链表中的节点
*给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
*你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*head:1->2->3->4    result:2->1->4->3
********************************************/
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

/*************************************************************
*No21:合并两个有序链表
*将两个升序链表合并为一个新的 升序 链表并返回。
*新链表是通过拼接给定的两个链表的所有节点组成的。
*输入：l1 = [1,2,4], l2 = [1,3,4]
*输出：[1,1,2,3,4,4]
**************************************************************/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//引入哑结点
	head := &ListNode{Val: -1, Next: nil}
	prevHead := head
	//这里 preHead 和 head里面存的都是ListNode的地址
	//下面preHead移动，并不影响head的值
	fmt.Printf("%p\n", &prevHead)
	fmt.Println(prevHead)
	fmt.Println(*prevHead)
	fmt.Printf("%p\n", &head)
	fmt.Println(*head)
	fmt.Println(head)

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prevHead.Next = l1
			l1 = l1.Next
		} else {
			prevHead.Next = l2
			l2 = l2.Next
		}
		prevHead = prevHead.Next
	}

	if l1 == nil {
		prevHead.Next = l2
	} else {
		prevHead.Next = l1
	}

	return head.Next
}

/*************************************************************
*No2. 两数相加
*给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，
*并且每个节点只能存储 一位 数字。
*请你将两个数相加，并以相同形式返回一个表示和的链表。
*你可以假设除了数字 0 之外，这两个数都不会以 0开头。
**************************************************************/
func AddTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func TestAddTwoNum() {
	L1 := MakeListNode([]int{2, 4, 3})
	Traverse(L1)
	fmt.Println(L1)
	//如果不使用Traversal遍历，而是就地遍历，需要保存L1的值
	//遍历需要移动
	/*tmp := L1
	for tmp !=nil{
		fmt.Print(tmp.Val," ")
		tmp = tmp.Next
	}
	fmt.Println()
	fmt.Println("-----------------就地遍历----------------")

	L2 := MakeListNode([]int{5, 6, 4})
	Traversal(L2)

	result := AddTwoNumbers(L1, L2)
	Traversal(result)*/
}
