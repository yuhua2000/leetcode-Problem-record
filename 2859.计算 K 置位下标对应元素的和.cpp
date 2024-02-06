#include <vector>

using namespace std;

class Solution {
    public:
    int sumIndicesWithKSetBits(vector<int>& nums, int k)
    {
        int result = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int num = i;
            int bits = 0;
            while (num > 0)
            {
                bits++;
                num &= (num - 1);
            }
            if (bits == k)
            {
                result += nums[i];
            }
        }
        return result;
    }
};
