import random
import math
import numpy as np

SEED = 0

def random_number_generator(i):

    multiplier = 24693
    increment = 3517
    modulus = 2**17
    seed = 1000

    res = seed
    for j in range(1, i + 1):
        res = (multiplier * res + increment) % modulus

    return res / modulus


def get_random():
    global SEED
    SEED += 1
    return random_number_generator(SEED)

def random_variable_generator(random_num):
    def inverse_cdf(p):
        t = 57
        a = 1 / t
        return math.sqrt(-2 * math.log(1-p) / a**2)

    return inverse_cdf(random_num)

def sample_cdf(sample, x):
    return np.where(sample <= x)[0].size / sample.size

def print_cdf(sample, x, complementary=False):

    if not complementary:
        print(f"P[W <= {x}] = {sample_cdf(sample, x)}")
    else:
        print(f"P[W > {x}] = {1 - sample_cdf(sample, x)}")

def print_cdf_table(sample, x):
    print(f"{x},{sample_cdf(sample, x)}")


def log_sample_mean(n):
    distances = [random_variable_generator(get_random()) for _ in range(n)]
    # print(f"n={n}")
    # print(np.mean(distances))
    # print(f"{n},{np.mean(distances)}")
    with open('log.csv', 'a') as f:
        f.write(f"{n},{sum(distances) / len(distances)}\n")


def simulate():

    n_vals = [10, 30, 50, 100, 150, 250, 500, 1000]

    for n in n_vals:
        print(f"n={n}")
        for _ in range(110):
            log_sample_mean(n)



def main():

    print("Random numbers:")
    for i in range(1, 4):
        print(f"\tu_{i} = {random_number_generator(i)}")
    print()
    for i in range(51, 54):
        print(f"\tu_{i} = {random_number_generator(i)}")
    print()

    simulate()

    # N = 1000
    # print(f"Running {N} trials.")

    # times = []

    # for trial in range(1, N+1):
    #     random_num = random_number_generator(trial)
    #     X = random_variable_generator(random_num)
    #     time = simulate(X)

    #     times.append(time)

    # times = np.array(times)

    # print(f"Finished running {N} trials.\n")

    # print(f"Mean: {np.mean(times)}")
    # print(f"First quartile: {np.quantile(times, .25)}")
    # print(f"Median: {np.median(times)}")
    # print(f"Third quartile: {np.quantile(times, .75)}\n")

    # print_cdf(times, 15)
    # print_cdf(times, 20)
    # print_cdf(times, 30)
    # print_cdf(times, 60, complementary=True)
    # print_cdf(times, 90, complementary=True)
    # print_cdf(times, 110, complementary=True)
    # print_cdf(times, 125, complementary=True)

    # print_cdf_table(times, 15)
    # print_cdf_table(times, 20)
    # print_cdf_table(times, 30)
    # print_cdf_table(times, 60)
    # print_cdf_table(times, 90)
    # print_cdf_table(times, 110)
    # print_cdf_table(times, 125)




if __name__ == "__main__":
    main()