#include <vector>
#include <stack>

using namespace std;

class Solution {
    public:
    long long maximumSumOfHeights(vector<int>& maxHeights)
    {
        int n = maxHeights.size();
        long long result = 0;
        vector<long long> prefix(n), suffix(n);
        stack<int> stack1, stack2;
        for (int i = 0; i < n; i++)
        {
            while (!stack1.empty() && maxHeights[stack1.top()] > maxHeights[i])
            {
                stack1.pop();
            }
            if (stack1.empty())
            {
                prefix[i] = (long long)(i + 1) * maxHeights[i];
            }
            else
            {
                prefix[i] = prefix[stack1.top()] + (long long)(i - stack1.top()) * maxHeights[i];
            }
            stack1.emplace(i);
        }

        for (int i = n - 1; i >= 0; i--)
        {
            while (!stack2.empty() && maxHeights[stack2.top()] > maxHeights[i])
            {
                stack2.pop();
            }
            if (stack2.empty())
            {
                suffix[i] = (long long)(n - i) * maxHeights[i];
            }
            else
            {
                suffix[i] = suffix[stack2.top()] + (long long)(stack2.top() - i) * maxHeights[i];
            }
            stack2.emplace(i);
            result = max(result, prefix[i] + suffix[i] - maxHeights[i]);
        }

        return result;
    }
};
