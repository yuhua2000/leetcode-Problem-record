#include <vector>
#include <unordered_map>

using namespace std;
/*
 * @lc app=leetcode.cn id=447 lang=cpp
 *
 * [447] 回旋镖的数量
 */

 // @lc code=start
class Solution {
    public:
    int numberOfBoomerangs(vector<vector<int>>& points)
    {
        int ans = 0;
        for (auto& p : points)
        {
            unordered_map<int, int> cnt;
            for (auto&q:points)
            {
                int dis = (p[0] - q[0]) * (p[0] - q[0]) + (p[1] - q[1]) * (p[1] - q[1]);
                ++cnt[dis];
            }
            for (auto &[_,m]:cnt)
            {
                ans += m * (m - 1);
            }
        }
        return ans;
    }
};
// @lc code=end

