/*
 * @lc app=leetcode.cn id=894 lang=cpp
 *
 * [894] 所有可能的真二叉树
 */

#include <vector>

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
    public:
    vector<TreeNode*> allPossibleFBT(int n)
    {
        vector<TreeNode*> fullBinaryTrees;
        if (n % 2 == 0)
            return fullBinaryTrees;

        if (n == 1)
        {
            fullBinaryTrees = { new TreeNode(0) };
            return fullBinaryTrees;
        }

        for (int i = 1; i < n; i += 2)
        {
            vector<TreeNode*> leftSubtrees = allPossibleFBT(i);
            vector<TreeNode*> rightSubtrees = allPossibleFBT(n - i - 1);
            for (TreeNode* leftSubtree : leftSubtrees)
            {
                for (TreeNode* rightSubtree : rightSubtrees)
                {
                    TreeNode* node = new TreeNode(0, leftSubtree, rightSubtree);
                    fullBinaryTrees.emplace_back(node);
                }
            }
        }
        return fullBinaryTrees;
    }
};
// @lc code=end

