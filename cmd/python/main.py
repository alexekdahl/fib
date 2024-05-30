import sys
import time


def fibonacci_big(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    prev, curr = 0, 1
    for _ in range(2, n + 1):
        prev, curr = curr, prev + curr
    return curr


def main():
    num = int(sys.argv[1])
    time_start = time.time()
    result = fibonacci_big(num)
    time_end = time.time()
    duration = time_end - time_start

    print(f"fibonacci({num}) is: {result}")
    print(f"Time taken to calculate: {duration} milliseconds")


if __name__ == "__main__":
    main()
