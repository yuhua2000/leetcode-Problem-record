#include <string>
#include <vector>

using namespace std;

class Solution {
    public:
    int maximumLength(string s)
    {
        vector<int> chs[26];
        int n = s.size();
        for (int i = 0, cnt = 0; i < n; i++)
        {
            cnt++;
            if (i + 1 == n || s[i + 1] != s[i])
            {
                int ch = s[i] - 'a';
                chs[ch].push_back(cnt);
                cnt = 0;
                for (int j = chs[ch].size() - 1; j > 0; j--)
                {
                    if (chs[ch][j] > chs[ch][j - 1])
                        swap(chs[ch][j], chs[ch][j - 1]);
                    else
                        break;
                }
                if (chs[ch].size() > 3)
                    chs[ch].pop_back();
            }
        }

        int result = -1;

        for (int i = 0; i < 26; i++)
        {
            vector<int> ch = chs[i];
            switch (ch.size())
            {
                case 3:
                    result = max(result, ch[2]);
                case 2:
                    if (ch[0] > 1)
                        result = max(result, min(ch[0] - 1, ch[1]));
                case 1:
                    result = max(result, ch[0] - 2);
            }
        }

        return result > 0 ? result : -1;
    }
};

/*
给你一个仅由小写英文字母组成的字符串 s 。

如果一个字符串仅由单一字符组成，那么它被称为 特殊 字符串。例如，字符串 "abc" 不是特殊字符串，而字符串 "ddd"、"zz" 和 "f" 是特殊字符串。

返回在 s 中出现 至少三次 的 最长特殊子字符串 的长度，如果不存在出现至少三次的特殊子字符串，则返回 -1 。

子字符串 是字符串中的一个连续 非空 字符序列。



示例 1：

输入：s = "aaaa"
输出：2
解释：出现三次的最长特殊子字符串是 "aa" ：子字符串 "aaaa"、"aaaa" 和 "aaaa"。
可以证明最大长度是 2 。
示例 2：

输入：s = "abcdef"
输出：-1
解释：不存在出现至少三次的特殊子字符串。因此返回 -1 。
示例 3：

输入：s = "abcaba"
输出：1
解释：出现三次的最长特殊子字符串是 "a" ：子字符串 "abcaba"、"abcaba" 和 "abcaba"。
可以证明最大长度是 1 。


提示：

3 <= s.length <= 5 * 105
s 仅由小写英文字母组成。
*/