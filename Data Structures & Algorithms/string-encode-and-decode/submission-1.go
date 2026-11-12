/*
Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

Machine 1 (sender) has the function:

String encode(List<String> strs) {
    // ... your code
    return encoded_string;
}
Machine 2 (receiver) has the function:

List<String> decode(String encoded_string) {
    // ... your code
    return decoded_strs;
}
So Machine 1 does:

String encoded_string = encode(strs);
and Machine 2 does:

List<String> decoded_strs = decode(encoded_string);
decoded_strs in Machine 2 should be the same as the input strs in Machine 1.

Implement the encode and decode methods.

Example 1:

Input: strs = ["Hello","World"]

Output: ["Hello","World"]
*/
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
