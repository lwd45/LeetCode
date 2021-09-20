import java.util.*;

public class Main {
    public static void main(String[] args) {
        // minHeap();
        // maxHeap();
        // arrayCmp();
        // linkList();
        // stack();
    }

    private static void stack() {
        Stack<Integer> stack = new Stack<>();
        stack.push(1);
        stack.pop();
        stack.peek();
    }

    private static void linkList() {
        LinkedList<Integer> list = new LinkedList<>();
        list.addFirst(1);
        list.addLast(2);
        list.removeLast();
        list.removeFirst();
    }

    private static void arrayCmp() {
        List<String> list = new ArrayList<>();
        list.sort((i, j) -> {
            if (i.length() == j.length())
                return i.compareTo(j);
            return j.length() - i.length();
        });
    }

    private static void maxHeap() {
        // PriorityQueue<Integer> maxHeap = new PriorityQueue<>((i, j) -> j - i);
        // maxHeap.add(4);
        // maxHeap.add(1);
        // maxHeap.add(5);
        // maxHeap.add(2);
        // maxHeap.add(10);
        // while (maxHeap.size() > 0) {
        //     System.out.println(maxHeap.poll());
        // }

        PriorityQueue<String> maxHeap = new PriorityQueue<>((i, j) -> {
            if (i.length() == j.length()) return i.compareTo(j);
            return j.length() - i.length();
        });
        maxHeap.add("4");
        maxHeap.add("1");
        maxHeap.add("5");
        maxHeap.add("2");
        maxHeap.add("10");
        maxHeap.add("30");
        maxHeap.add("220");
        maxHeap.add("20");
        while (maxHeap.size() > 0) {
            System.out.println(maxHeap.poll());
        }
    }

    private static void minHeap() {
        PriorityQueue<Integer> min = new PriorityQueue<>();
        // min.add(4);
        // min.add(1);
        // min.add(5);
        // min.add(2);
        // min.add(10);
        // while(min.size() > 0){
        //     System.out.println(min.poll());
        // }

        PriorityQueue<String> minStr = new PriorityQueue<>();
        minStr.add("4");
        minStr.add("1");
        minStr.add("5");
        minStr.add("2");
        minStr.add("10");
        while (minStr.size() > 0) {
            System.out.println(minStr.poll());
        }
    }
}