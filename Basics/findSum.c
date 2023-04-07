// #include <stdio.h>
// #include <time.h>

// unsigned long get_current_time_ms()
// {
//     struct timespec spec;
//     clock_gettime(CLOCK_REALTIME, &spec);
//     return spec.tv_sec * 1000 + spec.tv_nsec / 1000000;
// }

// int main()
// {
//     int n, sum = 0;
//     printf("Enter a positive integer: ");
//     scanf("%d", &n);

//     unsigned long start_time = get_current_time_ms();

//     for (int i = 1; i <= n; i++)
//     {
//         sum += i;
//     }

//     unsigned long end_time = get_current_time_ms();
//     printf("Sum = %d\n", sum);
//     printf("Time taken = %lu ms\n", end_time - start_time);

//     return 0;
// }
#include <stdio.h>
#include <time.h>

int main()
{
    int n, sum = 0;
    printf("Enter a positive integer: ");
    scanf("%d", &n);

    clock_t start_time = clock();

    // for (int i = 1; i <= n; i++)
    // {
    //     sum += i;
    // }
    sum = n * (n + 1) / 2;

    clock_t end_time = clock();
    printf("Sum = %d\n", sum);
    printf("Time taken = %f ms\n", ((double)(end_time - start_time)) / CLOCKS_PER_SEC * 1000);

    return 0;
}
