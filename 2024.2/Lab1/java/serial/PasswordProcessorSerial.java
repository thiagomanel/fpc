import java.io.*;
import java.nio.file.*;
import java.util.*;

public class PasswordProcessorSerial {
    public static void main(String[] args) {
        if (args.length < 1) {
            System.out.println("Uso: java PasswordProcessorSerial <caminho_do_diretorio>");
            return;
        }

        String directoryPath = args[0]; // Recebe o caminho como argumento

        File directory = new File(directoryPath);
        if (!directory.exists() || !directory.isDirectory()) {
            System.out.println("Erro: Diretório não encontrado ou inválido.");
            return;
        }

        File[] files = directory.listFiles((dir, name) -> name.endsWith(".txt"));
        if (files == null) {
            System.out.println("Erro ao listar arquivos no diretório.");
            return;
        }

        for (File file : files) {
            processFile(file);
        }
    }

    private static void processFile(File file) {
        System.out.println("Processing file: " + file.getName());
        List<String> obfuscatedLines = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(new FileReader(file))) {
            String line;
            while ((line = reader.readLine()) != null) {
                obfuscatedLines.add(rot13(line)); // Adiciona a linha ofuscada à lista
            }
        } catch (IOException e) {
            System.out.println("Erro ao ler o arquivo " + file.getName() + ": " + e.getMessage());
            return;
        }

        try (BufferedWriter writer = new BufferedWriter(new FileWriter(file))) {
            for (String obfuscatedLine : obfuscatedLines) {
                writer.write(obfuscatedLine);
                writer.newLine();
            }
        } catch (IOException e) {
            System.out.println("Erro ao escrever no arquivo " + file.getName() + ": " + e.getMessage());
        }
    }

    private static String rot13(String input) {
        StringBuilder result = new StringBuilder();
        for (char c : input.toCharArray()) {
            if (c >= 'a' && c <= 'z') {
                result.append((char) (((c - 'a' + 13) % 26) + 'a'));
            } else if (c >= 'A' && c <= 'Z') {
                result.append((char) (((c - 'A' + 13) % 26) + 'A'));
            } else {
                result.append(c);
            }
        }
        return result.toString();
    }
}

