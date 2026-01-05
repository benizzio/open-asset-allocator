// Test the updated mathHelper implementation
const BigNumber = require('bignumber.js');

// Simulate the updated mathHelper implementation with nullish coalescing
function updatedMathHelper(a, b) {
    let leftBN = new BigNumber(a ?? 0);
    if (leftBN.isNaN()) {
        leftBN = new BigNumber(0);
    }

    let rightBN = new BigNumber(b ?? 0);
    if (rightBN.isNaN()) {
        rightBN = new BigNumber(0);
    }
    
    let resultBN = leftBN.multipliedBy(rightBN);
    return resultBN.toNumber();
}

// Test case from the issue
const testValue = 0.00111;
const multiplier = 100;

console.log('Test case: 0.00111 * 100');
console.log('Expected: 0.111');
console.log('');

const result = updatedMathHelper(testValue, multiplier);
console.log(`Updated implementation result: ${result}`);
console.log(`Updated implementation correct: ${result === 0.111}`);
console.log('');

// Test with null/undefined
console.log('Testing null/undefined handling:');
console.log(`  null * 100 = ${updatedMathHelper(null, 100)} (expected: 0)`);
console.log(`  undefined * 100 = ${updatedMathHelper(undefined, 100)} (expected: 0)`);
console.log(`  100 * null = ${updatedMathHelper(100, null)} (expected: 0)`);
console.log(`  100 * undefined = ${updatedMathHelper(100, undefined)} (expected: 0)`);
