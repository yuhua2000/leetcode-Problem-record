/*
 * @lc app=leetcode.cn id=994 lang=cpp
 *
 * [994] 腐烂的橘子
 */
#include <vector>
#include <unordered_set>
#include <iostream>

using namespace std;

// @lc code=start
class Solution {
    private:
    int dx[4] = { 0, 0, -1, 1 };
    int dy[4] = { -1, 1, 0, 0 };
    public:
    int orangesRotting(vector<vector<int>>& grid)
    {
        vector<int> maps(100);
        vector<int> queue;
        unordered_set<int> visit(100);
        int goodOrange = 0;
        int n = grid.size(), m = grid[0].size();
        for (size_t i = 0; i < n; i++)
        {
            for (size_t j = 0; j < m; j++)
            {
                if (grid[i][j] == 1)
                {
                    goodOrange++;
                    maps[i * 10 + j] = 1;
                }
                else if (grid[i][j] == 2)
                {
                    queue.push_back(i * 10 + j);
                }
            }
        }

        int result = 0;
        while (!queue.empty())
        {
            if (visit.size() == goodOrange)
                return result;
            auto temp = queue;
            queue.clear();
            result++;
            for (auto point : temp)
            {
                int x = point / 10;
                int y = point % 10;
                for (int i = 0;i < 4;i++)
                {
                    int new_x = x + dx[i];
                    int new_y = y + dy[y];
                    cout << new_x << new_y << "<--" << x << y;
                    int new_point = new_x * 10 + new_y;
                    if (new_x >= 0 && new_x < n && new_y >= 0 && new_y < m &&
                        maps[new_point] && visit.find(new_point) == visit.end())
                    {
                        visit.emplace(new_point);
                        queue.push_back(new_point);
                    }
                }
            }
        }

        if (visit.size() == goodOrange)
            return result;

        return -1;
    }
};
// @lc code=end

