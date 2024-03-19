/*
 * @lc app=leetcode.cn id=1686 lang=cpp
 *
 * [1686] 石子游戏 VI
 */

#include <vector>
#include <algorithm>

using namespace std;

// @lc code=start
class Solution {
    public:
    int stoneGameVI(vector<int>& aliceValues, vector<int>& bobValues)
    {
        int n = aliceValues.size();
        vector<pair<int, int>> values;

        for (int i = 0; i < aliceValues.size(); i++)
            values.push_back(make_pair(aliceValues[i] + bobValues[i], i));

        sort(values.begin(), values.end(), [] (pair<int, int>& a, pair<int, int>& b) {
            return a.first > b.first;
            }
        );

        int aliceSum = 0, bobSum = 0;
        for (int i = 0; i < n; i++)
            if (i % 2 == 0)
                aliceSum += aliceValues[values[i].second];
            else
                bobSum += bobValues[values[i].second];


        if (aliceSum > bobSum)
            return 1;
        else if (aliceSum == bobSum)
            return 0;
        else
            return -1;

    }
};

// @lc code=end

