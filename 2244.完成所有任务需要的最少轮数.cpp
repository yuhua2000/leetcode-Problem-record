/*
 * @lc app=leetcode.cn id=2244 lang=cpp
 *
 * [2244] 完成所有任务需要的最少轮数
 */
#include <vector>
#include <unordered_map>

using namespace std;

// @lc code=start
class Solution {
    public:
    int minimumRounds(vector<int>& tasks)
    {
        unordered_map<int, int> counts;
        for (int i : tasks)
        {
            counts[i]++;
        }
        int result = 0;
        for (auto [_, count] : counts)
        {
            if (count < 2)
            {
                return -1;
            }
            result += (count + 2) / 3;
        }

        return result;
    }
};
// @lc code=end

