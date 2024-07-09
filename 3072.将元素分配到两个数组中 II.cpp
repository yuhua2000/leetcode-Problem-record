#include <vector>
#include <algorithm>
#include <iostream>

using namespace std;

class Solution {
    public:
    vector<int> resultArray(vector<int>& nums)
    {
        vector<int> nums1{ nums[0] };
        vector<int> nums2{ nums[1] };
        vector<int> sortNums1 = nums1;
        vector<int> sortNums2 = nums2;
        for (size_t i = 2; i < nums.size(); i++)
        {
            auto num = nums[i];
            auto it = std::lower_bound(sortNums1.begin(), sortNums1.end(), num + 1);
            int sum1 = std::distance(it, sortNums1.end());

            it = std::lower_bound(sortNums2.begin(), sortNums2.end(), num + 1);
            int sum2 = std::distance(it, sortNums2.end());

            cout << sum1 << sum2 << nums1.size() << nums2.size() << sortNums1.size() << sortNums2.size() << endl;
            if (sum1 > sum2 || (sum1 == sum2 && nums1.size() <= nums2.size()))
            {
                nums1.emplace_back(num);
                auto it = lower_bound(sortNums1.begin(), sortNums1.end(), num);
                sortNums1.insert(it, num);

            }
            else
            {
                nums2.emplace_back(num);
                auto it = lower_bound(sortNums2.begin(), sortNums2.end(), num);
                sortNums2.insert(it, num);
            }
        }

        nums1.insert(nums1.end(), nums2.begin(), nums2.end());
        return nums1;
    }
};