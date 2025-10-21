import java.lang.Thread;
import java.lang.Runnable;
import java.util.Random; 

public class SimpleSerialSolution {

    // Lógica da Tarefa de Verificação de Recursos (Execução Sequencial)
    private static void executeResourceCheck() {
        // Nome da Thread principal para o log
        String threadName = Thread.currentThread().getName(); 
        System.out.println("[" + threadName + "] INÍCIO: Verificação de Recursos.");

        Random random = new Random();

        for (int i = 1; i <= 5; i++) {
            try {
                // Gera um valor aleatório entre 1000ms (1s) e 3000ms (3s)
                int sleepTime = 1000 + random.nextInt(2000); 
                System.out.println("[" + threadName + "] Verificando Recurso " + i + " (Duração: " + sleepTime + "ms)...");
                Thread.sleep(sleepTime);
            } catch (InterruptedException e) {
                System.err.println("[" + threadName + "] Verificação interrompida.");
                Thread.currentThread().interrupt();
                return;
            }
        }
        System.out.println("[" + threadName + "] FIM: Verificação concluída.");
    }

    // Lógica da Tarefa de Inicialização de Logs (Execução Sequencial)
    private static void executeLogSetup() {
        String threadName = Thread.currentThread().getName();
        System.out.println("[" + threadName + "] INÍCIO: Configuração de Logs.");
        
        try {
            Thread.sleep(4000); // TEMPO FIXO: 4.0 segundos
        } catch (InterruptedException e) {
            System.err.println("[" + threadName + "] Configuração de Logs interrompida.");
            Thread.currentThread().interrupt();
        }
        
        System.out.println("[" + threadName + "] FIM: Configuração de Logs concluída.");
    }

    public static void main(String[] args) {
        Thread currentThread = Thread.currentThread();
        System.out.println("[" + currentThread.getName() + "] --- INÍCIO DO PROGRAMA SERIAL JAVA ---");

        executeLogSetup();
        executeResourceCheck();


        System.out.println("[" + currentThread.getName() + "] --- FIM DO PROGRAMA SERIAL JAVA ---");
    }
}
