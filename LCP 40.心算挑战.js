/**
 * @param {number[]} cards
 * @param {number} cnt
 * @return {number}
 */
var maxmiumScore = function (cards, cnt) {
    cards.sort((a, b) => {
        return b - a;
    });

    let uneven = Array();
    let even = Array();
    for (let i of cards) {
        if (i % 2 == 0) {
            even.push(i);
        } else {
            uneven.push(i);
        }
    }

    let result = 0;
    if (cnt % 2 != 0) {
        if (even.length == 0) {
            return result;
        }
        result += even[0];
        even.shift();
    }

    for (let i = 1; i <= (cnt / 2); i++) {
        let num1 = even.length >= 2 ? even[0] + even[1] : 0;
        let num2 = uneven.length >= 2 ? uneven[0] + uneven[1] : 0;
        if (num1+num2==0){
            return 0;
        }
        if (num1 > num2) {
            result += num1;
            even.splice(0, 2);
        } else {
            result += num2;
            uneven.splice(0, 2);
        }
    }

    return result;
};