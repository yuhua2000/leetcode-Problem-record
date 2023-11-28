/*
 * @lc app=leetcode.cn id=1386 lang=cpp
 *
 * [1386] 安排电影院座位
 */

#include <vector>
#include <map>
#include <unordered_map>
using namespace std;


// @lc code=start
class Solution {
    public:
    // 超时
    int maxNumberOfFamilies1(int n, vector<vector<int>>& reservedSeats)
    {
        int result = 0;
        map<int, bool> reservedSeatsMap;
        for (size_t i = 0; i < reservedSeats.size(); i++)
        {
            reservedSeatsMap[reservedSeats[i][0] * 10 + reservedSeats[i][1]] = true;

        }

        for (size_t i = 1; i <= n; i++)
        {

            if (!reservedSeatsMap[i * 10 + 4] && !reservedSeatsMap[i * 10 + 5] &&
                !reservedSeatsMap[i * 10 + 6] && !reservedSeatsMap[i * 10 + 7])
            {
                if (!reservedSeatsMap[i * 10 + 2] && !reservedSeatsMap[i * 10 + 3] &&
                    !reservedSeatsMap[i * 10 + 8] && !reservedSeatsMap[i * 10 + 9])
                {
                    result += 2;
                }
                else
                {
                    result++;
                }
            }
            else if (!reservedSeatsMap[i * 10 + 2] && !reservedSeatsMap[i * 10 + 3] &&
                !reservedSeatsMap[i * 10 + 4] && !reservedSeatsMap[i * 10 + 5])
            {
                result++;
            }
            else if (!reservedSeatsMap[i * 10 + 6] && !reservedSeatsMap[i * 10 + 7] &&
                !reservedSeatsMap[i * 10 + 8] && !reservedSeatsMap[i * 10 + 9])
            {
                result++;
            }
        }
        return result;
    }

    int maxNumberOfFamilies(int n, vector<vector<int>>& reservedSeats)
    {
        int left = 0b11110000;
        int middle = 0b11000011;
        int right = 0b00001111;

        unordered_map<int, int> occupied;
        for (const vector<int>& seat : reservedSeats)
        {
            if (seat[1] >= 2 && seat[1] <= 9)
            {
                occupied[seat[0]] |= (1 << (seat[1] - 2));
            }
        }
        int ans = (n - occupied.size()) * 2;

        for (auto &[row,bitmsk]:occupied)
        {
            if (((bitmsk|left)==left)||((bitmsk|middle)==middle)||((bitmsk|right)==right))
            {
                ans++;
            }
        }
        return ans;
    }
};
// @lc code=end

