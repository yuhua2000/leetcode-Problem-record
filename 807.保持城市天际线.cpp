/*
 * @lc app=leetcode.cn id=807 lang=cpp
 *
 * [807] 保持城市天际线
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int maxIncreaseKeepingSkyline(vector<vector<int>>& grid)
    {
        int n = grid.size();
        vector<int> rowMax(n, 0);
        vector<int> columnMax(n, 0);
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
            {
                rowMax[i] = max(rowMax[i], grid[i][j]);
                columnMax[j] = max(columnMax[j], grid[i][j]);
            }
        }

        int result = 0;
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
            {
                result += min(rowMax[i], columnMax[j]) - grid[i][j];
            }
        }

        return result;
    }
};
// @lc code=end

