#include <string>
#include <vector>
using namespace std;
/*
 * @lc app=leetcode.cn id=1410 lang=cpp
 *
 * [1410] HTML 实体解析器
 */

 // @lc code=start
class Solution {
    public:
    using EntityChar = pair<string, char>;
    vector<EntityChar> entityList;
    string entityParser(string text)
    {
        vector<EntityChar> entityList = {
            { "&quot;", '"' },
            { "&apos;", '\'' },
            { "&amp;", '&' },
            { "&gt;", '>' },
            { "&lt;", '<' },
            { "&frasl;", '/' }
        };

        string r = "";
        for (int pos = 0; pos < text.size();)
        {
            bool isEntity = false;
            if (text[pos]=='&')
            {
                for (const auto &[e,c]:entityList)
                {
                    if (text.substr(pos,e.size())==e)
                    {
                        r.push_back(c);
                        pos += e.size();
                        isEntity = true;
                        break;
                    }
                    
                }
            }
            if (!isEntity)
            {
                r.push_back(text[pos]);
                pos++;
                continue;
            }
        }
        return r;
    }
};
// @lc code=end

