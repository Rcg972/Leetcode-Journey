func buildArray(target []int, n int) []string {
    ans := make([]string, 0, n)
    count := 0

    for i:=1 ; i<=n ; i++ {
        if count == len(target){
            return ans
        }

        if i == target[count]{
            ans = append(ans, "Push")
            count += 1
        }else {
            ans = append(ans, "Push" , "Pop")
        }
    }

    return ans
}
