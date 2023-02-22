/*
 * @lc app=leetcode.cn id=1604 lang=golang
 *
 * [1604] 警告一小时内使用相同员工卡大于等于三次的人
 */

// @lc code=start
func alertNames(keyName []string, keyTime []string) (ret []string) {

	parseTime := func(time string) int {
		h := int(time[0]-'0')*10 + int(time[1]-'0')
		m := int(time[3]-'0')*10 + int(time[4]-'0')
		return h*60 + m
	}
	credit := make(map[string][]int)
	for i, name := range keyName {
		time := parseTime(keyTime[i])
		credit[name] = append(credit[name], time)
	}
	for name, dit := range credit {
		sort.Ints(dit)
		for i, t := range dit[2:] {
			if t-dit[i] <= 60 {
				ret = append(ret, name)
				break
			}
		}

	}
	sort.Strings(ret)
	return ret
}

// @lc code=end
