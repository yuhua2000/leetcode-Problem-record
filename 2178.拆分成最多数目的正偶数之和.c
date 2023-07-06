/*
 * @lc app=leetcode.cn id=2178 lang=c
 *
 * [2178] 拆分成最多数目的正偶数之和
 */

// @lc code=start

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
long long* maximumEvenSplit(long long finalSum, int *returnSize){
    long long* rest = NULL;
    *returnSize = 0;
    if (finalSum % 2 >0){
        return rest;
    }
    int k = sqrt(finalSum) + 1;
    rest=(long long*)malloc(sizeof(long long)*k);
    for (int i = 2; i <= finalSum; i+=2)
    {
        rest[++(*returnSize)-1]=i;
        finalSum-=i;   
    }
     rest[(*returnSize)-1]+=finalSum;
     return rest;
}



// @lc code=end
