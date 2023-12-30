#include <vector>

using namespace std;

class Solution {
    public:
    int source(vector<int>& player)
    {
        int source = 0;
        int index = -3;
        for (int i = 0; i < player.size(); i++)
        {
            source += player[i];
            if (i - index <= 2)
            {
                source += player[i];
            }
            if (player[i] == 10)
            {
                index = i;
            }
        }
        return source;
    }
    int isWinner(vector<int>& player1, vector<int>& player2)
    {
        int s1 = source(player1), s2 = source(player2);
        if (s1 == s2)
        {
            return 0;
        }
        return s1 > s2 ? 1 : 2;
    }
};