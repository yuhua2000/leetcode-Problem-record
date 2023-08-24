/*
 * @lc app=leetcode.cn id=1267 lang=c
 *
 * [1267] 统计参与通信的服务器
 */

 // @lc code=start


int countServers(int** grid, int gridSize, int* gridColSize) {
    int n = gridSize, m = gridColSize[0];
    int rows[n], cols[m];
    memset(rows, 0, sizeof(rows));
    memset(cols, 0, sizeof(cols));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            if (grid[i][j] == 1) {
                rows[i]++;
                cols[j]++;
            }
        }
    }

    int ans = 0;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            if (grid[i][j] == 1 && (rows[i] >= 2 || cols[j] >= 2)) {
                ans++;
            }
        }
    }
    return ans;
}
// @lc code=end

