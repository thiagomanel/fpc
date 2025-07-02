import java.io.*;
import java.util.concurrent.*;

public class ContadorPalavras {
    public static void main(String[] args) {
	System.out.println("Lab8");    

    }

    static int contarPalavras(String nomeArquivo) throws IOException {
        BufferedReader br = new BufferedReader(new FileReader(nomeArquivo));
        int count = 0;
        String linha;
        while ((linha = br.readLine()) != null) {
            count += linha.split("\\s+").length;
        }
        br.close();
        return count;
    }
}

