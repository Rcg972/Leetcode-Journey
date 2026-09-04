func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, 0)
    ans := make([]int, len(temperatures))

    for index, val := range temperatures {
        for len(stack) > 0 && val > temperatures[stack[len(stack)-1]]{
            ans[stack[len(stack)-1]] = index - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }

        stack = append(stack, index)

    }
    return ans
}
