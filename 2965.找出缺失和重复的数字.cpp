#include <vector>
#include <unordered_set>

using namespace std;

class Solution {
    public:
    vector<int> findMissingAndRepeatedValues(vector<vector<int>>& grid)
    {
        int n = grid.size();
        int totalCount = n * n;  // n x n grid will have numbers from 1 to n*n
        unordered_set<int> seenNumbers;
        int repeatedNumber = 0;
        int xorAll = 0;

        // XOR all numbers from 1 to totalCount
        for (int i = 1; i <= totalCount; i++)
        {
            xorAll ^= i;
        }

        // Traverse the grid and XOR all numbers in the grid
        for (size_t i = 0; i < n; i++)
        {
            for (size_t j = 0; j < n; j++)
            {
                int num = grid[i][j];
                // Check if the number is repeated
                if (repeatedNumber == 0)
                {
                    if (seenNumbers.count(num) == 1)
                        repeatedNumber = num;

                    seenNumbers.insert(num);
                }
                xorAll ^= num;
            }
        }

        // The missing number can be found by XORing xorAll with repeatedNumber
        int missingNumber = xorAll ^ repeatedNumber;

        return { repeatedNumber, missingNumber };
    }
};
