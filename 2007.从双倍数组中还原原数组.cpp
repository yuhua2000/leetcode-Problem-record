/*
 * @lc app=leetcode.cn id=2007 lang=cpp
 *
 * [2007] 从双倍数组中还原原数组
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    vector<int> findOriginalArray(vector<int>& changed)
    {
        if (changed.size() % 2 != 0)
            return {};


        vector<int> nums(100001);
        for (size_t i = 0; i < changed.size(); i++)
            nums[changed[i]]++;

        if (nums[0] % 2 != 0)
            return {};

        vector<int> result(nums[0] / 2, 0);
        for (size_t i = 1; i < nums.size(); i++)
        {
            while (nums[i])
            {
                if ((i * 2) < 100001 && nums[i * 2])
                {
                    result.emplace_back(i);
                    nums[i]--;
                    nums[i * 2]--;
                }
                else
                {
                    return{};
                }
            }
        }

        return result;

    }
};
// @lc code=end

