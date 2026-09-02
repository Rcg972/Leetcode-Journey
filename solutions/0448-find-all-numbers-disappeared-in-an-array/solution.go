func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    ans := make([]int, n+1)
    missing := make([]int, 0, n)

    for _, val := range nums {
        ans[val]++
    }

    for i:=1 ; i<=n ; i++ {
        if ans[i] == 0{
            missing = append(missing, i)
        }
    }

    return missing
}
