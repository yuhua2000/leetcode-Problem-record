#include <vector>
#include <queue>

using namespace std;

/*
 * @lc app=leetcode.cn id=1696 lang=cpp
 *
 * [1696] 跳跃游戏 VI
 */

// @lc code=start
class Solution {
public:
    int maxResult(vector<int>& nums, int k) {
         int n = nums.size();
        vector<int> dp(n);
        dp[0] = nums[0];
        deque<int> queue;
        queue.push_back(0);
        for (int i = 1; i < n; i++)
        {
            while (!queue.empty()&&queue.front()<i-k)
            {
                queue.pop_front();
            }
            dp[i] = dp[queue.front()] + nums[i];
            while (!queue.empty()&&dp[queue.back()]<=dp[i])
            {
                queue.pop_back();
            }
            queue.push_back(i);
        }
        return dp[n - 1];
    }
};
// @lc code=end

