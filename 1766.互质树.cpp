/*
 * @lc app=leetcode.cn id=1766 lang=cpp
 *
 * [1766] 互质树
 */
#include <vector>
#include <numeric>

using namespace std;

// @lc code=start

class Solution {
    private:
    vector<vector<int>> gcds;
    vector<int> dep;
    vector<vector<int>> tmp;
    vector<vector<int>> g;
    vector<int> ans;

    void dfs(vector<int>& nums, int x, int depth)
    {
        dep[x] = depth;
        for (auto val : gcds[nums[x]])
        {
            if (tmp[val].empty())
                continue;

            int las = tmp[val].back();
            if (ans[x] == -1 || dep[las] > dep[ans[x]])
                ans[x] = las;

        }
        tmp[nums[x]].emplace_back(x);
        for (auto i : g[x])
            if (dep[i] == -1)
                dfs(nums, i, depth + 1);



        tmp[nums[x]].pop_back();

    }

    public:
    vector<int> getCoprimes(vector<int>& nums, vector<vector<int>>& edges)
    {
        int n = nums.size();
        gcds.resize(51);
        tmp.resize(51);
        ans.resize(n, -1);
        dep.resize(n, -1);
        g.resize(n);

        for (int i = 1; i < 51; i++)
            for (int j = 1  ; j < 51; j++)
                if (gcd(i, j) == 1)
                    gcds[i].emplace_back(j);

        for (auto& edge : edges)
        {
            g[edge[0]].emplace_back(edge[1]);
            g[edge[1]].emplace_back(edge[0]);
        }

        dfs(nums, 0, 1);
        return ans;
    }
};
// @lc code=end

