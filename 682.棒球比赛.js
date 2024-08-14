/*
 * @lc app=leetcode.cn id=682 lang=javascript
 *
 * [682] 棒球比赛
 */

// @lc code=start
/**
 * @param {string[]} operations
 * @return {number}
 */
var calPoints = function(operations) {
    let points = new Array();
    for (let i  = 0; i   < operations.length; i++) {
        if (operations[i] == 'C'){
            points.pop();
        } else if (operations[i] == 'D') {
            points.push(points[points.length - 1] * 2);
        } else if (operations[i] == '+') {
            points.push(points[points.length - 1]+points[points.length - 2])
        } else {
            points.push(operations[i] - '0');
        }
    }

    let result = 0;
    points.forEach(function (v, _) {
        result += v;
    })
    return result;
};
// @lc code=end

