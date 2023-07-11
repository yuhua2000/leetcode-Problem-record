/*
 * @lc app=leetcode.cn id=1911 lang=c
 *
 * [1911] 最大子序列交替和
 */

// @lc code=start
long long maxAlternatingSum(int* nums, int numsSize) {
    long long even = nums[0], odd = 0;
    for (int i = 0; i < numsSize; i++) {
        even = fmax(even, odd + nums[i]);
        odd = fmax(odd, even - nums[i]);
    }
    return even;
}
// @lc code=end

