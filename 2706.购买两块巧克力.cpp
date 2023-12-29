#include <vector>
#include<algorithm>
using namespace std;

class Solution {
    public:
    int buyChoco(vector<int>& prices, int money)
    {
        sort(prices.begin(), prices.end());
        int reside = money - prices[0] - prices[1];
        if (reside>=0)
        {
            return reside;
        }
        return money;
    }
};