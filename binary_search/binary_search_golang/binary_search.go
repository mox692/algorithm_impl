package binary_search_golang

// BinarySearch takes slice and target number.
// If target number found, return true and index of it.
func BinarySearch(a []int, targetNumber int) (bool, int) {
	cur_len := len(a)
	if cur_len == 0 {
		return false, 0
	}
	cur_pos := 0
	for {
		if cur_len%2 == 0 {
			center := (cur_len >> 1) - 1
			if a[center] > targetNumber {
				a = a[0:center]
				cur_len = (cur_len >> 1) - 1
				if cur_len == 0 {
					return false, 0
				}
				continue
			} else if a[center] < targetNumber {
				a = a[cur_len>>1:]
				cur_pos += cur_len >> 1
				cur_len = cur_len >> 1
				if cur_len == 0 {
					return false, 0
				}
				continue
			} else if a[center] == targetNumber {
				return true, cur_pos + center
			} else {
				return false, 0
			}
		} else {
			if cur_len == 0 {
				return false, 0
			}
			center := (cur_len - 1) >> 1
			if a[center] > targetNumber {
				a = a[0:center]
				cur_len = (cur_len >> 1) - 1
				if cur_len <= 0 {
					return false, 0
				}
				continue
			} else if a[center] < targetNumber {
				a = a[center+1:]
				cur_pos += center
				cur_len = (cur_len - 1) >> 1
				if cur_len <= 0 {
					return false, 0
				}
				continue
			} else if a[center] == targetNumber {
				return true, cur_pos + center
			} else {
				return false, 0
			}
		}
	}
}
