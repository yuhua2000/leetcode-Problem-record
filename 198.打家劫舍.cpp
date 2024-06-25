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
        int first = 0, second = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int temp = second;
            second = max(first + nums[i], second);
            first = temp;
        }

        return second;
    }
};
// @lc code=end

