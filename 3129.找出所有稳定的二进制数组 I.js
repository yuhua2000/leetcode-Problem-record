const MOD = 1000000007;

/**
 * @param {number} zero
 * @param {number} one
 * @param {number} limit
 * @return {number}
 */
var numberOfStableArrays = function (zero, one, limit) {
    const dp = Array.from({ length: zero + 1 }, () =>
        Array.from({ length: one + 1 }, () => [0, 0])
    );

    for (let i = 0; i <= Math.min(limit, zero); i++) {
        dp[i][0][0] = 1;
    }

    for (let i = 0; i <= Math.min(limit, one); i++) {
        dp[0][i][1] = 1;
    }

    for (let i = 1; i <= zero; i++) {
        for (let j = 1; j <= one; j++) {
            if (i > limit) {
                dp[i][j][0] = dp[i - 1][j][0] + dp[i - 1][j][1] - dp[i - limit - 1][j][1];
            } else {
                dp[i][j][0] = dp[i - 1][j][0] + dp[i - 1][j][1];
            }
            dp[i][j][0] = (dp[i][j][0] % MOD + MOD) % MOD;

            if (j > limit) {
                dp[i][j][1] = dp[i][j - 1][1] + dp[i][j - 1][0] - dp[i][j - limit - 1][0];
            } else {
                dp[i][j][1] = dp[i][j - 1][1] + dp[i][j - 1][0];
            }
            dp[i][j][1] = (dp[i][j][1] % MOD + MOD) % MOD;
        }
    }

    return (dp[zero][one][0] + dp[zero][one][1]) % MOD;
};

