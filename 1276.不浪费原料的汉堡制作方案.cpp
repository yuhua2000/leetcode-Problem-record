#include <vector>

using namespace std;
/*
 * @lc app=leetcode.cn id=1276 lang=cpp
 *
 * [1276] 不浪费原料的汉堡制作方案
 */

 // @lc code=start
class Solution {
    public:
    vector<int> numOfBurgers(int tomatoSlices, int cheeseSlices)
    {
        if (tomatoSlices % 2 != 0 || tomatoSlices<cheeseSlices * 2 || tomatoSlices>cheeseSlices * 4)
        {
            return {};
        }
        return { tomatoSlices / 2 - cheeseSlices,cheeseSlices * 2 - tomatoSlices / 2 };
    }
};
// @lc code=end

