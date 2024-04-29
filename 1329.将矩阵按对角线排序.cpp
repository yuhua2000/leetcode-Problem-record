/*
 * @lc app=leetcode.cn id=1329 lang=cpp
 *
 * [1329] 将矩阵按对角线排序
 */
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

// @lc code=start
class Solution {
    public:
    vector<vector<int>> diagonalSort(vector<vector<int>>& mat)
    {
        int n = mat.size(), m = mat[0].size();

        vector<vector<int>> spreadMat(m + n);

        for (size_t i = 0; i < n; i++)
        {
            for (size_t j = 0; j < m; j++)
            {
                spreadMat[i - j + m].push_back(mat[i][j]);
            }
        }

        for (size_t i = 0; i < spreadMat.size(); i++)
        {
            sort(spreadMat[i].rbegin(), spreadMat[i].rend());
        }

        
        for (size_t i = 0; i < n; i++)
        {
            for (size_t j = 0; j < m; j++)
            {
                mat[i][j] = spreadMat[i - j + m].back();
                spreadMat[i - j + m].pop_back();    
            }
        }

        return mat;
    }
};
// @lc code=end

