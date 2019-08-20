package br.edu.ufcg.lsd.pc;

public class IDMaker {

    private int lastID;

    public synchronized int getNext() {
        return lastID++;
    }
}
