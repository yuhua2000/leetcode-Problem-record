#include <vector>
#include <algorithm>

using namespace std;

class Solution {
    public:
    int countPairs1(vector<int>& nums, int target) {
        int result = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int tempTarget = target - nums[i];
            for (int j = i + 1; j < nums.size(); j++)
            {
                if (nums[j] < tempTarget)
                    result++;
            }
        }
        return result;
    }
    
    int countPairs2(vector<int>& nums, int target)
    {
        sort(nums.begin(), nums.end());
        int result = 0;
        for (int i = 0, j = nums.size() - 1; i < j; i++)
        {
            while (i < j && nums[i] + nums[j] >= target)
            {
                j--;
            }
            result += j - i;
        }
        return result;
    }

    int countPairs(vector<int>& nums, int target) {
        sort(nums.begin(), nums.end());
        int result = 0;
        for (int i = 1; i < nums.size(); i++)
        {
            result += lower_bound(nums.begin(), nums.begin() + i, target - nums[i]) - nums.begin();
        }
        return result;
    }
};