import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class LogAnalyzer {

    static int total200 = 0;
    static int total500 = 0;

    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("Uso: java LogAnalyzer <arquivos_de_log>");
            System.exit(1);
        }

        //int total200 = 0;
        //int total500 = 0;

        for (String fileName : args) {
            System.out.println("Processando arquivo: " + fileName);
	    processFile(fileName);
        }

        System.out.println("===== RESULTADO FINAL =====");
        System.out.println("Total 200: " + total200);
        System.out.println("Total 500: " + total500);
    }

    public static void processFile(String fileName) {
	try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
            String line;
            while ((line = br.readLine()) != null) {
                String[] parts = line.split(" ");
                if (parts.length == 3) {
                    String code = parts[2];
                    if (code.equals("200")) {
                        total200++;
                    } else if (code.equals("500")) {
                        total500++;
                    }
                }
            }
	} catch (IOException e) {
            System.err.println("Erro ao ler arquivo: " + fileName);
            e.printStackTrace();
        }
    }
}
