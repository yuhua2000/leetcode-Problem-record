/**
 * @param {number[]} nums
 * @return {number}
 */
var findValueOfPartition = function (nums) {
    nums.sort(function (x, y) {
        return x - y;
    })
    let minPartition = Math.abs(nums[1] - nums[0]);
    for (let i = 2; i < nums.length; i++) {
        minPartition = Math.min(Math.abs(nums[i] - nums[i - 1]), minPartition);
    }
    return minPartition;
};