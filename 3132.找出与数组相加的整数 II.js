/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var minimumAddedInteger = function (nums1, nums2) {
    nums1.sort((a, b) => { return a - b })
    nums2.sort((a, b) => { return a - b })
    const diffMap = new Set();
    for (let i = 0; i < 3; i++) {
        diffMap.add(nums2[0] - nums1[i])
    }

    for (let i = 1; i < nums2.length; i++) {
        const diff = new Set()
        for (let j = i - 2; j <= i + 2; j++) {
            if (j >= 0 && j < nums1.length) {
                diff.add(nums2[i] - nums1[j])
            }
        }
        for (let value of diffMap) {
            if (!diff.has(value)) {
                diffMap.delete(value)
            }
        }
    }


    for (let diff of Array.from(diffMap).sort((a, b) => a - b)) {
        let idx = 0;
        for (let j = 0; j < nums2.length; j++) {
            let num = nums2[j] - diff;
            for (; idx < nums1.length; idx++) {
                if (nums1[idx] == num) {
                    if (j != nums2.length - 1) {
                        idx++
                    }
                    break
                }
            }
            if (idx == num.length) {
                break;
            }
        }
        if (idx < nums1.length) {
            return diff
        }
    }

    return -1;
};

var minimumAddedInteger2 = function (nums1, nums2) {
    nums1.sort((a, b) => { return a - b })
    nums2.sort((a, b) => { return a - b })
    for (let diff of [nums2[0] - nums1[2], nums2[0] - nums1[1], nums2[0] - nums1[0]]) {
        let idx = 0;
        for (let j = 0; j < nums2.length; j++) {
            let num = nums2[j] - diff;
            for (; idx < nums1.length; idx++) {
                if (nums1[idx] == num) {
                    if (j != nums2.length - 1) {
                        idx++
                    }
                    break
                }
            }
            if (idx == num.length) {
                break;
            }
        }
        if (idx < nums1.length) {
            return diff
        }
    }

    return -1;
};

// [3, 4, 4, 9, 9] 总和
// [7, 8, 8] 总和
