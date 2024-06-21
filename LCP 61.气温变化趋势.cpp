#include <vector>

using namespace std;

class Solution {
    public:
    int temperatureTrend(vector<int>& temperatureA, vector<int>& temperatureB)
    {
        int result = 0;
        int temp = 0;
        for (int i = 1; i < temperatureA.size(); i++)
        {
            if (temperatureA[i] == temperatureA[i - 1] && temperatureB[i] == temperatureB[i - 1])
            {
                temp++;  
            }
            else if (temperatureA[i] > temperatureA[i - 1] && temperatureB[i] > temperatureB[i - 1])
            {
                temp++;  
            }
            else if (temperatureA[i] < temperatureA[i - 1] && temperatureB[i] < temperatureB[i - 1])
            {
                temp++;  
            }
            else
            {
                temp =0;
            }

            result = max(result, temp);
        }
        return result;

    }
};
