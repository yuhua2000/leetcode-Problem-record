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
        for (size_t i = 1; i <= nums.size() / 2; i++)
        {
            if (nums[i] == 0)
                continue;

            if (nums[i] <= nums[i * 2])
            {
                nums[i * 2] -= nums[i];
                result.insert(result.end(), nums[i], i);
            }
            else
            {
                return{};
            }
        }


        if (result.size() == changed.size() / 2)
            return result;

        return {};
    }
};
// @lc code=end

