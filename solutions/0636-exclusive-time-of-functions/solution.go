func exclusiveTime(n int, logs []string) []int {
    ans := make([]int, n)
    stack := make([]int, 0)
    waktu := 0

    for _, val := range logs {
        val2 := strings.Split(val, ":")
        id, _ := strconv.Atoi(val2[0])
        time, _ := strconv.Atoi(val2[2])
        aksi := val2[1]

        switch aksi{
            case "start" :
            n := len(stack)
            if n == 0 {
                stack = append(stack, id)
                waktu = time
            }else{
                ans[stack[n-1]] += time - waktu
                stack = append(stack, id)
                waktu = time
            }

            case "end" :
            n:= len(stack)
            ans[stack[n-1]] = time - waktu + 1 + ans[stack[n-1]]
            waktu = time +1
            stack = stack[:len(stack)-1]
        }
    }
    return ans
}
