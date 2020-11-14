import java.util.concurrent.Semaphore;

class Thread1 extends Thread {
    Semaphore mutex = null;

    public Thread1(Semaphore semaphore){
        this.mutex = semaphore;
    }

    public void run(){
        while(true){
            try {
                mutex.acquire();
                Main.cont++;
            }catch (InterruptedException e) {
                e.printStackTrace();
            }finally {
                mutex.release();
            }
            return;
        }
    }
}


class Main {
    static int cont = 0;
    public static void main(String[] args)  throws InterruptedException {
        Semaphore semaphore = new Semaphore(1);
        Thread1 thread1 = new Thread1(semaphore);
        Thread1 thread2 = new Thread1(semaphore);

        thread1.start();
        thread2.start();

        thread1.join();
        thread2.join();

        System.out.println(cont);
    }
}
  