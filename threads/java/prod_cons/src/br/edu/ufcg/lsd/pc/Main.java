package br.edu.ufcg.lsd.pc;

import java.util.Random;

public class Main {

    public static void main(String args[]) throws InterruptedException {


        Data data = new Data();
        Producer producer = new Producer(data);
        Consumer consumer = new Consumer(data);

        Thread t0 = new Thread(producer,"producer-thread");
        Thread t1 = new Thread(consumer,"consumer-thread");

        t0.start();
        t1.start();

        t0.join();;
        t1.join();
    }

    public static class Producer implements Runnable {

        private Data data;

        public Producer(Data data) {
            this.data = data;
        }

        @Override
        public void run() {

            while (true) {
                synchronized (this.data) {
                    while (!this.data.isEmpty()) {
                        try {
                            //the buffer is full. 
                            //wait until something is consumed
                            this.data.wait();
                        } catch (InterruptedException e) { }
                    }
                    int produced = new Random().nextInt(11);
                    this.data.put(produced);
                    System.err.println("value produced: " + produced);
                    this.data.notifyAll();
                }
            }
        }
    }

    public static class Consumer implements Runnable {

        private final Data data;

        public Consumer(Data data) {
            this.data  = data;
        }

        public void run() {

            while (true) {
                synchronized (this.data) {
                    while (this.data.isEmpty()) {
                        try {
                            //the buffer is empty
                            //wait until something be produced
                            this.data.wait();
                        } catch (InterruptedException e) { }
                    }
                    int taken = this.data.take();
                    System.err.println("value consumed: " + taken);
                    this.data.notifyAll();
                }
            }
        }
    }

    public static class Data {

        //a buffer of 1-capacity to ease the code,
        //we hold only positive numbers, so -1
        //indicates the buffer is available
        private int value = -1;

        public void put(int v) {
            if (v < 0) {
                throw new 
                IllegalArgumentException(
                    "Cannot hold negative numbers");
            }
            this.value = v;
        }

        public int take() {
            int valueToTake = value;
            value = -1;
            return valueToTake;
        }

        public boolean isEmpty() {
            return value == -1;
        }
    }

}
