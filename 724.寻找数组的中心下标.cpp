/*
 * @lc app=leetcode.cn id=724 lang=cpp
 *
 * [724] 寻找数组的中心下标
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int pivotIndex(vector<int>& nums)
    {
        int sum = 0;
        for (size_t i = 0; i < nums.size(); i++)
        {
            sum += nums[i];
        }
        int temp = 0;
        for (size_t i = 0; i < nums.size(); i++)
        {
            sum -= nums[i];
            if (sum == temp)
            {
                return i;
            }
            temp += nums[i];
        }

        return -1;
    }
};
// @lc code=end

