/*
 * @lc app=leetcode.cn id=303 lang=cpp
 *
 * [303] 区域和检索 - 数组不可变
 */
#include <vector>

using namespace std;

// @lc code=start
class NumArray {
    private:
    vector<int> sums;
    public:
    NumArray(vector<int>& nums) {
        int sum = 0;
        for (auto i : nums)
        {
            sum += i;
            sums.push_back(sum);
        }
        
    }
    
    int sumRange(int left, int right) {
        if (left==0)
        {
            return sums[right];
        }
        return sums[right] - sums[left - 1];
    }
};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray* obj = new NumArray(nums);
 * int param_1 = obj->sumRange(left,right);
 */
// @lc code=end

