func dailyTemperatures(temperatures []int) []int {
    r := make([]int, len(temperatures))
    stack := make([]int, 0)
    for i := 0; i < len(temperatures); i++ {
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            index := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            r[index] = i - index
        }
        stack = append(stack, i)
    }
    return r
}
