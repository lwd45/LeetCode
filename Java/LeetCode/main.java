import java.util.Collections;
import java.util.HashMap;

/**
 * Created by lwd at 2021/8/3
 *
 * @Description:
 */
public class main {
    public static void main(String[] args) {
        MyThread t1 = new MyThread();
        new Thread(t1,"t1").start();
        new Thread(t1,"t2").start();
    }
}

class MyThread implements Runnable {
    @Override
    public void run() {
        synchronized (this) {
            for (int i = 0; i < 10; i++) {
                System.out.println(Thread.currentThread().getName() + ":" + i);
            }
        }
    }
}
