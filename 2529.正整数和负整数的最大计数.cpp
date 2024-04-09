/*
 * @lc app=leetcode.cn id=2529 lang=cpp
 *
 * [2529] 正整数和负整数的最大计数
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int lowerBound(vector<int>& nums, int val)
    {
        int l = 0, r = nums.size() ;
        while (l < r)
        {
            int m = (l + r) >> 1;
            if (nums[m] >= val)
                r = m;
            else
                l = m + 1;

        }
        return l;
    }

    int maximumCount(vector<int>& nums)
    {
        int pos1 = lowerBound(nums, 0);
        int pos2 = lowerBound(nums, 1);
        return max(pos1, (int)nums.size() - pos2);
    }
};
// @lc code=end

