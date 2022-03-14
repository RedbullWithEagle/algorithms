package dynamic

import "fmt"

//fib 斐波那契数列  递归算法
func fib(target int) int {
	if target == 0 {
		return 0
	}

	if target == 1 || target == 2 {
		return 1
	}

	return fib(target-1) + fib(target-2)
}

//fibDynamic 斐波那契数列  动态规划版本
func fibDynamic(target int) int {
	if target == 0 {
		return 0
	}

	if target == 1 || target == 2 {
		return 1
	}

	prev := 1
	curr := 1

	// 采用交换不占用额外空间
	for i := 3; i < target; i++ {
		sum := prev + curr
		prev, curr = curr, sum
	}
	return curr
}

//coinChange 找出给定数组需要最少的硬币
func CoinChange(coins []int, target int, mem map[int]int) int {
	if target == 0 {
		return 0
	}

	if target < 0 {
		return -1
	}

	if v, ok := mem[target]; ok {
		return v
	}

	res := target + 1
	for _, coin := range coins {
		subProblem := CoinChange(coins, target-coin, mem)
		if subProblem == -1 {
			continue
		}

		if res > subProblem+1 {
			res = subProblem + 1
		}
	}

	if res == target+1 {
		mem[target] = -1
		return -1
	}

	mem[target] = res
	return res
}

var mem map[int]int

func TestCoin() {
	//硬币
	/*coins := []int{1, 2, 5}

	mem = make(map[int]int)

	fmt.Println(CoinChange(coins, 11, mem))*/

	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
}

/*********************************************
*No.53 最大子数组和
*给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*子数组 是数组中的一个连续部分。
*
********************************************/
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

/*********************************************
*No.70 爬楼梯
*假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
*每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
********************************************/
func climbStairs(n int) int {
	mem := make([]int, n+1, n+1)
	if n == 0 {
		return 0
	}
	mem[0] = 1
	mem[1] = 1

	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}
