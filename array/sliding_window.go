package array

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

/***************************************************
*No76  最小覆盖子串
*例如：给定字符串s,一个字符串t，返回s中涵盖t所有字符的最小子串。如果不存在，返回NULL
*注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
*输入：s = "ADOBECODEBANC", t = "ABC"      输出："BANC
*1.t是否有重复字符？   按照有重复字符
*****************************************************/
func GetMinString(source, target string) string {
	lenSource := len(source)
	lenTarget := len(target)
	min := lenSource + 1
	minStr := ""

	for i := 0; i < lenSource; i++ {
		for j := i + 1; j < lenSource+1; j++ {
			findCount := 0
			//判断target中的每个元素 是否包含在source[i:j]中
			for _, value := range target {
				if !strings.Contains(source[i:j], string(value)) {
					break
				}
				findCount++
			}

			if findCount == lenTarget {
				findMin := j - i
				if findMin < min {
					min = findMin
					minStr = source[i:j]
				}
			}
		}
	}

	if min > 0 && min <= lenSource {
		return minStr
	}

	return ""
}

const INT_MAX = 100000

//滑动窗口解决
func minWindowMy(s, t string) string {
	lenSource := len(s)
	lenTarget := len(t)
	fmt.Println(lenSource, lenTarget)
	fmt.Println("---------------------------")
	need := make(map[string]int)
	window := make(map[string]int)

	for _, value := range t {
		need[string(value)]++
	}

	var left, right int
	var valid int
	var start int
	len := INT_MAX
	for ; right < lenSource; {
		c := s[right]
		right++
		fmt.Println(string(c))
		fmt.Printf("window:[%d,%d]\n", left, right)
	}

	fmt.Println("------------------")
	fmt.Println(valid, start, len)
	fmt.Println(window)
	return ""
}

