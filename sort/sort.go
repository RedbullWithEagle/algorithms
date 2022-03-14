package sort

//BubbleSort 冒泡排序
func BubbleSort(s []int) error {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		for j := 0; j < lenS-i-1; j++ {
			if s[j] > s[j+1] {
				tmp := s[j+1]
				s[j+1] = s[j]
				s[j] = tmp
			}
		}
	}
	return nil
}

func sliceEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
