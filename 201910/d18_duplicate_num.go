package _01910

func ContainsDuplicate1(nums []int) bool {
	for i:=0;i<len(nums);i++{
		for j,num:=range nums{
			if j>0&&num==nums[i]{
				return true
			}
		}
	}
	return false
}

func ContainsDuplicate2(nums []int)bool{
	numMap:=make(map[int]int,0)
	for _,v:=range nums{
		if _,ok:=numMap[v];ok{
			return true
		}
	}
	return false
}
