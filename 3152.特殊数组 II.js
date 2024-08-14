/**
 * @param {number[]} nums
 * @param {number[][]} queries
 * @return {boolean[]}
 */
var isArraySpecial1 = function (nums, queries) {
    let evenSpecialStr = [];
    for (let i = 0; i < nums.length; i++) {
        nums[i] %= 2;
        evenSpecialStr.push(i % 2);
    }

    let evenSpecialStrJoined = evenSpecialStr.join("");
    let numStr = nums.join("");

    let result = new Array(queries.length);
    for (let i = 0; i < queries.length; i++) {
        let subNumStr = numStr.substring(queries[i][0], queries[i][1] + 1);
        result[i] = evenSpecialStrJoined.startsWith(subNumStr)
        if (subNumStr.startsWith("10")) {
            result[i] |= evenSpecialStrJoined.startsWith(subNumStr.substring(1))
        }
    }

    return result;
};

var isArraySpecial = function (nums, queries) {
    const n = nums.length;
    const dp = new Array(n).fill(1);
    for (let i = 1; i < nums.length; i++) {
        if ((nums[i - 1] + nums[i]) % 2 == 1) {
            dp[i] = dp[i - 1] + 1
        }
    }

    const res = new Array(queries.length);
    for (let i = 0; i < queries.length; i++) {
        res[i] = dp[queries[i][1]] >= queries[i][1] - queries[i][0] + 1
    }
};