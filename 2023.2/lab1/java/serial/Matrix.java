import java.util.Random;

public class Matrix {
    public static int[][] generateMatrix(int size) {
        int[][] matrix = new int[size][size];
        Random random = new Random();

        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                matrix[i][j] = random.nextInt(29500 - 250 + 1) + 250;
            }
        }

        return matrix;
    }

    public static int Min(int[][] matrix) {
        int smallest = Integer.MAX_VALUE;

        for (int[] row : matrix) {
            for (int element : row) {
                if (element < smallest) {
                    smallest = element;
                }
            }
        }

        return smallest;
    }

    public static int Max(int[][] matrix) {
        int largest = Integer.MIN_VALUE;

        for (int[] row : matrix) {
            for (int element : row) {
                if (element > largest) {
                    largest = element;
                }
            }
        }

        return largest;
    }
}
