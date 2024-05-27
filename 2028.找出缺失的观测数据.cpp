/*
 * @lc app=leetcode.cn id=2028 lang=cpp
 *
 * [2028] 找出缺失的观测数据
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    vector<int> missingRolls(vector<int>& rolls, int mean, int n)
    {

        int  diff = (n + rolls.size()) * mean;
        for (auto roll : rolls)
            diff -= roll;

        if (diff < n || diff >(n * 6))
            return {};


        vector<int> result(n);
        for (int i = 0;i < n;i++)
        {
            int d = min(6, diff - n + i + 1);
            result[i] = d;
            diff -= d;
        }

        return result;
    }
};
// @lc code=end

