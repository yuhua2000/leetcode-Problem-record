#include <string>
#include <deque>

using namespace std;

class Solution {
    public:
    string finalString(string s)
    {
        deque<char> q;
        bool head = false;
        for (auto c : s)
            if (c == 'i')
                head = !head;
            else
                if (head)
                    q.push_front(c);
                else
                    q.push_back(c);

        string ans;
        if (head)
            ans = string{q.rbegin(), q.rend()};
        else
            ans = string{ q.begin(), q.end() };
        
        return   ans;
    }
};