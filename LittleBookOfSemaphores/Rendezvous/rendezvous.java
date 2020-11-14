
class Semaphore {
  private boolean signal = false;

  public synchronized void acquire() {
    this.signal = true;
    this.notify();
  }

  public synchronized void release() throws InterruptedException {
    while(!this.signal) wait();
    this.signal = false;
  }
}

class Thread1 extends Thread {
  Semaphore semaphore1 = null;
  Semaphore semaphore2 = null;

  public Thread1(Semaphore semaphore1, Semaphore semaphore2){
    this.semaphore1 = semaphore1;
    this.semaphore2 = semaphore2;
  }

  public void run(){
    while(true){
      try {
        System.out.print("Thread1 signaling...");
        this.semaphore1.acquire();
        this.semaphore2.release();
        System.out.println("Thread1 receiving!!");
      }catch (InterruptedException e) {
        e.printStackTrace();
      }  
    }
  }
}

class Thread2 extends Thread {
  Semaphore semaphore1 = null;
  Semaphore semaphore2 = null;

  public Thread2(Semaphore semaphore1, Semaphore semaphore2){
    this.semaphore1 = semaphore1;
    this.semaphore2 = semaphore2;
  }

  public void run() {
    while(true) {
    try {
      this.semaphore1.release();
      System.out.println(" Thread2 receiving!!");
      System.out.print("Thread2 signaling...");
      this.semaphore2.acquire();
    } catch (InterruptedException e) {
        e.printStackTrace();
    }  
    }
  }
}

class Main {
  public static void main(String[] args)  throws InterruptedException {
      

        Semaphore semaphore1 = new Semaphore();
        Semaphore semaphore2 = new Semaphore();
        Thread1 sender = new Thread1(semaphore1, semaphore2);
        Thread2 receiver = new Thread2(semaphore1, semaphore2);

        receiver.start();
        sender.start();

        receiver.join();
        sender.join();
  }
}
