/**
 * @param {number[]} nums
 * @param {number[]} moveFrom
 * @param {number[]} moveTo
 * @return {number[]}
 */
var relocateMarbles = function (nums, moveFrom, moveTo) {
    const numsMap = new Map();
    for (let i = 0; i < nums.length; i++) {
        numsMap.set(nums[i], true);
    }

    for (let i = 0; i < moveFrom.length; i++) {
        numsMap.delete(moveFrom[i]);
        numsMap.set(moveTo[i]);
    }

    const result = new Array();
    numsMap.forEach(function (_, key) {
        result.push(key);
    });

    result.sort(function (x, y) {
        return x - y;
    });
    return result;
};