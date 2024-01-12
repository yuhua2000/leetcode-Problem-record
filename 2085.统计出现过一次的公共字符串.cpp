/*
 * @lc app=leetcode.cn id=2085 lang=cpp
 *
 * [2085] 统计出现过一次的公共字符串
 */

#include <vector>
#include <string>
#include <unordered_map>

using namespace std;

// @lc code=start
class Solution {
    public:
    int countWords(vector<string>& words1, vector<string>& words2)
    {
        unordered_map<string, int> words;
        for (int i = 0; i < words1.size(); i++)
        {
            words[words1[i]]++;
        }

        for (auto it = words.begin();it != words.end();)
        {
            if (it->second != 1)
            {
                it = words.erase(it);
            }
            else
            {
                ++it;
            }
        }

        for (int i = 0; i < words2.size(); i++)
        {
            if (words[words2[i]] >= 1)
            {
                words[words2[i]]++;
            }
        }
        
        int result = 0;
        for (auto it = words.begin();it != words.end();it++)
        {
            if (it->second == 2)
            {
                result++;
            }
        }
        return result;
    }
};
// @lc code=end

