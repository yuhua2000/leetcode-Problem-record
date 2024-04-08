/*
 * @lc app=leetcode.cn id=2009 lang=cpp
 *
 * [2009] 使数组连续的最少操作数
 */
#include <vector>
#include <algorithm>
#include <iostream>

using namespace std;

// @lc code=start
class Solution {
    public:
    int minOperations(vector<int>& nums)
    {
        int n = nums.size();
        sort(nums.begin(), nums.end());
        int l = 0, r = 0, numCount = 1;
        int result = n;
        for (;r < n;r++)
        {
            while (nums[r] - nums[l] > n - 1)
            {
                l++;
                if (l > 0 && nums[l] == nums[l - 1])
                {
                    continue;
                }
                numCount--;
            }
            cout << l << "---" << r << "----" << nums[l] << "---" << nums[r] << "---" << nums[r - 1] << "---" << numCount << "---" << n << endl;
            result = min(result, n - numCount);
            if (r < n - 1 && nums[r] == nums[r + 1])
            {
                continue;
            }
            cout << "---" << endl;
            numCount++;

        }
        return result;
    }
};
// @lc code=end

//[8,5,9,9,8,4] [10,49,16,22,28,34,35,7] [4,2,5,3]