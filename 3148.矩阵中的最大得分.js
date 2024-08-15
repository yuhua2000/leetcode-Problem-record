/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxScore = function (grid) {
    const minScore = -1e5 - 1
    let n = grid.length
    let m = grid[0].length
    let result = minScore
    let columnMinArray = Array(m).fill(minScore)

    for (let i = 0; i < n; i++) {
        let rowMin = minScore
        for (let j = 0; j < m; j++) {
            let num = grid[i][j]
            let maxScore = minScore
            if (rowMin != minScore) {
                maxScore = num - rowMin
            }

            let columnMin = columnMinArray[j]
            if (columnMin != minScore) {
                maxScore = Math.max(maxScore, num - columnMin)
            }

            result = Math.max(result, maxScore);

            if (maxScore > 0) {
                num -= maxScore
            }

            if (rowMin == minScore || rowMin > num) {
                rowMin = num
            }

            if (columnMin == minScore || columnMin > num) {
                columnMinArray[j] = num
            }
        }
    }

    return result
};
