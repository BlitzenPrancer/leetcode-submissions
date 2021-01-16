func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for k := range nums {
        a := target - nums[k]
        if index, doesExist := m[a]; doesExist {
            return []int{index, k}
        } else {
            m[nums[k]] = k
        }
    }
    return nil
}
