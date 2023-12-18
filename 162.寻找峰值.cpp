/*
 * @lc app=leetcode.cn id=162 lang=cpp
 *
 * [162] 寻找峰值
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int findPeakElement1(vector<int>& nums)
    {
        bool prior = true;
        for (size_t i = 0; i < nums.size() - 1; i++)
        {
            /* code */
            if (nums[i] > nums[i + 1])
            {
                /* code */
                if (prior)
                {
                    return i;
                }
                prior = false;
            }
            else
            {
                prior = true;
            }

        }
        return nums.size() - 1;
    }
    int findPeakElement(vector<int>& nums)
    {
        int n = nums.size();
        auto get = [&] (int i)->pair<int, int> {
            if (i == -1 || i == n)
            {
                return { 0,0 };
            }
            return { 1,nums[i] };
        };

        int left = 0, right = n - 1, ans = -1;
        while (left <= right)
        {
            int mid = (left + right) / 2;
            if (get(mid - 1) < get(mid) && get(mid) > get(mid + 1))
            {
                ans = mid;
                break;
            }

            if (get(mid) > get(mid + 1))
            {
                right = mid - 1;
            }
            else
            {
                left = mid + 1;
            }
        }
        return ans;
    }
};
// @lc code=end

