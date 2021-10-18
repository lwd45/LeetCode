import java.util.*;
import java.util.concurrent.TimeUnit;
import java.util.stream.IntStream;

public class Main {
    public static void main(String[] args) {
        //新建一个ThreadLocal
        ThreadLocal<String> local = new ThreadLocal<>();
        //新建一个随机数类
        Random random = new Random();
        //使用java8的Stream新建5个线程
        IntStream.range(0, 5).forEach(a-> new Thread(()-> {
            //为每一个线程设置相应的local值
            local.set(a+"  "+random.nextInt(10));
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println("线程和local值分别是  "+ local.get());
        }).start());
    }
}