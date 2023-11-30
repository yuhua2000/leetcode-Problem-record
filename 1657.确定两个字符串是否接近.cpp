/*
 * @lc app=leetcode.cn id=1657 lang=cpp
 *
 * [1657] 确定两个字符串是否接近
 */
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

// @lc code=start
class Solution {
public:
    bool closeStrings(string word1, string word2) {
        vector<int> count1(26), count2(26);
       for (char c: word1)
       {
           count1[c - 'a']++;
       }
       for (char c : word2)
       {
           count2[c - 'a']++;
       }

       for (size_t i = 0; i < 26; i++)
       {
           if ((count1[i] == 0 && count2[i] > 0) || (count1[i] > 0 && count2[i] == 0))
            {
                return false;
            }
       }
       sort(count1.begin(),count1.end());
       sort(count2.begin(), count2.end());
       return count1 == count2;
    }
};
// @lc code=end

