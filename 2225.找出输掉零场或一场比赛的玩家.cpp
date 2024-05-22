/*
 * @lc app=leetcode.cn id=2225 lang=cpp
 *
 * [2225] 找出输掉零场或一场比赛的玩家
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    vector<vector<int>> findWinners(vector<vector<int>>& matches)
    {
        vector<int> ans(100001);
        for (auto match : matches)
        {
            int winner = match[0], loser = match[1];
            if (ans[winner] == 0)
                ans[winner] = 1;

            if (ans[loser] <= 1)
                ans[loser] = 2;
            else if (ans[loser] == 2)
                ans[loser] = 3;
        }

        vector<int> champion, runner;
        for (size_t i = 0; i < ans.size(); i++)
        {
            if (ans[i] == 1)
                champion.emplace_back(i);
            else if (ans[i] == 2)
                runner.emplace_back(i);
        }
        return vector<vector<int>>{ champion, runner };
    }
};
// @lc code=end

