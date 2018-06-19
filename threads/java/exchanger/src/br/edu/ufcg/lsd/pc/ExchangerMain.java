package br.edu.ufcg.lsd.pc;

import java.util.ArrayList;
import java.util.Random;
import java.util.concurrent.Exchanger;

public class ExchangerMain {

    public static class DataChunk {

        private final int size;
        private final ArrayList<Integer> buffer;

        public DataChunk(int size) {
            this.size = size;
            this.buffer = new ArrayList<Integer>();
        }

        public void put(int i) {
            buffer.add(i);
        }

        public int remove() {
            return buffer.remove(0);
        }

        public boolean isEmpty() {
            return this.buffer.isEmpty();
        }

        public boolean isFull() {
            return this.buffer.size() == this.size;
        }
    }

    public static void main(String args[]) {
        Exchanger<DataChunk> exchanger = new Exchanger<DataChunk>();

        Taker taker = new Taker(exchanger);
        Filler filler = new Filler(exchanger);

        Thread takerThread = new Thread(taker);
        Thread fillerThread = new Thread(filler);

        takerThread.start();
        fillerThread.start();
    }

    private static class Filler implements Runnable {

        private Exchanger<DataChunk> exchanger;

        public Filler(Exchanger<DataChunk> exchanger) {
            this.exchanger = exchanger;
        }

        @Override
        public void run() {

            DataChunk dc = new DataChunk(5);

            while (dc != null) {
                //fill the buffer
                while (!dc.isFull()) {
                    int item = new Random().nextInt();
                    dc.put(item);
                }

                //exchange data
                try {
                    dc = exchanger.exchange(dc);
                } catch (InterruptedException e) {}
            }
        }
    }

    private static class Taker implements Runnable {

        private Exchanger<DataChunk> exchanger;

        public Taker(Exchanger<DataChunk> exchanger) {
            this.exchanger = exchanger;
        }

        @Override
        public void run() {
            DataChunk dc = new DataChunk(5);

            while (dc != null) {
                //empty the buffer
                while (!dc.isEmpty()) {
                    int taken = dc.remove();
                }

                //exchange data

                try {
                    dc = exchanger.exchange(dc);
                } catch (InterruptedException e) {}
            }
        }
    }
}
