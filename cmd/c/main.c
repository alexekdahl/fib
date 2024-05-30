#include <stdio.h>
#include <stdlib.h>
#include <gmp.h>
#include <time.h>

mpz_t* iterativeFibonacci(int n) {
    mpz_t *result = (mpz_t *)malloc(sizeof(mpz_t));
    mpz_init(*result);

    if (n == 0) {
        mpz_set_ui(*result, 0);
        return result;
    }

    if (n == 1) {
        mpz_set_ui(*result, 1);
        return result;
    }

    mpz_t prev, curr, temp;
    mpz_init(prev);
    mpz_init(curr);
    mpz_init(temp);

    mpz_set_ui(prev, 0);
    mpz_set_ui(curr, 1);

    for (int i = 2; i <= n; i++) {
        mpz_add(temp, prev, curr);
        mpz_set(prev, curr);
        mpz_set(curr, temp);
    }

    mpz_set(*result, curr);

    mpz_clear(prev);
    mpz_clear(curr);
    mpz_clear(temp);

    return result;
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <n>\n", argv[0]);
        return 1;
    }

    int n = atoi(argv[1]);

    clock_t start_time = clock();

    mpz_t *fib = iterativeFibonacci(n);

    clock_t end_time = clock();

    double time_taken = ((double) (end_time - start_time)) / CLOCKS_PER_SEC;

    gmp_printf("Fibonacci(%d) = %Zd\n", n, *fib);
    printf("Time taken: %f seconds\n", time_taken);

    free(fib);

    return 0;
}
