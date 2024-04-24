/*
 * @lc app=leetcode.cn id=2385 lang=cpp
 *
 * [2385] 感染二叉树需要的总时间
 */
#include <unordered_map>
#include <vector>
#include <queue>
#include <unordered_set>

using namespace std;

// Definition for a binary tree node.
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode* left, TreeNode* right) : val(x), left(left), right(right) {}
};

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
    private:

    public:
    int amountOfTime(TreeNode* root, int start)
    {
        unordered_map<int, vector<int>> graph;

        function<void(TreeNode*)> dfs = [&] (TreeNode* node) {
            for (TreeNode* child : vector<TreeNode*>{ node->left, node->right })
            {
                if (child!=nullptr)
                {
                    graph[node->val].push_back(child->val);
                    graph[child->val].push_back(node->val);
                    dfs(child);
                }
            }
            };
        dfs(root);
        queue<vector<int>> q;
        q.push({ start,0 });
        unordered_set<int> visited;
        visited.insert(start);
        int time = 0;

        while (!q.empty())
        {
            auto arr = q.front();
            q.pop();
            int nodeVal = arr[0];
            time = arr[1];
            for (auto val : graph[nodeVal])
            {
                if (!visited.count(val))
                {
                    q.push({ val,time + 1 });
                    visited.insert(val);
                }
            }
        }
        return time;
    }
};
// @lc code=end

