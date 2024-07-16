/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var findIntersectionValues = function (nums1, nums2) {
    var exitsNum = new Array(202);
    var result = [0, 0];
    for (let i = 0; i < nums1.length; i++) {
        exitsNum[nums1[i]] = true;
    }

    for (let i = 0; i < nums2.length; i++) {
        exitsNum[nums2[i] + 101] = true;

        if (exitsNum[nums2[i]]) {
            result[1]++;
        }
    }

    for (let i = 0; i < nums1.length; i++) {
        if (exitsNum[nums1[i] + 101]) {
            result[0]++;
        }
    }

    return result;
};