/*
 * @lc app=leetcode.cn id=213 lang=cpp
 *
 * [213] 打家劫舍 II
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int rob(vector<int>& nums)
    {
        if (nums.size() == 1)
            return nums[0];

        int first = 0, second = 0;
        for (int i = 0; i < nums.size() - 1; i++)
        {
            int temp = first;
            first = max(first, second + nums[i]);
            second = temp;
        }
        int result = first;

        first = 0, second = 0;
        for (int i = 1; i < nums.size(); i++)
        {
            int temp = first;
            first = max(first, second + nums[i]);
            second = temp;
        }

        result = max(result, first);
        return result;

    }
};
// @lc code=end

