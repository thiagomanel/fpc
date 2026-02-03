public class Main {
    public static void main(String[] args) {
        if (args.length != 5) {
            System.out.println("Use: java Main <num_producers> <max_items_per_producer> <producing_time> <num_consumers> <consuming_time>");
            return;
        }
        
        int numProducers = Integer.parseInt(args[0]);
        int maxItemsPerProducer = Integer.parseInt(args[1]);
        int producingTime = Integer.parseInt(args[2]);
        int numConsumers = Integer.parseInt(args[3]);
        int consumingTime = Integer.parseInt(args[4]);
        
        Buffer buffer = new Buffer();
        
        for (int i = 1; i <= numProducers; i++) {
            Producer producer = new Producer(i, buffer, maxItemsPerProducer, producingTime);
            producer.produce();
        }
        
        for (int i = 1; i <= numConsumers; i++) {
            Consumer consumer = new Consumer(i, buffer, consumingTime);
            consumer.process();
        }
    }
}
