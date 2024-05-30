const { performance } = require("perf_hooks");

function fibonacciBig(n) {
  if (n === 0) {
    return BigInt(0);
  }
  if (n === 1) {
    return BigInt(1);
  }
  let prev = BigInt(0);
  let curr = BigInt(1);
  for (let i = 2; i <= n; i++) {
    [prev, curr] = [curr, prev + curr];
  }
  return curr;
}

function recursiveFibonacci(n) {
  if (n === 0) {
    return BigInt(0);
  }
  if (n === 1) {
    return BigInt(1);
  }
  return recursiveFibonacci(n - 1) + recursiveFibonacci(n - 2);
}

function recursiveFibonacciWithCache(n, cache) {
  if (n === 0) {
    return BigInt(0);
  }
  if (n === 1) {
    return BigInt(1);
  }
  if (cache[n] === undefined) {
    cache[n] =
      recursiveFibonacciWithCache(n - 1, cache) +
      recursiveFibonacciWithCache(n - 2, cache);
  }
  return cache[n];
}

function main() {
  const num = parseInt(process.argv[2], 10);

  const timeStart = performance.now();
  const result = fibonacciBig(num);
  const timeEnd = performance.now();
  const duration = timeEnd - timeStart;

  const timeStart2 = performance.now();
  const result2 = recursiveFibonacci(num);
  const timeEnd2 = performance.now();
  const duration2 = timeEnd2 - timeStart2;

  const timeStart3 = performance.now();
  const result3 = recursiveFibonacciWithCache(num, {});
  const timeEnd3 = performance.now();
  const duration3 = timeEnd3 - timeStart3;

  console.log("Recursive Fibonacci without cache:");
  console.log(`fibonacci(${num}) is: ${result2.toString()}`);
  console.log(`Time taken to calculate: ${duration2} milliseconds`);
  console.log();

  console.log("Recursive Fibonacci with cache:");
  console.log(`fibonacci(${num}) is: ${result3.toString()}`);
  console.log(`Time taken to calculate: ${duration3} milliseconds`);
  console.log();

  console.log("Iterative Fibonacci:");
  console.log(`fibonacci(${num}) is: ${result.toString()}`);
  console.log(`Time taken to calculate: ${duration} milliseconds`);
}

main();
