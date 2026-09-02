func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    ans := make([]int, 0 ,n)
    count := 0

    for _, val := range nums {
        for _, val2 := range nums {
            if val2 < val {
                count++
            }
        }
        ans = append(ans, count)
        count = 0
    }

    return ans
}
