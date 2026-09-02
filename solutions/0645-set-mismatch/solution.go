func findErrorNums(nums []int) []int {
    n := len(nums)
    ans := make([]int, n+1)
    var duplikat, hilang int

    for _, val := range nums {
        ans[val]++
    }

    for i:=1 ; i<=n ; i++ {
        if ans[i] == 2 {
            duplikat = i
        }else if ans[i] == 0 {
            hilang = i
        }
    }

    return []int{duplikat, hilang}
}
