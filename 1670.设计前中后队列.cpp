/*
 * @lc app=leetcode.cn id=1670 lang=cpp
 *
 * [1670] 设计前中后队列
 */
#include <queue>

using namespace std;

// @lc code=start
class FrontMiddleBackQueue {
    public:
    FrontMiddleBackQueue()
    {

    }

    void pushFront(int val)
    {
        left.push_front(val);
        if (left.size() == right.size() + 2)
        {
            right.push_front(left.back());
            left.pop_back();
        }
    }

    void pushMiddle(int val)
    {
        if (left.size() == right.size() + 1)
        {
            right.push_front(left.back());
            left.pop_back();
        }
        left.push_back(val);
    }

    void pushBack(int val)
    {
        right.push_back(val);
        if (right.size() == left.size() + 1)
        {
            left.push_back(right.front());
            right.pop_front();
        }
    }

    int popFront()
    {
        if (left.empty())
        {
            return -1;
        }

        int result = left.front();
        left.pop_front();
        if (right.size() == left.size() + 1)
        {
            left.push_back(right.front());
            right.pop_front();
        }
        return result;
    }

    int popMiddle()
    {
        if (left.empty())
        {
            return -1;
        }
        int result = left.back();
        left.pop_back();
        if (right.size() == left.size() + 1)
        {
            left.push_back(right.front());
            right.pop_front();
        }
        return result;

    }

    int popBack()
    {
        if (left.empty())
        {
            return -1;
        }
        int result;
        if (right.empty())
        {
            result = left.back();
            left.pop_back();
        }
        else
        {
            result = right.back();
            right.pop_back();
            if (left.size() == right.size() + 2)
            {
                right.push_front(left.back());
                left.pop_back();
            }
        }
        return result;

    }
    private:
    deque<int> left;
    deque<int> right;
};

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * FrontMiddleBackQueue* obj = new FrontMiddleBackQueue();
 * obj->pushFront(val);
 * obj->pushMiddle(val);
 * obj->pushBack(val);
 * int param_4 = obj->popFront();
 * int param_5 = obj->popMiddle();
 * int param_6 = obj->popBack();
 */
 // @lc code=end

