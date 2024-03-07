import random

def generate_matrix(size):
    return [[random.randint(250, 29500) for _ in range(size)] for _ in range(size)]

def min(matrix):
    smallest = float('inf')

    for row in matrix:
        for element in row:
            if element < smallest:
                smallest = element

    return smallest

def max(matrix):
    largest = float('-inf')

    for row in matrix:
        for element in row:
            if element > largest:
                largest = element

    return largest
