/*
 * @lc app=leetcode.cn id=1146 lang=cpp
 *
 * [1146] 快照数组
 */
#include <vector>
#include <iostream>

using namespace std;

// @lc code=start
class SnapshotArray {
    private:
    vector<vector<pair<int, int>>> SnapsArray;
    int Snap;

    public:
    SnapshotArray(int length)
    {
        Snap = 0;
        SnapsArray.resize(length);
    }

    void set(int index, int val)
    {
        if (SnapsArray[index].size() == 0)
        {
            SnapsArray[index].emplace_back(Snap, val);
            return;
        }
        if (SnapsArray[index].back().second == val)
        {
            return;
        }

        if (SnapsArray[index].back().first == Snap)
        {
            SnapsArray[index].back().second = val;
        }
        else
        {
            SnapsArray[index].emplace_back(Snap, val);
        }
    }

    int snap()
    {
        return Snap++;
    }

    int get(int index, int snap_id)
    {
        int l = 0, r = SnapsArray[index].size() - 1;
        if (r == -1)
        {
            return 0;
        }
        for (auto SpanArray : SnapsArray[index])
        {
            cout << SpanArray.first <<"--->" <<SpanArray.second << "\n";
        }

        while (l <= r)
        {
            int mid = (l + r) >> 1;
            if (SnapsArray[index][mid].first > snap_id)
            {
                r = mid - 1;
            }
            else 
            {
                l = mid + 1;
            }
        }
        if (l - 1 >= 0 && SnapsArray[index][l - 1].first <= snap_id)
        {
            return SnapsArray[index][l - 1].second;
        }
        if (SnapsArray[index][l].first <= snap_id)
        {
            return SnapsArray[index][l].second;
        }

        return 0;
    }
};

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * SnapshotArray* obj = new SnapshotArray(length);
 * obj->set(index,val);
 * int param_2 = obj->snap();
 * int param_3 = obj->get(index,snap_id);
 */
 // @lc code=end

