/*
 * @lc app=leetcode.cn id=1261 lang=cpp
 *
 * [1261] 在受污染的二叉树中查找元素
 */
#include <vector>
#include <unordered_set>
using namespace std;

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
class FindElements {
    private:
    unordered_set<int> valSet;
    public:
    FindElements(TreeNode* root)
    {
        if (!root)
        {
            return;
        }
        root->val = 0;
        vector<TreeNode*> level{ root };
        while (!level.empty())
        {
            vector<TreeNode*> temp;
            for (auto& it : level)
            {
                valSet.insert(it->val);
                if (it->left)
                {
                    it->left->val = it->val * 2 + 1;
                    temp.push_back(it->left);
                }
                if (it->right)
                {
                    it->right->val = it->val * 2 + 2;
                    temp.push_back(it->right);
                }
            }
            level = temp;
        }
    }

    bool find(int target)
    {
        return valSet.count(target) > 0;
    }
};

/**
 * Your FindElements object will be instantiated and called as such:
 * FindElements* obj = new FindElements(root);
 * bool param_1 = obj->find(target);
 */
 // @lc code=end

