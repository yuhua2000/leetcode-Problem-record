/*
 * @lc app=leetcode.cn id=2105 lang=cpp
 *
 * [2105] 给植物浇水 II
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int minimumRefill(vector<int>& plants, int capacityA, int capacityB)
    {
        int n = plants.size();
        int result = 0;
        int currentCapacityA = capacityA, currentCapacityB = capacityB;
        for (int i = 0; i < (n + 1) / 2; i++)
        {
            if (i == n - i - 1)
            {
                if (currentCapacityA < plants[i] && currentCapacityB < plants[n - i - 1])
                {
                    result++;
                    break;
                }
            }
            else
            {
                if (currentCapacityA < plants[i])
                {
                    result++;
                    currentCapacityA = capacityA - plants[i];
                }
                else
                {
                    currentCapacityA -= plants[i];
                }
                if (currentCapacityB < plants[n - i - 1])
                {
                    result++;
                    currentCapacityB = capacityB - plants[n - i - 1];
                }
                else
                {
                    currentCapacityB -= plants[n - i - 1];
                }
            }

        }

        return result;
    }
};
// @lc code=end

