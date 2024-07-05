#include <vector>

using namespace std;

class Solution {
    public:
    vector<vector<int>> modifiedMatrix(vector<vector<int>>& matrix)
    {
        int n = matrix.size();
        int m = matrix[0].size();
        vector<int> idx;
        vector<int> maxSum(m, -2);
        for (size_t i = 0; i < n; i++)
        {
            for (size_t j = 0; j < m; j++)
            {
                maxSum[j] = max(maxSum[j], matrix[i][j]);
                if (matrix[i][j] == -1)
                {
                    idx.emplace_back(i * m + j);
                }
            }
        }

        for (size_t i = 0; i < idx.size(); i++)
        {
            int x = idx[i] / m;
            int y = idx[i] % m;
            matrix[x][y] = maxSum[y];
        }
        
        return matrix;
    }
};