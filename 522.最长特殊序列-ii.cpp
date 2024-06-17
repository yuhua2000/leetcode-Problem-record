/*
 * @lc app=leetcode.cn id=522 lang=cpp
 *
 * [522] 最长特殊序列 II
 */
#include <vector>
#include <string>

using namespace std;

// @lc code=start
class Solution {
    public:
    int findLUSlength(vector<string>& strs)
    {
        auto is_subseq = [] (const string& s, const string& t) -> bool {
            int pt_s = 0, pt_t = 0;
            while (pt_s < s.size() && pt_t < t.size())
            {
                if (s[pt_s] == t[pt_t])
                    ++pt_s;
                ++pt_t;
            }

            return pt_s == s.size();
        };

        int n = strs.size();
        int ans = -1;
        for (int i = 0; i < n; i++)
        {
            bool check = true;
            for (int j = 0; j < n; j++)
            {

                if (i != j && is_subseq(strs[i], strs[j]))
                {
                    check = false;
                    break;
                }

            }
            if (check)
                ans = max(ans, static_cast<int>(strs[i].size()));

        }

        return ans;
    }
};
// @lc code=end

/*
对于给定的某个字符串 str[i]\textit{str}[i]str[i]，如果它的一个子序列 sub\textit{sub}sub 是「特殊序列」，那么 str[i]\textit{str}[i]str[i] 本身也是一个「特殊序列」。

这是因为如果 sub\textit{sub}sub 没有在除了 str[i]\textit{str}[i]str[i] 之外的字符串中以子序列的形式出现过，那么给 sub\textit{sub}sub 不断地添加字符，它也不会出现。而 str[i]\textit{str}[i]str[i] 就是 sub\textit{sub}sub 添加若干个（可以为零个）字符得到的结果。

因此我们只需要使用一个双重循环，外层枚举每一个字符串 str[i]\textit{str}[i]str[i] 作为特殊序列，内层枚举每一个字符串 str[j] (i≠j)\textit{str}[j]~(i \neq j)str[j] (i

=j)，判断 str[i]\textit{str}[i]str[i] 是否不为 str[j]\textit{str}[j]str[j] 的子序列即可。

要想判断 str[i]\textit{str}[i]str[i] 是否为 str[j]\textit{str}[j]str[j] 的子序列，我们可以使用贪心 + 双指针的方法：即初始时指针 pti\textit{pt}_ipt
i
​
  和 ptj\textit{pt}_jpt
j
​
  分别指向两个字符串的首字符。如果两个字符相同，那么两个指针都往右移动一个位置，表示匹配成功；否则只往右移动指针 ptj\textit{pt}_jpt
j
​
 ，表示匹配失败。如果 pti\textit{pt}_ipt
i
​
  遍历完了整个字符串，就说明 str[i]\textit{str}[i]str[i] 是 str[j]\textit{str}[j]str[j] 的子序列。

在所有满足要求的 str[i]\textit{str}[i]str[i] 中，我们选出最长的那个，返回其长度作为答案。如果不存在满足要求的字符串，那么返回 −1-1−1。


*/