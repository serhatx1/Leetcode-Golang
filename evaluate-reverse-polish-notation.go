func evalRPN(tokens []string) int {
    if len(tokens)<2{
        number,_:=strconv.Atoi(tokens[0])
        return number
    }
    stack:=make([]string,0)
    result:=0
    for i:=0;i<len(tokens);i++{
        val:=string(tokens[i])
    if len(stack)<2{
        stack = append(stack, val) 
        continue

    }
    num1,_ := strconv.Atoi(stack[len(stack)-1])
    num2,_:= strconv.Atoi(stack[len(stack)-2])
    switch val {
    case "+":
        result = num2 + num1
        stack = stack[:len(stack)-2]
        stack = append(stack, strconv.Itoa(result)) 
    case "-":
        result = num2 - num1
        stack = stack[:len(stack)-2]
        stack = append(stack, strconv.Itoa(result)) 
    case "*":
        result = num2 * num1
        stack = stack[:len(stack)-2]
        stack = append(stack, strconv.Itoa(result)) 
    case "/":
        result = num2 / num1
        stack = stack[:len(stack)-2]
        stack = append(stack, strconv.Itoa(result)) 
    default:
        stack = append(stack, val) 
    }
    }
    return result
}
