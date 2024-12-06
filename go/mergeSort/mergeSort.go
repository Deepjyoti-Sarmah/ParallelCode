package mergesort

type MergeSort struct {
	nums []int
}

func NewMergeSort(nums []int) *MergeSort {
	return &MergeSort{nums: nums}
}

func (m *MergeSort) Sort() {
	if len(m.nums) == 0 {
		return
	}

	m.recursiveSort(0, len(m.nums)-1)
}

func (m *MergeSort) recursiveSort(left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	m.recursiveSort(left, mid)
	m.recursiveSort(mid+1, right)

	result := make([]int, 0, right-left+1)
	i, j := left, mid+1

	for i <= mid && j <= right {
		if m.nums[i] <= m.nums[j] {
			result = append(result, m.nums[i])
			i++
		} else {
			result = append(result, m.nums[j])
			j++
		}
	}

	for i <= mid {
		result = append(result, m.nums[i])
		i++
	}

	for j <= right {
		result = append(result, m.nums[j])
		j++
	}

	copy(m.nums[left:right+1], result)
}
