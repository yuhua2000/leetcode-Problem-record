/*
 * @lc app=leetcode.cn id=1017 lang=cpp
 *
 * [1017] 负二进制转换
 */
#include <string>
#include <iostream>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
    public:
    string baseNeg2(int n)
    {
        if (n == 0)
        {
            return "0";
        }

        vector<int> bits(32);
        for (size_t i = 0; i < 32 && n != 0; i++)
        {
            if (n & 1)
            {
                bits[i]++;
                if (i & 1)
                {
                    bits[i + 1]++;
                }
            }
            n >>= 1;
        }
        int carry = 0;
        for (size_t i = 0; i < 32; i++)
        {
            int var = carry + bits[i];
            bits[i] = var & 1;
            carry = (var - bits[i]) / (-2);

        }
        int pos = 31;
        string result;
        while (pos >= 0 && bits[pos] == 0)
        {
            pos--;
        }

        while (pos >= 0)
        {
            result.push_back(bits[pos] + '0');
            pos--;
        }

        return result;
    }
};
// @lc code=end

