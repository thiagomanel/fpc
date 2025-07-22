import java.io.IOException;
import java.nio.file.*;
import java.util.*;

public class FileIndexingPipeline {

    static Map<String, Map<String, Integer>> fileIndex = new HashMap<>();

    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("Uso: java IndexadorPipeline <arquivo1.txt> <arquivo2.txt> ...");
            return;
        }

        Buffer<FileData> readBuffer = new Buffer<>();
        Buffer<FileData> tokenBuffer = new Buffer<>();

        for (String pathStr : args) {
            // Etapa 1: Leitura do arquivo
            if (readFile(pathStr, readBuffer)) {
                // Etapa 2: Tokenização
                tokenize(readBuffer, tokenBuffer);

                // Etapa 3: Indexação
                index(tokenBuffer);
            }
        }

        System.out.println("fileIndex:");
        for (var word : fileIndex.keySet()) {
            System.out.println(word + " -> " + fileIndex.get(word));
        }
    }

    public static boolean readFile(String pathStr, Buffer<FileData> readBuffer) {
        try {
            Path path = Paths.get(pathStr);
            String content = Files.readString(path);
            readBuffer.insert(new FileData(path.getFileName().toString(), content));
            return true;
        } catch (IOException e) {
            System.err.println("Erro ao ler arquivo " + pathStr + ": " + e.getMessage());
            return false;
        }
    }

    public static void tokenize(Buffer<FileData> readBuffer, Buffer<FileData> tokenBuffer) {
        FileData fileData = readBuffer.remove();
        if (fileData == null) return;
        String[] words = fileData.content.split("\\s+");
        String newContent = String.join(",", words);
        tokenBuffer.insert(new FileData(fileData.name, newContent));
    }

    public static void index(Buffer<FileData> tokenBuffer) {
        FileData fileData = tokenBuffer.remove();
        if (fileData == null) return;
        String[] words = fileData.content.split(",");
        for (String word : words) {
            fileIndex.putIfAbsent(word, new HashMap<>());
            Map<String, Integer> fileDatas = fileIndex.get(word);
            fileDatas.put(fileData.name, fileDatas.getOrDefault(fileData.name, 0) + 1);
        }
    }

    static class FileData {
        public final String name;
        public final String content;

        public FileData(String name, String content) {
            this.name = name;
            this.content = content;
        }
    }

    static class Buffer<T> {
        private final Queue<T> queue = new LinkedList<>();

        public void insert(T item) {
            queue.add(item);
        }

        public T remove() {
            return queue.poll();
        }

        public boolean isEmpty() {
            return queue.isEmpty();
        }
    }
}
