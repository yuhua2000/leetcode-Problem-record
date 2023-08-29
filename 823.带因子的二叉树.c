/*
 * @lc app=leetcode.cn id=823 lang=c
 *
 * [823] 带因子的二叉树
 */

 // @lc code=start
int cmp(const void* p1, const void* p2) {
    return *(int*)p1 - *(int*)p2;
}

int numFactoredBinaryTrees(int* arr, int arrSize) {
    qsort(arr, arrSize,sizeof(int), cmp);
    long long* dp = (long long*)malloc(arrSize * sizeof(long long));
    long long res = 0, mod = 1e9 + 7;
    for (int i = 0; i < arrSize; i++) {
        dp[i] = 1;
        for (int left = 0, right = i - 1;left <= right;left++) {
            while (left <= right && (long long)arr[left] * arr[right] > arr[i]) {
                right--;
            }
            if (left <= right && (long long)arr[left] * arr[right] == arr[i]) {
                if (left == right) {
                    dp[i] = (dp[i] + dp[right] * dp[left]) % mod;
                }
                else {
                    dp[i] = (dp[i] + dp[right] * dp[left] * 2) % mod;
                }
            }
        }
        res = (res + dp[i]) % mod;
    }
    return res;
}
// @lc code=end

