class Consumer {
    private final Buffer buffer;
    private final int sleepTime;
    private final int id;
    
    public Consumer(int id, Buffer buffer, int sleepTime) {
        this.id = id;
        this.buffer = buffer;
        this.sleepTime = sleepTime;
    }
    
    public void process() {
        while (true) {
            int item = buffer.remove();
            if (item == -1) break;
            System.out.println("Consumer " + id + " consumed item " + item);
            try {
                Thread.sleep(sleepTime);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
    }
}