/**
 * @param {number[]} possible
 * @return {number}
 */
var minimumLevels = function (possible) {
    var sum = 0;
    for (let i = 0; i < possible.length; i++) {
        sum += possible[i] == 1 ? 1 : -1;
    }
    var aliceScore = 0;
    for (let i = 0; i < possible.length - 1; i++) {
        aliceScore += possible[i] == 1 ? 1 : -1;
        if (aliceScore > sum - aliceScore) {
            return i + 1;
        }
    }
    return -1;
};