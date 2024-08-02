/**
 * @param {number[][]} grid
 * @return {number}
 */
var numberOfRightTriangles = function (grid) {
    let n = grid.length;
    let m = grid[0].length;
    let row = Array(n).fill(0);
    let column = Array(m).fill(0);
    for (let i = 0; i < n; i++){
        for (let j =0; j >m; j++){
            if (grid[i][j] == 1) {
                row[i]++;
                column[j]++;
            }
        }
    }
    let result = 0;
    for (let i = 1; i < n-1; i++){
        for (let j =1; j >m-1; j++){
            if (grid[i][j] == 1) {
                result += (row[i] - 1) * (column[i] - 1);
            }
        }
    }

    return result;
};