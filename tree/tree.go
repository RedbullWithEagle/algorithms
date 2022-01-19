package tree

import (
	"fmt"
	"math"
)

//定义二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//功能：打印节点的值
//参数：nil
//返回值：nil
func (node *TreeNode) Print() {
	fmt.Printf("%d ", node.Val)
}

//功能：设置节点的值
//参数：节点的值
//返回值：nil
func (node *TreeNode) SetValue(value int) {
	node.Val = value
}

//功能：创建节点
//参数：节点的值
//返回值：nil
func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

//功能：查找节点，利用递归进行查找
//参数：根节点，查找的值
//返回值：查找值所在节点
func (node *TreeNode) FindTreeNode(n *TreeNode, x int) *TreeNode {
	if n == nil {
		return nil
	} else if n.Val == x {
		return n
	} else {
		p := node.FindTreeNode(n.Left, x)
		if p != nil {
			return p
		}
		return node.FindTreeNode(n.Right, x)
	}
}

//功能：求树的高度
//参数：根节点
//返回值：树的高度，树的高度=Max(左子树高度，右子树高度)+1
func (node *TreeNode) GetTreeHeight(n *TreeNode) int {
	if n == nil {
		return 0
	} else {
		lHeight := node.GetTreeHeight(n.Left)
		rHeight := node.GetTreeHeight(n.Right)
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
func (node *TreeNode) PreOrder(n *TreeNode) {
	if n != nil {
		fmt.Printf("%d ", n.Val)
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

//功能：递归中序遍历二叉树
//参数：根节点
//返回值：nil
func (node *TreeNode) InOrder(n *TreeNode) {
	if n != nil {
		node.InOrder(n.Left)
		fmt.Printf("%d ", n.Val)
		node.InOrder(n.Right)
	}
}

//功能：递归后序遍历二叉树
//参数：根节点
//返回值：nil
func (node *TreeNode) PostOrder(n *TreeNode) {
	if n != nil {
		node.PostOrder(n.Left)
		node.PostOrder(n.Right)
		fmt.Printf("%d ", n.Val)
	}
}

//功能：打印所有的叶子节点
//参数：root
//返回值：nil
func (node *TreeNode) GetLeafTreeNode(n *TreeNode) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			fmt.Printf("%d ", n.Val)
		}
		node.GetLeafTreeNode(n.Left)
		node.GetLeafTreeNode(n.Right)
	}
}

func TestTree() {
	//创建一颗树
	root := CreateTreeNode(3)
	root.Left = CreateTreeNode(9)
	root.Right = CreateTreeNode(20)
	root.Right.Left = CreateTreeNode(15)
	root.Right.Right = CreateTreeNode(7)

	levelOrder2(root)
	//root.Right.Right = CreateTreeNode(9)

	//fmt.Printf("%d\n", root.FindTreeNode(root, 4).Val)
	//fmt.Printf("%d\n", root.GetTreeHeight(root))

	//fmt.Println(numTrees(4))
	/*root.PreOrder(root)
	fmt.Printf("\n")
	root.InOrder(root)
	fmt.Printf("\n")
	root.PostOrder(root)
	fmt.Printf("\n")

	root.GetLeafTreeNode(root)
	fmt.Printf("\n")*/
}

/**********************************************************
*No.94 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
*方法一：递归
**********************************************************/
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)
	return result
}

//result *类型，外面可以看到函数内的修改
func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, result)
	*result = append(*result, root.Val)
	traversal(root.Right, result)
}

/******************************************************
*方法2：迭代
*******************************************************/
func inorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

/********************************************
*95. 不同的二叉搜索树 II
*给你一个整数 n ，
*请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。
*可以按 任意顺序 返回答案
*方法：回溯
********************************************/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	allTrees := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}

	return allTrees
}

/****************************************************
*No.96 不同的二叉搜索树
*给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？
*返回满足题意的二叉搜索树的种数。
***************************************************/
func numTrees(n int) int {
	treeCount := make([]int, n+1, n+1)
	treeCount[0] = 1
	treeCount[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			Left := j - 1
			right := i - j
			treeCount[i] += treeCount[Left] * treeCount[right]
		}
	}
	return treeCount[n]
}

/****************************************************************
*No.98 验证二叉搜索树
*给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

*有效 二叉搜索树定义如下：

*节点的左子树只包含 小于 当前节点的数。
*节点的右子树只包含 大于 当前节点的数。
*所有左子树和右子树自身必须也是二叉搜索树。

*方法一：递归
****************************************************************/
func isValidBST(root *TreeNode) bool {
	return helperValidBSt(root.Left, math.MinInt32, math.MaxInt32)
}

func helperValidBSt(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helperValidBSt(root.Left, lower, root.Val) && helperValidBSt(root.Right, root.Val, upper)
}

/****************************************************************
* 方法二：中序遍历
****************************************************************/
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}

	return true
}

/*****************************************************************
*No.99. 恢复二叉搜索树
*给你二叉搜索树的根节点 root ，该树中的两个节点的值被错误地交换。
*请在不改变其结构的情况下，恢复这棵树。
*****************************************************************/
func recoverTree(root *TreeNode) {
	//1.先中序遍历
	nums := inorderTraversal2(root)
	//2.在数组中找出需要交换的两个节点
	x, y := findTwoSwapped(nums)
	//3.交换这两个节点
	recover(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	x, y := nums[index1], nums[index2]
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}

	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}

		count--
		if count == 0 {
			return
		}
	}

	recover(root.Right, count, x, y)
	recover(root.Left, count, x, y)
}

/**************************************************************
*No.100 相同的树
*给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
*如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
***************************************************************/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right)
}

/**************************************************************
*No.101. 对称二叉树
*给你一个二叉树的根节点 root ， 检查它是否轴对称。
***************************************************************/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if !checkSymmetric(p.Left, q.Right) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Left)
}

/**************************************************************
*No.102. 二叉树的层序遍历
*给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
（即逐层地，从左到右访问所有节点）。
***************************************************************/
func levelOrder(root *TreeNode) [][]int {
	nums := make([][]int, 0)
	helperLevelOrder(root, 0, &nums)
	return nums
}

//这个算法需要注意二维数组的使用
func helperLevelOrder(root *TreeNode, depth int, nums *[][]int) {
	if root == nil {
		return
	}
	//这里先生成一维数组
	if len(*nums) < depth+1 {
		*nums = append(*nums, []int{})
	}
	(*nums)[depth] = append((*nums)[depth], root.Val)

	depth += 1
	helperLevelOrder(root.Left, depth, nums)
	helperLevelOrder(root.Right, depth, nums)
}

func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

/**************************************************************
*No.103. 二叉树的锯齿形层序遍历
*给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
*（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）
***************************************************************/
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}

		q = p
		if i%2 == 1 {
			for m, n := 0, len(ret[i]); m < n/2; m++ {
				ret[i][m], ret[i][n-1-m] = ret[i][n-1-m], ret[i][m]
			}
		}
	}
	return ret
}

/************************************************************
*No.104 二叉树的最大深度
*给定一个二叉树，找出其最大深度。
************************************************************/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if rightDepth > leftDepth {
		return rightDepth + 1
	}

	return leftDepth + 1
}

/************************************************************
*No.111. 二叉树的最小深度
*给定一个二叉树，找出其最小深度
*最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
************************************************************/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt32
	if root.Left != nil {
		leftDepth := minDepth(root.Left)
		if leftDepth < minD {
			minD = leftDepth
		}
	}

	if root.Right != nil {
		rightDepth := minDepth(root.Right)
		if rightDepth < minD {
			minD = rightDepth
		}
	}

	return minD + 1
}
