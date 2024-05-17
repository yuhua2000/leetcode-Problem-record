/*
 * @lc app=leetcode.cn id=826 lang=cpp
 *
 * [826] 安排工作以达到最大收益
 */
#include <vector>
#include <algorithm>

using namespace std;

// @lc code=start
class Solution {
    public:
    int maxProfitAssignment(vector<int>& difficulty, vector<int>& profit, vector<int>& worker)
    {
        int n = difficulty.size();
        vector<int> jobIds(n);
        for (size_t i = 0; i < n; i++)
            jobIds[i] = i;

        sort(jobIds.begin(), jobIds.end(), [&] (const int& a, const int& b)->bool {
            return difficulty[a] < difficulty[b];
            }
        );

        sort(worker.begin(), worker.end(), [&] (const int& a, const int& b)->bool {
            return a < b;
            }
        );

        int maxProfit = 0;
        int result = 0;
        int numWorkers = worker.size();
        int i = 0;
        for (auto workerSkill : worker)
        {
            for (;i < n && workerSkill >= difficulty[jobIds[i]];i++)
                maxProfit = max(maxProfit, profit[jobIds[i]]);

            result += maxProfit;
        }

        return result;
    }
};
// @lc code=end



