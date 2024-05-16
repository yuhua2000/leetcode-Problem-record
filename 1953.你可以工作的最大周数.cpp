/*
 * @lc app=leetcode.cn id=1953 lang=cpp
 *
 * [1953] 你可以工作的最大周数
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    long long numberOfWeeks(vector<int>& milestones)
    {
        long long maxMilestones = 0;
        long long countMilestones = 0;
        for (auto milestones : milestones)
        {
            countMilestones += milestones;
            maxMilestones = max(maxMilestones, (long long)milestones);
        }

        long long rest = countMilestones - maxMilestones;
        if (maxMilestones > rest + 1)
        {
            return rest * 2 + 1;
        }
        else
        {
            return countMilestones;
        }

    }
};
// @lc code=end

