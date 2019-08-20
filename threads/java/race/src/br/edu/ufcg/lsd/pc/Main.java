package br.edu.ufcg.lsd.pc;

public class Main {

    public static void main(String[] args) throws InterruptedException {
        IDMaker idMaker = new IDMaker();

        Service s0 = new Service(idMaker, "s0");
        Service s1 = new Service(idMaker, "s1");

        Thread t0 = new Thread(s0, "thread-s0");
        Thread t1 = new Thread(s1, "thread-s1");

        t0.start();
        t1.start();

        t0.join();
        t1.join();
    }

    public static class Service implements Runnable {

        private final IDMaker idMaker;
        private final String serviceId;

        public Service(IDMaker idMaker, String serviceId) {
            this.idMaker = idMaker;
            this.serviceId = serviceId;
        }

        @Override
        public void run() {
            while(true) {
                System.err.println("ServiceId=" + this.serviceId +
                        " requestId=" + idMaker.getNext());
            }
        }
    }

}
