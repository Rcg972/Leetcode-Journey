func shuffle(nums []int, n int) []int {
    ans := make([]int,0, 2*n)

    for i:=0; i<n; i++{
        ans = append(ans, nums[i])
        ans = append(ans, nums[i+n])
    }

    return ans
}
