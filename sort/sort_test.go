package sort

import "testing"

/***************************************
*单元测试要点
1.文件 xx_test.go
2.方法名TestXXX   参数：t *testing.T
3.go test ./文件的包路径
****************************************/
func TestSort(t *testing.T) {
	s := []int{8, 2, 11, 7, 24, 1, 6}
	expected := []int{1, 2, 6, 7, 8, 11, 24}
	BubbleSort(s)
	if !sliceEq(s, expected) {
		t.Errorf("s=%v; expected %v", s, expected)
	}
}

func TestSortArray(t *testing.T) {
	var sortTests = []struct {
		source   []int // input
		expected []int // expected result
	}{
		{[]int{5, 4, 6, 7, 3, 8, 2, 9, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{3, 6, 7, 3}, []int{3, 3, 6, 7}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range sortTests {
		BubbleSort(tt.source)
		if !sliceEq(tt.source, tt.expected) {
			t.Errorf("source is %v; expected %v", tt.source, tt.expected)
			continue
		}
		t.Logf("soucce is %v,is ok", tt.source)
	}
}
