#include <vector>
#include <unordered_set>

using namespace std;

class Solution {
    public:
    int firstCompleteIndex(vector<int>& arr, vector<vector<int>>& mat)
    {
        int n = mat.size(), m = mat[0].size();
        vector<int> row(n), column(m);
        vector<tuple <int, int>> numsIndex((n * m) + 1);
        for (size_t i = 0; i < n; i++)
        {
            for (size_t j = 0; j < m; j++)
            {
                numsIndex[mat[i][j]] = tuple<int, int>{ i,j };
            }
        }
        for (int i = 0;i < arr.size();i++)
        {
            auto& t = numsIndex[arr[i]];

            row[get<0>(t)]++;
            column[get<1>(t)]++;
            if (row[get<0>(t)] == m || column[get<1>(t)] == n)
            {
                return i;
            }
        }
        return -1;
    }
};