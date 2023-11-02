package ang

func Divide(dividend int, divisor int) int {

	sameSign := true
	// 绝对值
	if dividend < 0 {
		dividend = 0 - dividend
		sameSign = false
	}

	// 绝对值
	if divisor < 0 {
		divisor = 0 - divisor
		sameSign = false
	}

	if dividend < divisor {
		return 0
	}

	mul, i := divisor, 0
	for mul < dividend {
		mul = mul + divisor //相当于*2
		i++
	}
	res := i

	// 正负数
	if !sameSign {
		res = 0 - res
	}
	return res
}
