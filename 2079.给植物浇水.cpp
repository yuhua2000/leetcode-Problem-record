/*
 * @lc app=leetcode.cn id=2079 lang=cpp
 *
 * [2079] 给植物浇水
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
public:
    int wateringPlants(vector<int>& plants, int capacity) {
        int result = 0;
        int loc = -1;
        int currentCapacity = capacity;
        for (int i = 0; i < plants.size(); i++)
        {
            if (currentCapacity>=plants[i])
            {
                result += i - loc;
                loc = i; 
                currentCapacity -= plants[i];
            }
            else
            {
                result += (loc - (-1)) + ((i - (-1)));
                loc = i;
                currentCapacity = capacity - plants[i];
            }
        }
        return result;
    }
};
// @lc code=end

