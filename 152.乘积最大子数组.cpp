/*
 * @lc app=leetcode.cn id=152 lang=cpp
 *
 * [152] 乘积最大子数组
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int maxProduct(vector<int>& nums)
    {
        double maxF = nums[0], minF = nums[0];
         double  result = nums[0];

        for (int i = 1; i < nums.size(); i++)
        {
             double mx = maxF, mn = minF;
            maxF = max(static_cast< double>(nums[i]), max(mx * nums[i], mn * nums[i]));
            minF = min(static_cast< double>(nums[i]), min(mx * nums[i], mn * nums[i]));
            result = max(result, maxF);
        }

        return static_cast<int>(result);
    }
};
// @lc code=end

