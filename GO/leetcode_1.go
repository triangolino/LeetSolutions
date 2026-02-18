//I'll use comments a lot, if you don't like it, i'm sorry
//but i use this files to study

package main

import "fmt"

func main() {
	//space for func
    
    fmt.Print(twoSum([]int{2,7}[:],9))
}
func twoSum(nums []int, target int) []int {
    //basic controls for constraints
    if target > 1_000_000_000 || target < -1_000_000_000{return nil}
    if len(nums) < 2 || len(nums) > 100000{return nil}
    seek := make(map[int]int)
  
    for i := range nums{
        if nums[i] > 1_000_000_000 || nums[i] < -1_000_000_000{return nil}
        temp := target - nums [i]
        
        
        if idx, ok := seek[temp]; ok{
            return []int{idx, i}
        }
        
        seek[nums[i]] = i
    }
    return nil
    

}