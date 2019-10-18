package _01910

import "testing"


func TestContainsDuplicate1(t *testing.T) {
	type testData struct{
		data []int
		expect bool
	}
	inputlist := []testData{
		testData{[]int{1, 2, 3},false},
		testData{[]int{1, 2, 2, 3},true},
		testData{[]int{0, 1, 2, 3, 4, 5, 0},true},
	}
	for _, input := range inputlist {
		if ContainsDuplicate1(input.data)&&input.expect{
			t.Logf("input:%v, expect:%v, output:%v",input.data, input.expect, ContainsDuplicate1(input.data))
		}
	}

}

func TestContainsDuplicate2(t *testing.T) {
	type testData struct{
		data []int
		expect bool
	}
	inputlist := []testData{
		testData{[]int{1, 2, 3},false},
		testData{[]int{1, 2, 2, 3},true},
		testData{[]int{0, 1, 2, 3, 4, 5, 0},true},
	}
	for _, input := range inputlist {
		if (ContainsDuplicate2(input.data)&&input.expect){
			t.Logf("input:%v, expect:%v, output:%v",input.data, input.expect, ContainsDuplicate1(input.data))
		}
	}
}
