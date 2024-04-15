/*
 * @lc app=leetcode.cn id=705 lang=cpp
 *
 * [705] 设计哈希集合
 */
#include <malloc.h>
#include <string.h>

// @lc code=start
class MyHashSet {
    private:
    bool* nums;
    public:
    MyHashSet()
    {
        nums = (bool*)malloc((1000000 + 1) * sizeof(bool));
        memset(nums, 0, (1000000 + 1) * sizeof(bool));
    }

    void add(int key)
    {
        nums[key] = true;
    }

    void remove(int key)
    {
        nums[key] = false;
    }

    bool contains(int key)
    {
        return nums[key];
    }
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet* obj = new MyHashSet();
 * obj->add(key);
 * obj->remove(key);
 * bool param_3 = obj->contains(key);
 */
 // @lc code=end

