/*
 * @lc app=leetcode.cn id=706 lang=cpp
 *
 * [706] 设计哈希映射
 */

// @lc code=start
class MyHashMap {
    private:
    int nums[1000000+1] = {};
    public:
    MyHashMap() {
        
    }
    
    void put(int key, int value) {
        nums[key] = value + 1;
    }
    
    int get(int key)
    {
        return nums[key] - 1;
    }
    
    void remove(int key) {
        nums[key] = 0;
    }
};

/**
 * Your MyHashMap object will be instantiated and called as such:
 * MyHashMap* obj = new MyHashMap();
 * obj->put(key,value);
 * int param_2 = obj->get(key);
 * obj->remove(key);
 */
// @lc code=end

