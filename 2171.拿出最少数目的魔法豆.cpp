#include <vector>
#include <numeric>
#include <algorithm>

using namespace std;

/*
 * @lc app=leetcode.cn id=2171 lang=cpp
 *
 * [2171] 拿出最少数目的魔法豆
 */

// @lc code=start
class Solution {
public:
    long long minimumRemoval(vector<int>& beans) {
        int n = beans.size();
        sort(beans.begin(), beans.end());
        long long total = accumulate(beans.begin(), beans.end(), 0LL); // 豆子总数
        long long res = total;
        for (int i = 0; i < n; i++)
        {
            res = min(res, total - (long long)beans[i] * (n - i));
        }
        
        return res;
    }
};
// @lc code=end

