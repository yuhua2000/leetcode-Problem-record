/**
 * @param {number[]} nums
 * @return {number}
 */
var maxStrength = function (nums) {
    nums.sort(function (x, y) {
        return y - x;
    });

    let result = 1
    let resultEffective = false;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] > 0) {
            result *= nums[i];
            resultEffective = true;
        } else {
            break
        }
    }

    for (let i = nums.length - 1; i > 0; i--) {
        if (nums[i] < 0 && i - 1 >= 0 && nums[i - 1] < 0) {
            result *= nums[i] * nums[i - 1];
            i--
            resultEffective = true;
        } else {
            break
        }
    }

    if (!resultEffective) {
        return nums[0]
    }
    
    return result
};