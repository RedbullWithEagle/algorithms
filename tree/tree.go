package tree

import (
	"fmt"
	"math"
)

//定义二叉树的节点
type Node struct {
	value int
	left  *Node
	right *Node
}

//功能：打印节点的值
//参数：nil
//返回值：nil
func (node *Node) Print() {
	fmt.Printf("%d ", node.value)
}

//功能：设置节点的值
//参数：节点的值
//返回值：nil
func (node *Node) SetValue(value int) {
	node.value = value
}

//功能：创建节点
//参数：节点的值
//返回值：nil
func CreateNode(value int) *Node {
	return &Node{value, nil, nil}
}

//功能：查找节点，利用递归进行查找
//参数：根节点，查找的值
//返回值：查找值所在节点
func (node *Node) FindNode(n *Node, x int) *Node {
	if n == nil {
		return nil
	} else if n.value == x {
		return n
	} else {
		p := node.FindNode(n.left, x)
		if p != nil {
			return p
		}
		return node.FindNode(n.right, x)
	}
}

//功能：求树的高度
//参数：根节点
//返回值：树的高度，树的高度=Max(左子树高度，右子树高度)+1
func (node *Node) GetTreeHeight(n *Node) int {
	if n == nil {
		return 0
	} else {
		lHeight := node.GetTreeHeight(n.left)
		rHeight := node.GetTreeHeight(n.right)
		if lHeight > rHeight {
			return lHeight + 1
		} else {
			return rHeight + 1
		}
	}
}

//功能：递归前序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.value)
		node.PreOrder(n.left)
		node.PreOrder(n.right)
	}
}

//功能：递归中序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) InOrder(n *Node) {
	if n != nil {
		node.InOrder(n.left)
		fmt.Printf("%d ", n.value)
		node.InOrder(n.right)
	}
}

//功能：递归后序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) PostOrder(n *Node) {
	if n != nil {
		node.PostOrder(n.left)
		node.PostOrder(n.right)
		fmt.Printf("%d ", n.value)
	}
}

//功能：打印所有的叶子节点
//参数：root
//返回值：nil
func (node *Node) GetLeafNode(n *Node) {
	if n != nil {
		if n.left == nil && n.right == nil {
			fmt.Printf("%d ", n.value)
		}
		node.GetLeafNode(n.left)
		node.GetLeafNode(n.right)
	}
}

func TestTree() {
	//创建一颗树
	root := CreateNode(1)
	//root.left = CreateNode(2)
	root.right = CreateNode(2)
	//root.left.right = CreateNode(7)
	//root.left.right.left = CreateNode(6)
	root.right.left = CreateNode(3)
	//root.right.right = CreateNode(9)

	//fmt.Printf("%d\n", root.FindNode(root, 4).value)
	//fmt.Printf("%d\n", root.GetTreeHeight(root))

	fmt.Println(numTrees(4))
	/*root.PreOrder(root)
	fmt.Printf("\n")
	root.InOrder(root)
	fmt.Printf("\n")
	root.PostOrder(root)
	fmt.Printf("\n")

	root.GetLeafNode(root)
	fmt.Printf("\n")*/
}

func inorderTraversal(root *Node) []int {
	result := make([]int, 0)
	traversal(root, &result)
	return result
}

func traversal(root *Node, result *[]int) {
	if root == nil {
		return
	}
	traversal(root.left, result)
	*result = append(*result, root.value)
	traversal(root.right, result)
}

/********************************************
*95. 不同的二叉搜索树 II
*给你一个整数 n ，
*请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。
*可以按 任意顺序 返回答案
********************************************/
func generateTrees(n int) []*Node {
	if n == 0 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) []*Node {
	if start > end {
		return []*Node{nil}
	}

	allTrees := []*Node{}
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &Node{i, nil, nil}
				currTree.left = left
				currTree.right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}

	return allTrees
}

/*
*No.97
 */
func numTrees(n int) int {
	treeCount := make([]int, n+1, n+1)
	treeCount[0] = 1
	treeCount[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			left := j - 1
			right := i - j
			treeCount[i] += treeCount[left] * treeCount[right]
		}
	}
	return treeCount[n]
}

/*
*No.98
 */
func isValidBST(root *Node) bool {
	return helperValidBSt(root.left, math.MinInt32, math.MaxInt32)
}

func helperValidBSt(root *Node, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.value <= lower || root.value >= upper {
		return false
	}
	return helperValidBSt(root.left, lower, root.value) && helperValidBSt(root.right, root.value, upper)
}

