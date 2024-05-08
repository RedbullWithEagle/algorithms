package array

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "ReverseArray1", args: args{source: []int{5, 7, 8, 2, 4, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseArray(tt.args.source)
			jsData, _ := json.Marshal(got)
			fmt.Println(string(jsData))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "BSearch", args: args{
			nums:   []int{1, 5, 8, 12, 34, 45, 46, 78, 88},
			target: 34,
		}},
		{name: "BSearch2", args: args{
			nums:   []int{2, 2},
			target: 2,
		}},
		{name: "BSearch3", args: args{
			nums:   []int{},
			target: 2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BSearch(tt.args.nums, tt.args.target)
			jsData, _ := json.Marshal(got)
			fmt.Println(string(jsData))
			if got != tt.want {
				t.Errorf("BSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
