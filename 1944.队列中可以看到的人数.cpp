#include <vector>

using namespace std;

/*
 * @lc app=leetcode.cn id=1944 lang=cpp
 *
 * [1944] 队列中可以看到的人数
 */

// @lc code=start
class Solution {
public:
    vector<int> canSeePersonsCount(vector<int>& heights) {
        int n = heights.size();
        vector<int> stack;
        vector<int> res(n, 0);
        for (int i = n-1; i >=0; i--)
        {
            int h = heights[i];
            while (!stack.empty()&&stack.back()<h)
            {
                stack.pop_back();
                res[i]++;
            }
            if (!stack.empty())
            {
                res[i]++;
            }
            stack.push_back(h);
        }

        return res;
    }
};
// @lc code=end

