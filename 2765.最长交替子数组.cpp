#include <vector>

using namespace std;

class Solution {
    public:
    int alternatingSubarray(vector<int>& nums) {
        int index = 0;
        int alter = 1;
        int res = -1;
        for (int i = 1 ; i < nums.size()  ; i++)
        {
            int length = i - index + 1;
            if (nums[i] - nums[index] == (length - 1) % 2)
            {
                 res = max(res, length);   
            }
            else
            {
                 if (nums[i] - nums[i - 1] == 1) {
                    index = i - 1;
                    res = max(res, 2);
                } else {
                    index = i;
                } 
            }
        }
        return res;
    }
};