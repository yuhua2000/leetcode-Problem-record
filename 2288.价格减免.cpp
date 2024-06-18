/*
 * @lc app=leetcode.cn id=2288 lang=cpp
 *
 * [2288] 价格减免
 */
#include <string>
#include <iostream>
#include <sstream>   
#include <iomanip>

using namespace std;

// @lc code=start
class Solution {
    public:
    string discountPrices(string sentence, int discount)
    {
        stringstream sout;
        sout << fixed << setprecision(2);
        bool isPrice = false;
        string word;
        long long price = 0;
        for (size_t i = 0; i < sentence.size(); i++)
        {
            word += sentence[i];
            if (sentence[i] == ' ' || i == sentence.size() - 1)
            {
                if (isPrice)
                {
                    if (i == sentence.size() - 1)
                    {
                        if (sentence[i] >= '0' && sentence[i] <= '9')
                            price = price * 10 + sentence[i] - '0';
                        else
                            isPrice = false;
                    }
                }
                if (isPrice && price > 0)
                {
                    sout << "$" << double(price) * (1.0 - discount / 100.0);
                    if (i != sentence.size() - 1)
                        sout << " ";
                }
                else
                {
                    sout << word;
                }
                word = "";
                isPrice = false;
            }
            else if (sentence[i] == '$' && (i == 0 || sentence[i - 1] == ' '))
            {
                isPrice = true;
                price = 0;
            }
            else if (isPrice && sentence[i] >= '0' && sentence[i] <= '9')
            {
                price = price * 10 + sentence[i] - '0';
            }
            else if (isPrice)
            {
                isPrice = false;
            }
        }
        if (!word.empty())
            sout << word;

        return sout.str();
    }
};
// @lc code=end

