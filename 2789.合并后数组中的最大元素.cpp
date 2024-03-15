#include <vector>

using namespace std;

class Solution {
    public:
    long long maxArrayValue(vector<int>& nums)
    {
        long long curr = nums[nums.size() - 1];
        long long maxValue = nums[nums.size() - 1];
        for (int i = nums.size() - 2; i >= 0; i--)
        {
            if (curr>nums[i])
            {
                curr = curr + nums[i];
            }
            else
            {
                curr = nums[i];
            }
            maxValue = max(maxValue, curr);
        }
        
        return maxValue;
    }
};