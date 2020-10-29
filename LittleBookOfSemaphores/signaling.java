
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

class SendingThread extends Thread {
  Semaphore semaphore = null;

  public SendingThread(Semaphore semaphore){
    this.semaphore = semaphore;
  }

  public void run(){
    while(true){
      System.out.print("Sending...");
      this.semaphore.acquire();
      return;
    }
  }
}

class ReceivingThread extends Thread {
  Semaphore semaphore = null;

  public ReceivingThread(Semaphore semaphore){
    this.semaphore = semaphore;
  }
  
  public void run() {
    while(true) {
    try {
      this.semaphore.release();
      System.out.println(" Received!!");  
    } catch (InterruptedException e) {
        e.printStackTrace();
    }  
      return;
    }
  }
}

class Main {
    public static void main(String[] args)  throws InterruptedException {
        
        for (int i = 0; i < 1000; i++) {

            Semaphore semaphore = new Semaphore();
            SendingThread sender = new SendingThread(semaphore);
            ReceivingThread receiver = new ReceivingThread(semaphore);

            receiver.start();
            sender.start();

            receiver.join();
            sender.join();
        }
    }
}
