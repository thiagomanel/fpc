import java.util.Timer;
import java.util.TimerTask;

public class TimerTaskOldJava {

    public static void main(String[] args) {

        long delay = 5000;
        long period = 1000;

        MyTimerTask myTimerTask = new MyTimerTask();
        Timer timer = new Timer();

        timer.schedule(myTimerTask, delay, period);
    }


    public static class MyTimerTask extends TimerTask {

        public int i = 0;

        @Override
        public void run() {
            ++i;
            if (i == 5) {
                cancel();
            }
        }
    }
}
