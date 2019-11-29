package top1

func TwoSum(num []int, target int) []int {
	mymap := make(map[int]int)
	for i, v := range num {
		complement := target - v
		if ii, ok := mymap[complement]; ok {
			return []int{ii, i}
		}
		mymap[v] = i
	}
	return nil
}
