package dynamic

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

//coinChange 找出给定数组需要最少的硬币
func CoinChange(coins []int, target int,mem map[int]int) int {
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
		subProblem := CoinChange(coins, target-coin,mem)
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
