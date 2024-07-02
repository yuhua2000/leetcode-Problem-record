#include <vector>
#include <unordered_set>

using namespace std;

class Solution {
    public:
    int maximumPrimeDifference(vector<int>& nums)
    {
        unordered_set<int> primes = {
         2, 3, 5, 7, 11,
         13, 17, 19, 23, 29,
         31, 37, 41, 43, 47,
         53, 59, 61, 67, 71,
         73, 79, 83, 89, 97
        };

        int first = -1, end = -1;
        int n = nums.size();
        for (int i = 0; i < n; i++)
        {
            if (first == -1 && primes.count(nums[i]))
                first = i;

            if (end == -1 && primes.count(nums[n - i - 1]))
                end = n - i - 1;

            if (first != -1 && end != -1)
                break;

        }

        return end - first;

    }
};