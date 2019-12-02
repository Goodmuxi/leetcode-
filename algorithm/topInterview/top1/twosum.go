package top1

type Result struct {
	First  int
	Second int
}

func TwoSumAll(num []int, target int) []Result {
	r := make([]Result, 0)
	mymap := make(map[int]int)
	for i, v := range num {
		complement := target - v
		if ii, ok := mymap[complement]; ok {
			r = append(r, Result{ii, i})
			continue
		}
		mymap[v] = i
	}
	return r
}

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
