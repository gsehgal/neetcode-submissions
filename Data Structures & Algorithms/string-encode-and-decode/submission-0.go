type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, val := range strs {
		builder.WriteString(strconv.Itoa(len(val)))
		builder.WriteString("#")
		builder.WriteString(val)
	}
	return builder.String()
}


func (s *Solution) Decode(encoded string) []string {
	var res []string
	for i:=0;i<len(encoded); {
		j := strings.Index(encoded[i:],"#")
		l, _ := strconv.Atoi(encoded[i:i+j])
		res = append(res,encoded[i+j+1:i+j+1+l])
		i += j+l+1
	}
	return res

}