func minWindow(s string, t string) string {
	//ori 给定target中的字符个数
	//cnt 滑动窗口中在target中的元素的个数
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	//判断cnt中是否包含了t中的所有元素
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

//力扣官方优化版
func minWindowOptimization(s string, t string) string {
	lenSource := len(s)
	lenTarget := len(t)
	//异常判断
	if lenSource == 0 || lenTarget == 0 || lenSource < lenTarget {
		return ""
	}

	mapWindow := make(map[byte]int, 128)
	mapTarget := make(map[byte]int, 128)

	for _, v := range t {
		mapTarget[byte(v)]++
	}

	var distance, begin int
	minLen := lenSource + 1

	var left, right int
	for ; right < lenSource; {
		if mapTarget[s[right]] == 0 {
			right++
			continue
		}

		//维护distance逻辑
		if mapWindow[s[right]] < mapTarget[s[right]] {
			distance++
		}
		mapWindow[s[right]]++
		right++

		for ; distance == lenTarget; {
			if right-left < minLen {
				minLen = right - left
				begin = left
			}

			if mapTarget[s[left]] == 0 {
				left++
				continue
			}

			//维护distance逻辑
			if mapWindow[s[left]] == mapTarget[s[left]] {
				distance--
			}
			mapWindow[s[left]]--
			left++
		}
	}

	if minLen == lenSource+1 {
		return ""
	}
	return s[begin : begin+minLen]
}

/***************************************************
*No3  无重复字符的最长子串
*给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*****************************************************/
func lengthOfLongestSubstring(s string) int {
	lenS := len(s)
	if lenS <= 1 {
		return lenS
	}

	mapWindow := make(map[string]int)
	var left, right int
	minLen := 0

	for ; right < lenS; {
		if _, ok := mapWindow[string(s[right])]; !ok {
			mapWindow[string(s[right])]++
			if len(mapWindow) > minLen {
				minLen = len(mapWindow)
			}
			right++
			continue
		}

		//如果在里面
		for ; left < right; {
			delete(mapWindow, string(s[left]))
			if s[left] == s[right] {
				left++
				break
			}
			left++
		}
	}

	if minLen == lenS+1 {
		return 0
	}

	return minLen
}

func TestMinStr() {
	source := "adbecfebac"
	//target := "ABCC"
	//fmt.Println(GetMinString(source,target))
	//fmt.Println(minWindowOptimization(source, target))
	//minWindow(source,target)

	fmt.Println(lengthOfLongestSubstring(source))
}

/***************************************************
*No7  整数反转
*给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
*如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。
*****************************************************/
func reverse(x int) int {
	rev := 0
	for x != 0 {
		//这里注意，为什么要/10，防止越界
		if rev > math.MaxInt32/10 || rev < math.MinInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}

	return rev
}

/***************************************************
*No5  最长回文子串
*给你一个字符串 s，找到 s 中最长的回文子串。
*输入：s = "babad"
*输出："bab"
*解释："aba" 同样是符合题意的答案。
*****************************************************/
func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	} else if sLen == 1 {
		return s
	}

	begin := -1
	maxLen := 0

	check := func(tmp string) bool {
		left := 0
		right := len(tmp) - 1
		for left < right {
			if tmp[left] != tmp[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		//这里j是小于==  ，否则“bb”这种，不通过
		for j := i + 1; j <= len(s); j++ {
			if check(s[i:j]) {
				if j-i > maxLen {
					maxLen = j - i
					begin = i
				}
			}
		}
	}

	if maxLen == 0 {
		return ""
	}

	return s[begin : begin+maxLen]
}

//CheckHW 检测一个字符串是否是回文
func CheckHW(tmp string) bool {
	left := 0
	right := len(tmp) - 1
	for left < right {
		if tmp[left] != tmp[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*************************************
*No9 判断一个数字是否是回文数
************************************/
func checkIntHW(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

/*****************************************
*No6  Z 字形变换
*将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
*比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：
*P   A   H   N
*A P L S I I G
*Y   I   R
*之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
******************************************/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	sLen := len(s)
	cycleLen := numRows + numRows - 2
	var ret string

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < sLen; j += cycleLen {
			ret += string(s[i+j])
			if i != 0 && i != numRows-1 && j+cycleLen-i < sLen {
				ret += string(s[j+cycleLen-i])
			}
		}
	}

	return ret
}

/*****************************************
No6 符串转换整数 (atoi)
请你来实现一个myAtoi(string s)函数，
使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
函数 myAtoi(string s) 的算法如下：
读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。
确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。
字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。
如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，
需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，
大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：
本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
******************************************/
func MyAtoi(s string) int {
	var ret int
	isDivide := false
	isPlus := false
	isNum := false
	for _, v := range s {
		switch v {
		case 32: //" "
			if ret == 0 {
				continue
			} else {
				break
			}
		case 43:
			if isDivide || isPlus || isNum {
				return 0
			}
			isPlus = true
		case 45: //"-"
			if isPlus || isDivide || isNum {
				return 0
			}

			isDivide = true
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			isNum = true
			if ret > (math.MaxInt32/10 - (int(v) - 48)) {
				return math.MaxInt32
			}

			if ret < (math.MinInt32/10 - (int(v) - 48)) {
				return math.MinInt32
			}

			tmp := int(v) - 48
			if ret < 0 {
				ret = ret*10 - (tmp)
			} else {
				ret = ret*10 + (tmp)
				if ret < 10 && ret > 0 && isDivide {
					ret = -ret
				}
			}
		default:
			if ret == 0 {
				return 0
			} else {
				return ret
			}
		}
	}

	return ret
}

func myAtoi2(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs

}

/**************************************************
No13  罗马数字转整数
*************************************************/
func romanToInt(s string) int {
	mapRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	ret := 0
	n := len(s)
	for i := range s {
		value := mapRoman[s[i]]
		if i < n-1 && value < mapRoman[s[i+1]] {
			ret -= value
		} else {
			ret += value
		}
	}

	return ret
}

func romanToIntOptimize(s string) int {
	mapRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	ret := 0
	n := len(s)
	for i := range s {
		value := mapRoman[s[i]]
		if i < n-1 && value < mapRoman[s[i+1]] {
			ret -= value
		} else {
			ret += value
		}
	}

	return ret
}
/*********************************************************
No.20有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
**********************************************************/
func isValid(s string) bool {
	mapSign := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	stack := list.New()

	for _, v := range s {
		if _, ok := mapSign[byte(v)]; ok {
			stack.PushBack(byte(v))
		} else {
			tmp := stack.Back()
			if tmp == nil {
				return false
			}
			if mapSign[tmp.Value.(byte)] != byte(v) {
				return false
			}
			stack.Remove(tmp)
		}
	}

	if stack.Len() > 0 {
		return false
	}
	return true
}

func TestReverse() {
	strSign := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, v := range strSign {
		tmp := isValid(v)
		fmt.Println(v, " result is :", tmp)
	}
	str := "MCMXCIV"
	fmt.Println(romanToInt(str))

	//fmt.Println(romanToInt("MCMXCIV"))
	//strABC := "00000-42a1234"
	//fmt.Println(MyAtoi(strABC))
	/*strSource := "PAYPALISHIRING"
	fmt.Println(convert(strSource, 3))
	str := []string{"babad", "cbbd", "aa", "ac", "ababa"}
	for _, v := range str {
		result := longestPalindrome(v)
		fmt.Println(v, " is ", result)
	}*/

	//测试反转数字
	/*a1:=1534236469
	  fmt.Println(reverse(a1))

	  a2:= -3435634
	  fmt.Println(reverse(a2))*/
}
