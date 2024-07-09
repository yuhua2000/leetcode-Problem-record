/*
 * @lc app=leetcode.cn id=1958 lang=cpp
 *
 * [1958] 检查操作是否合法
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    bool checkMove(vector<vector<char>>& board, int rMove, int cMove, char color)
    {
        int n = board.size();
        int m = board[0].size();
        vector<pair<int, int>> diretions{ {-1,0},{1,0},{0,1},{0,-1},{1,1},{-1,-1},{1,-1},{-1,1} };
        for (auto dire : diretions)
        {
            int tempRMove = rMove + dire.first;
            int tempCMove = cMove + dire.second;
            int tempColor = color;
            int len = 1;
            while (tempRMove >= 0 && tempRMove < n && tempCMove >= 0 && tempCMove < m)
            {
                len++;
                if (board[tempRMove][tempCMove] == '.')
                {
                    break;
                }
                else if (board[tempRMove][tempCMove] != tempColor)
                {
                    tempRMove += dire.first;
                    tempCMove += dire.second;
                }
                else if (len >= 3)
                {
                    return true;
                }
                else
                {
                    break;
                }
            }
        }
        return false;
    }
};
// @lc code=end

