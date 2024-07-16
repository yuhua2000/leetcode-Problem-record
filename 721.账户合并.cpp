/*
 * @lc app=leetcode.cn id=721 lang=cpp
 *
 * [721] 账户合并
 */
#include <vector>
#include <string>
#include <map>
#include <algorithm>
#include <iostream>

using namespace std;

// @lc code=start

class UnionFind {
    private:
    vector<int> parent;

    public:
    UnionFind(int n)
    {
        parent.resize(n);
        for (int i = 0; i < n; i++)
        {
            parent[i] = i;
        }
    }


    int find(int x)
    {
        if (parent[x] != x)
        {
            parent[x] = find(parent[x]);
        }

        return parent[x];
    }

    void unionSe(int to, int from)
    {
        parent[find(from)] = find(to);
    }
};

class Solution {
    public:
    vector<vector<string>> accountsMerge(vector<vector<string>>& accounts)
    {
        map<string, int> emailToIndex;
        map<string, string> emailToName;
        int emailsCount = 0;

        for (auto& account : accounts)
        {
            string& name = account[0];
            for (int i = 1; i < account.size(); i++)
            {
                string& email = account[i];
                if (!emailToIndex.count(email))
                {
                    emailToIndex[email] = emailsCount++;
                    emailToName[email] = name;
                }
            }
        }

        UnionFind uf(emailsCount);
        for (auto& account : accounts)
        {
            int firstIndex = emailToIndex[account[1]];
            for (int i = 2; i < account.size(); i++)
                uf.unionSe(firstIndex, emailToIndex[account[i]]);
        }

        map<int, vector<string>> indexToEmails;
        for (auto& [email, _] : emailToIndex)
        {
            int index = uf.find(emailToIndex[email]);
            indexToEmails[index].emplace_back(email);
        }

        vector<vector<string>> merged;
        for (auto& [_, emails] : indexToEmails)
        {
            sort(emails.begin(), emails.end());
            string& name = emailToName[emails[0]];
            vector<string> account;
            account.emplace_back(name);
            for (auto& email : emails)
            {
                account.emplace_back(email);
            }
            merged.emplace_back(account);

        }

        return merged;
    }
};
// @lc code=end

 