/*
 * @lc app=leetcode.cn id=924 lang=cpp
 *
 * [924] 尽量减少恶意软件的传播
 */

#include <vector>
#include <unordered_map>
#include <queue>

using namespace std;
// @lc code=start
class Solution {
    public:
    int minMalwareSpread(vector<vector<int>>& graph, vector<int>& initial)
    {
        int n = graph.size();
        vector<int> ids(n);
        unordered_map<int, int> id_to_size;
        int id = 0;
        for (size_t i = 0; i < n; i++)
        {
            if (!ids[i])
            {
                id++;
                int size = 1;
                queue<int> q;
                q.push(i);
                ids[i] = id;
                while (!q.empty())
                {
                    int u = q.front();
                    q.pop();
                    for (int v = 0; v < n; v++)
                    {
                        if (!ids[v] && graph[u][v] == 1)
                        {
                            ++size;
                            ids[v] = id;
                            q.emplace(v);
                        }
                    }
                }
                id_to_size[id] = size;
            }
        }
        unordered_map<int, int> id_to_initials;
        for (auto u : initial)
        {
            id_to_initials[ids[u]]++;
        }
        int ans = n + 1, ans_removed = 0;
        for (int u : initial)
        {
            int removed = 0;
            if (id_to_initials[ids[u]] == 1)
            {
                removed = id_to_size[ids[u]];
            }
            if (removed > ans_removed || (removed == ans_removed && u < ans))
            {
                ans = u;
                ans_removed = removed;
            }
        }
        return ans;
    }
};
// @lc code=end

