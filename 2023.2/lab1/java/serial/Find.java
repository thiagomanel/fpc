public class Find {
    public static void main(String[] args) {
        if (args.length != 1) {
            System.out.println("Use: java Main <size>");
            System.exit(1);
        }

        try {
            int size = Integer.parseInt(args[0]);

            int[][] matrix = Matrix.generateMatrix(size);

            System.out.printf("Max value: %d\n", Matrix.Max(matrix));
            System.out.printf("Min value: %d\n", Matrix.Min(matrix));
            
        } catch (NumberFormatException e) {
            System.out.println("The size of the matrix must be an integer.");
            System.exit(1);
        }

        
    }
}
