/*
 * @lc app=leetcode.cn id=1673 lang=cpp
 *
 * [1673] 找出最具竞争力的子序列
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    vector<int> mostCompetitive(vector<int>& nums, int k)
    {
        vector<int> result;
        int n = nums.size();

        for (int i = 0; i < n; i++)
        {
            while (!result.empty() && nums[i] < result.back() && i < n - k + result.size())
            {
                result.pop_back();
            }
            result.push_back(nums[i]);

        }
        result.resize(k);
        return result;
    }
};
// @lc code=end
