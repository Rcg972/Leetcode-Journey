func evalRPN(tokens []string) int {
    ans := make([]int, 0)

    for _, val := range tokens {
        switch val {
            case "/":
            n := len(ans)
            result := ans[n - 2] / ans[n - 1]
            ans = ans[:len(ans)-2]
            ans = append(ans, result)

            case "*":
            n := len(ans)
            result := ans[n - 2] * ans[n - 1]
            ans = ans[:len(ans)-2]
            ans = append(ans, result)

            case "+":
            n := len(ans)
            result := ans[n - 2] + ans[n - 1]
            ans = ans[:len(ans)-2]
            ans = append(ans, result)
            
            case "-":
            n := len(ans)
            result := ans[n - 2] - ans[n - 1]
            ans = ans[:len(ans)-2]
            ans = append(ans, result)

            default:
            if s, err := strconv.Atoi(val); err == nil{
                ans = append(ans, s)
            }
            
        }
    }

    return ans[0]
}
