/*
 * @lc app=leetcode.cn id=1103 lang=cpp
 *
 * [1103] 分糖果 II
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    vector<int> distributeCandies(int candies, int num_people)
    {
        vector<int> result(num_people);
        int i = 1;
        while (candies > 0)
        {
            result[(i - 1) % num_people] += i;
            candies -= i;
            i++;
        }
        if (candies < 0)
        {
            result[(i - 2) % num_people] += candies;
        }
        return result;
    }
};
// @lc code=end

