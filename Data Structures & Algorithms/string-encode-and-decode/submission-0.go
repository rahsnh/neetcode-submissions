type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder

	for _, s := range strs {
		res.WriteString(strconv.Itoa(len(s)))
		res.WriteByte('#')
		res.WriteString(s)
	}

	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		wordLength, _ := strconv.Atoi(encoded[i:j])
		res = append(res, encoded[j+1:j+1+wordLength])
		i += len(encoded[i:j])+ wordLength + 1
	}

	return res
}
