package sorts

func QuickSort(nums []int) {
	mid := len(nums) / 2
	left, right := 0, len(nums)-1
	leftt, rightt := 0, len(nums)-1

	for left < right {
		if nums[left] > mid && nums[right] <= mid {
			nums[left], nums[right] = nums[right], nums[left]
		}

		if nums[left] <= mid {
			left++
		}

		if nums[right] > mid {
			right--
		}

	}
	QuickSort(nums[leftt:mid])
	QuickSort(nums[mid : rightt+1])
	return
}

func QuickSort1(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		QuickSort1(arr, low, pivot-1)
		QuickSort1(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i, j := low-1, low

	for ; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
