/*
 * @lc app=leetcode.cn id=198 lang=cpp
 *
 * [198] 打家劫舍
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int rob(vector<int>& nums)
    {
        int a = 0, b = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int ta = a;
            a = max(b + nums[i], a);
            b = ta;
        }
        return a;
    }
};
// @lc code=end

