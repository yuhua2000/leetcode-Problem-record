package leetcode

/*
 * @lc app=leetcode.cn id=1797 lang=golang
 *
 * [1797] 设计一个验证系统
 */

// @lc code=start
type AuthenticationManager struct {
	ttl      int
	tokenDic map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl:      timeToLive,
		tokenDic: make(map[string]int),
	}
}

func (a *AuthenticationManager) Generate(tokenId string, currentTime int) {
	a.tokenDic[tokenId] = currentTime + a.ttl
}

func (a *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if t, ok := a.tokenDic[tokenId]; ok {
		if t <= currentTime {
			delete(a.tokenDic, tokenId)
			return
		}
		a.tokenDic[tokenId] = currentTime + a.ttl
	}
}

func (a *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	for k, v := range a.tokenDic {
		if v <= currentTime {
			delete(a.tokenDic, k)
		}
	}
	return len(a.tokenDic)
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
// @lc code=end
