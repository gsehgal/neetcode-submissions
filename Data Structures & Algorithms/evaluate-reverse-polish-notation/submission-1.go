func evalRPN(tokens []string) int {
	var stack []int
	for _,token := range tokens {
		if token == "+" || token == "-" || token =="*" || token == "/" {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if token == "+" {
				stack = append(stack,num1+num2)
			}else if token == "-" {
				stack = append(stack,num2-num1)
			}else if token == "*" {
				stack = append(stack,num1*num2)
			}else {
				stack = append(stack,num2/num1)
			}
		}else {
			num,_ := strconv.Atoi(token)
			stack = append(stack,num)
		}
	}
	return stack[len(stack)-1]
}
