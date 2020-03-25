package problem0010

// https://leetcode.com/problems/regular-expression-matching/
func isMatch(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	dp[0][0] = true
	// 주제 설정에 따라 ""는 "a * b * c *"와 일치 할 수 있습니다.
	// 따라서 해당 dp를 true로 설정해야합니다.

	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				/* p [j] 및 s [i]는 일치 할 수 있으므로 이전 일치하는 한 여기에서 일치시킬 수 있습니다 */
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				/* 현재 p [j]의 일치는 p [j-1]의 내용과 관련이 있습니다. */
				if p[j-1] != s[i] && p[j-1] != '.' {
					/**
					* p [j]는 s [i]와 일치 할 수 없습니다
					* p [j-1 : j + 1]은 ""로만 취급 될 수 있습니다
					 */
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/**
					* p [j]는 s [i]와 일치
					* p [j-1; j + 1]은 "x *"로 세 가지 해석이 있습니다.
					 */
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x *"는 ""*/
						dp[i+1][j] || /* "x *"는 "x"로 해석됩니다 */
						dp[i][j+1] /* "x *"는 "xx ..."로 해석됩니다. */
				}
			}
		}
	}

	return dp[sSize][pSize]
}
