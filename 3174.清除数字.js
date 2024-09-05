/**
 * @param {string} s
 * @return {string}
 */
var clearDigits = function (s) {
    let result = "";
    for (let i = 0; i < s.length; i++) {
        if (s[i] >= '0' && s[i] <= '9') {
            result = result.substr(0, result.length - 1);
        } else {
            result += s[i];
        }
    }
    return result;
};