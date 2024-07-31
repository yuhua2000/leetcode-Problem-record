/**
 * @param {number[][]} points
 * @param {number} w
 * @return {number}
 */
var minRectanglesToCoverPoints = function (points, w) {
    points.sort(function (x, y) {
        return x[0] - y[0];
    });

    let idx = points[0][0];
    let result = 1;
    for (let i = 1; i < points.length; i++){
        if (points[i][0] - idx > w) {
            result++;
            idx = points[i][0];
        }
    }

    return result;
};