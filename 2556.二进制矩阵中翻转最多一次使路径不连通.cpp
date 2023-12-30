/*
 * @lc app=leetcode.cn id=2556 lang=cpp
 *
 * [2556] 二进制矩阵中翻转最多一次使路径不连通
 */

#include <vector>
using namespace std;

// @lc code=start
class Solution {
    public:
    bool isPossibleToCutPath(vector<vector<int>>& grid)
    {
        int m = grid.size(), n = grid[0].size();
        function<bool(int, int)> dfs = [&] (int x, int y)->bool {
            if (x==m-1&&y==n-1)
            {
                return true;
            }
            grid[x][y] = 0;
            return x < m - 1 && grid[x + 1][y] && dfs(x + 1, y) ||
                y < n - 1 && grid[x][y + 1] && dfs(x, y + 1);
        };
        return !dfs(0, 0) || !dfs(0, 0);
    }
};
// @lc code=end

