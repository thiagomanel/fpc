import java.io.FileInputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

public class Sum {

    public static void main(String[] args) throws Exception {
        if (args.length < 1) {
            System.err.println("Usage: java Sum filepath1 filepath2 filepathN");
            System.exit(1);
        }

	    //many exceptions could be thrown here. we don't care
        for (String path : args) {
            Thread thread = new Thread(new FileSum(path), path);
            thread.start();
        }
    }

    public static class FileSum implements Runnable {

        private final String path;
        
        public FileSum(String path) {
            this.path = path;
        }

        @Override
        public void run() {
            try {
                sum();
            } catch (IOException ex) {
            }
        }

        public void sum() throws IOException {
            Path filePath = Paths.get(path);
            if (Files.isRegularFile(filePath)) {
                FileInputStream fis = new FileInputStream(filePath.toString());
                System.out.println(path + ": " + sum(fis));
            } else {
                throw new RuntimeException("Non-regular file: " + path);
            }
        }

        public int sum(FileInputStream fis) throws IOException {
            int byteRead;
            int sum = 0;
            
            while ((byteRead = fis.read()) != -1) {
                sum += byteRead;
            }

            return sum;
        }

    }
}
