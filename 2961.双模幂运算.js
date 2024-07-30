/**
 * @param {number[][]} variables
 * @param {number} target
 * @return {number[]}
 */

function binpow(a, b, mod) {
    let res = 1;
    while (b > 0) {
        if (b & 1) {
            res = (res * a) % mod;
        }
        a = (a * a) % mod;
        b >>= 1;
    }
    return res;
}

var getGoodIndices = function (variables, target) {
    let res = Array();
    for (let i = 0; i < variables.length; i++) {
        if (binpow(binpow(variables[i][0], variables[i][1], 10), variables[i][2], variables[i][3]) == target) {
            res.push(i);
        }
    }
    return res;
};