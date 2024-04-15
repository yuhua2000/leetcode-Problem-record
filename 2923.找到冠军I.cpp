#include <vector>

using namespace std;

class Solution {
    public:
    int findChampion(vector<vector<int>>& grid)
    {
        int n = grid.size();
        for (int i = 0;i < n;i++)
        {
            int sum = 0;
            for (auto j : grid[i])
            {
                sum += j;
            }
            if (sum == n - 1)
            {
                return i;
            }
        }
        return -1;
    }
};

