/*
 * @lc app=leetcode.cn id=1969 lang=cpp
 *
 * [1969] 数组元素的最小非零乘积
 */

 // @lc code=start
class Solution {
    private:
    long long mod = 1e9 + 7;
    long long fastPow(long long x, long long n)
    {
        long long res = 1;
        for (; n != 0; n >>= 1)
        {
            if (n & 1)
            {
                res = res * x % mod;
            }
            x = x * x % mod;
        }
        return res;
    }

    public:
    int minNonZeroProduct(int p)
    {
        long long maxNum = ((long long)1 << p) - 1;
        return fastPow((maxNum - 1) % mod, ((maxNum - 1) >> 1)) * (maxNum % mod) % mod;
    }
};
// @lc code=end

