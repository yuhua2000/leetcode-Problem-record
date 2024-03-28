/*
 * @lc app=leetcode.cn id=1997 lang=cpp
 *
 * [1997] 访问完所有房间的第一天
 */
#include <vector>
#include <set>
using namespace std;

// @lc code=start
class Solution {
    public:
    // int firstDayBeenInAllRooms(vector<int>& nextVisit) {
    //     vector<int> visitNums(nextVisit.size());

    //     set<int> count;
    //     int result = 0;
    //     int mod = 1e9 + 7;

    //     int nowIndex = 0;
    //     count.insert(nowIndex);
    //     while (count.size() < nextVisit.size())
    //     {
    //         visitNums[nowIndex]++;
    //         if (visitNums[nowIndex]%2==1)
    //         {
    //             nowIndex = nextVisit[nowIndex];
    //         }
    //         else
    //         {
    //             nowIndex = (nowIndex + 1) % nextVisit.size();
    //         }
    //         result = (result + 1) % mod;
    //         count.insert(nowIndex);
    //     }
    //     return result;
    // }
    int firstDayBeenInAllRooms(vector<int>& nextVisit)
    {
        int mod = 1e9 + 7;
        int n = nextVisit.size();
        vector<long> s(n);
        for (int i = 0; i < n-1; i++)
        {
            int j = nextVisit[i];
            s[i + 1] = (s[i] * 2 - s[j] + 2 + mod) % mod;
        }
        
        return s[n - 1];
    }
};