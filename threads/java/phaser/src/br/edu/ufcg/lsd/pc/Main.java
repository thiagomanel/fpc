package br.edu.ufcg.lsd.pc;

import java.util.concurrent.Phaser;

public class Main {

    public static void main(String[] args) {

        //This phaser has two rounds. In the seconds, we wait three workers to finish their job.
        //When the workers of the first round finish, we start the second round with two new workers.

        Phaser ph = new Phaser(1);

        System.err.println("create threads for phaser " + ph);

        new Thread(new Worker("w0", ph)).start();
        new Thread(new Worker("w1", ph)).start();
        new Thread(new Worker("w2", ph)).start();

        System.err.println("wait for the workers for phaser " + ph);

        ph.arriveAndAwaitAdvance();

        System.err.println("start the second round on phaser "  + ph);

        new Thread(new Worker("w3", ph)).start();
        new Thread(new Worker("w4", ph)).start();

        System.err.println("wait for the workers for phaser " + ph);

        ph.arriveAndAwaitAdvance();
        ph.arriveAndDeregister();
    }

    private static class Worker implements Runnable {

        private final String workerId;
        private final Phaser phaser;

        Worker(String workerId, Phaser phaser) {
            this.workerId = workerId;
            this.phaser = phaser;
            System.err.println("worker " + workerId + " is going to register on phaser " +  phaser);
            phaser.register();
        }

        @Override
        public void run() {
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            System.err.println("worker " + workerId + " is going to arrive and deregister on phaser " +  phaser);
            phaser.arriveAndDeregister();
        }
    }

}

