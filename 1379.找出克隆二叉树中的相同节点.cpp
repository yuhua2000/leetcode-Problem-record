/*
 * @lc app=leetcode.cn id=1379 lang=cpp
 *
 * [1379] 找出克隆二叉树中的相同节点
 */
#include <queue>

using namespace std;

// Definition for a binary tree node.
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};


// @lc code=start
class Solution {
    public:
    TreeNode* getTargetCopy1(TreeNode* original, TreeNode* cloned, TreeNode* target)
    {
        queue<TreeNode*> originalQueue;
        queue<TreeNode*> clonedQueue;
        originalQueue.push(original);
        clonedQueue.push(cloned);
        while (!originalQueue.empty())
        {
            TreeNode* node1 = originalQueue.front(), * node2 = clonedQueue.front();
            originalQueue.pop();
            clonedQueue.pop();
            if (node1 == target)
            {
                return node2;
            }
            if (node1->left)
            {
                originalQueue.push(node1->left);
                clonedQueue.push(node2->left);
            }
            if (node1->right)
            {
                originalQueue.push(node1->right);
                clonedQueue.push(node2->right);
            }
        }
        return NULL;
    }
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target)
    {
        if (original==nullptr)
        {
            return nullptr;
        }
        if (original==target)
        {
            return cloned;
        }
        
        TreeNode* left = getTargetCopy(original->left, cloned->left, target);
        if (left)
        {
            return left;
        }
         TreeNode* right = getTargetCopy(original->right, cloned->right, target);
        if (right)
        {
            return right;
        }
        return NULL;
    }
};
// @lc code=end

