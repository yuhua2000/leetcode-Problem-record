/*
 * @lc app=leetcode.cn id=670 lang=cpp
 *
 * [670] 最大交换
 */
#include <string>
#include <stdexcept>

using namespace std;

// @lc code=start
class Solution {
    public:
    int maximumSwap1(int num)
    {
        string numStr = to_string(num);
        for (int i = 0; i < numStr.length() - 1; i++)
        {
            int  maxIndex = i + 1;
            for (int j = i + 1; j < numStr.length(); j++)
            {
                if (numStr[j] >= numStr[maxIndex])
                {
                    maxIndex = j;
                }
            }

            if (numStr[i] < numStr[maxIndex])
            {
                char temp = numStr[i];
                numStr[i] = numStr[maxIndex];
                numStr[maxIndex] = temp;
                return stoi(numStr);
            }

        }
        return num;
    }

    int maximumSwap(int num)
    {
        string numStr = to_string(num);
        int n = numStr.length();
        int maxIdx = n - 1;
        int idx1 = -1, idx2 = -1;
        for (int i = n-1; i>=0; i--)
        {
            if (numStr[i]>numStr[maxIdx])
            {
                maxIdx = i;
            }
            else if (numStr[i] < numStr[maxIdx])
            {
                idx1 = i;
                idx2 = maxIdx;
            }   
        }

        if (idx1>=0)
        {
            swap(numStr[idx1], numStr[idx2]);
            return stoi(numStr);
        }
        
        return num;
    }
};
// @lc code=end

