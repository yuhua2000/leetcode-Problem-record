/*
 * @lc app=leetcode.cn id=1600 lang=cpp
 *
 * [1600] 王位继承顺序
 */
#include <string>
#include <vector>
#include <map>
#include <set>
#include <iostream>

using namespace std;

// @lc code=start
class ThroneInheritance {
    private:
    map<string, vector<string>> genealogy;
    set<string> deaths;
    string kingName;
    public:
    ThroneInheritance(string kingName) :kingName(kingName)
    {

    }

    void birth(string parentName, string childName)
    {
        genealogy[parentName].emplace_back(childName);
    }

    void death(string name)
    {
        deaths.emplace(name);
    }

    vector<string> getInheritanceOrder()
    {
        vector<string> result;
        getNextGeneration(kingName, result);
        return result;
    }
    void getNextGeneration(string name, vector<string>& inheritance)
    {
        // cout << name << endl;
        if (deaths.find(name) == deaths.end())
        {
            inheritance.emplace_back(name);
        }
        auto childs = genealogy.find(name);
        if (childs == genealogy.end())
        {
            return;
        }
        // cout << childs->second.size() << endl;
        for (auto& child : childs->second)
        {
            getNextGeneration(child, inheritance);
        }
    }
};

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * ThroneInheritance* obj = new ThroneInheritance(kingName);
 * obj->birth(parentName,childName);
 * obj->death(name);
 * vector<string> param_3 = obj->getInheritanceOrder();
 */
 // @lc code=end

