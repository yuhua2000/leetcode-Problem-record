#include <vector>

using namespace std;
/*
 * @lc app=leetcode.cn id=419 lang=cpp
 *
 * [419] 甲板上的战舰
 */

 // @lc code=start
class Solution {
    public:
    int countBattleships(vector<vector<char>>& board)
    {
        int result = 0;
        int n = board.size(), m = board[0].size();
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if (board[i][j] == 'X')
                {
                    if (i - 1 >= 0 && board[i - 1][j] == 'X')
                        continue;

                    if (j - 1 >= 0 && board[i][j - 1] == 'X')
                        continue;

                    result++;
                }
            }
        }

        return result;
    }
};
// @lc code=end

