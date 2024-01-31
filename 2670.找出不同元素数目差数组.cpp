#include <vector>
#include <unordered_set>
using namespace std;

class Solution {
    public:
    vector<int> distinctDifferenceArray(vector<int>& nums)
    {

        unordered_set<int> diff;
        vector<int> sufDiff(nums.size() + 1);
        vector<int> result(nums.size());
        for (int i = nums.size() - 1; i >= 0; i--)
        {
            diff.emplace(nums[i]);
            sufDiff[i] = diff.size();
        }

        diff.clear();
        for (int i = 0; i < nums.size() - 1; i++)
        {
            diff.emplace(nums[i]);
            result[i] = diff.size() - sufDiff[i + 1];
        }

        return result;
    }
};
