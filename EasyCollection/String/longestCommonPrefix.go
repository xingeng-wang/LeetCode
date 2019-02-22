package String

// Write a function to find the longest common prefix string amongst an array of strings.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/887/
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
