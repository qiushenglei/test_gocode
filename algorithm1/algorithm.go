package algorithm1

import (
	"sort"
)

func SubtractProductAndSum(n int) int {

	if n <= 0 {
		return 0
	}

	add := 0
	mul := 1

	for n > 0 {
		num := n % 10
		n = n / 10
		add += num
		mul *= num
	}

	return mul - add
}

func reverseString(s []byte) []byte {
	i, j := 0, len(s)-1
	if i == j {
		return nil
	}

	for i < j {
		s[j], s[i] = s[i], s[j]
		i++
		j--
	}
	return s
}

func reverseStr(s string, k int) string {
	//bytes := []byte(s)
	//lengths := len(s)
	//if lengths == 0 {
	//	return s
	//}

	//if
	return ""
}

func ReplaceSpace(s string) string {
	space := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			space++
		}
	}

	bytes := []byte(s)
	resize := make([]byte, 2*space)
	bytes = append(bytes, resize...)

	// 快慢指针
	slow, fast := l-1, l+2*space-1

	for slow >= 0 {
		if s[slow] == ' ' {
			bytes[fast] = '0'
			fast--
			bytes[fast] = '2'
			fast--
			bytes[fast] = '%'
		} else {
			bytes[fast] = bytes[slow]
		}
		slow--
		fast--
	}

	return string(bytes)
}

func ReverseWords(s string) string {
	b := []byte(s)
	reverseString(b)

	for start, i := 0, 0; i < len(b); i++ {
		if b[i] == ' ' {
			reverseString(b[start:i])
		}
	}
	return string(b)
}

func binSearch(left, right, target int, nums []int) int {

	if left > right {
		return -1
	}

	mid := (right-left)/2 + left

	res := -1
	if target == nums[mid] {
		return mid
	} else if target < nums[mid] {
		right = mid - 1
	} else {
		left = mid + 1
	}
	res = binSearch(left, right, target, nums)
	return res
}

func Search(nums []int, target int) int {
	l := len(nums)
	res := binSearch(0, l-1, target, nums)
	return res
}

func RemoveElement(nums []int, val int) int {

	s, f := 0, 0
	for ; f < len(nums); f++ {
		if nums[f] != val {
			nums[s] = nums[f]
			s++
		}
	}
	return len(nums[0:s])
}

func BackspaceCompare(s string, t string) bool {
	if backSpace(s) != backSpace(t) {
		return false
	}
	return true
}

func backSpace(str string) string {
	b := []byte(str)
	n := make([]byte, 0)
	l := len(str)
	s, f := 0, 0
	for ; f < l; f++ {
		if b[f] != '#' {
			n = append(n, b[f])
			s++
		} else {
			if s > 0 {
				s--
			}
			n = n[:s]
		}
	}
	return string(n)
}

func SortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	arr := make([]int, r+1)
	for i := r; l <= r; i-- {
		lpow, rpow := nums[l]*nums[l], nums[r]*nums[r]
		if lpow >= rpow {
			arr[i] = lpow
			l++
		} else {
			arr[i] = rpow
			r--
		}
	}
	return arr
}

func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hashmap := make(map[int32]int, len(s))
	for _, v := range s {
		hashmap[v]++
	}

	for _, v := range t {
		hashmap[v]--
		if hashmap[v] < 0 {
			return false
		}
	}

	return true
}

func GroupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string, len(strs))

	for _, str := range strs {

		strBytes := []byte(str)

		// 字符数组排序
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})

		// 排序后的字符数组转成字符串做key，原string做value
		hash[string(strBytes)] = append(hash[string(strBytes)], str)
	}

	var res [][]string
	for _, v := range hash {
		res = append(res, v)
	}

	return res
}

func FindAnagrams1(s string, p string) []int {
	pl := len(p)

	var res []int

	for i, j := 0, pl; j <= len(s); {
		if IsAnagram(s[i:j], p) {
			res = append(res, i)
		}
		i++
		j++
	}
	return res
}

func FindAnagrams(s string, p string) []int {
	pLength := len(p)
	sLength := len(s)
	var res []int
	if pLength > sLength {
		return res
	}
	sHashTable, phashTable := [26]int{}, [26]int{}

	for k, v := range p {
		sHashTable[s[k]-'a']++ // 初始化s头部的滑动窗口
		phashTable[v-'a']++
	}
	if sHashTable == phashTable {
		res = append(res, 0)
	}

	for k, _ := range s[:sLength-pLength] {
		sHashTable[s[k]-'a']--         //第二轮，需要把滑动窗口后的前一个旧字符删除，例如1，2，3，4，窗口长度为2，把1删除，把3加进来
		sHashTable[s[k+pLength]-'a']++ //第二轮，需要把滑动窗口后的新字符加进去
		if sHashTable == phashTable {
			res = append(res, k+1)
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	tmp := head
	for tmp.Next != nil {
		n := tmp.Next
		tmp.Next = pre
		tmp, pre = n, tmp
	}

	return pre
}
