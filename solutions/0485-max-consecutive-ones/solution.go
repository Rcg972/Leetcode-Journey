func findMaxConsecutiveOnes(nums []int) int {
    var terbesar, sekarang int

    for _, val := range nums{
        if val == 1{
            sekarang += 1
        }else{
            if sekarang > terbesar {
                terbesar = sekarang
                sekarang = 0
            }else{
                sekarang = 0
            }
        }
    }
    return max(terbesar, sekarang)
}
