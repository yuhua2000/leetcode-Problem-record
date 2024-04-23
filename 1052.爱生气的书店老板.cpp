/*
 * @lc app=leetcode.cn id=1052 lang=cpp
 *
 * [1052] 爱生气的书店老板
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    int maxSatisfied(vector<int>& customers, vector<int>& grumpy, int minutes)
    {
        int sum = 0;
        int maxDiscontentCustomers = 0;
        int discontentCustomers = 0;

        for (size_t i = 0; i < minutes; i++)
        {
            if (!grumpy[i])
                sum += customers[i];

            if (grumpy[i])
                discontentCustomers += customers[i];
        }
        maxDiscontentCustomers = discontentCustomers;
        for (size_t i = minutes; i < customers.size(); i++)
        {
            if (!grumpy[i])
                sum += customers[i];
            if (grumpy[i])
            {
                discontentCustomers += customers[i];
            }
            if (grumpy[i - minutes])
            {
                discontentCustomers -= customers[i - minutes];
            }
            maxDiscontentCustomers = max(maxDiscontentCustomers, discontentCustomers);
        }

        return sum + maxDiscontentCustomers;
    }
};
// @lc code=end

