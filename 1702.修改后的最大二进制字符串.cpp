/*
 * @lc app=leetcode.cn id=1702 lang=cpp
 *
 * [1702] 修改后的最大二进制字符串
 */
#include <string>

using namespace std;

// @lc code=start
class Solution {
    public:
    string maximumBinaryString(string binary)
    {
        int n = binary.size();
        int j = 0;
        for (int index = 0; index < n;index++)
        {
            if (binary[index] == '0')
            {
                while (j <= index || (j < n && binary[j] == '1'))
                {
                    j++;
                }
                if (j < n)
                {
                    binary[j] = '1';
                    binary[index] = '1';
                    binary[index + 1] = '0';
                }
            }

        }
        return binary;
    }
};
// @lc code=end

