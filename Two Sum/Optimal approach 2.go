func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for k, v := range nums {
        diff := target - v
        index, doesExist := m[diff]
        if doesExist == true {
            return []int{index, k}
        }
        m[v] = k
    }
    return nil
}
