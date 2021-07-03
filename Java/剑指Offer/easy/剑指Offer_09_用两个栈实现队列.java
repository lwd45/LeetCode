package easy;

import java.util.Stack;

/**
 * Created by lwd at 2021/7/3
 *
 * @Description:
 */
public class 剑指Offer_09_用两个栈实现队列 {
    Stack<Integer> stack1, stack2;

    public CQueue() {
        stack1 = new Stack<>();
        stack2 = new Stack<>();
    }

    public void appendTail(int value) {
        stack2.push(value);
    }

    public int deleteHead() {
        if (stack1.isEmpty() && stack2.isEmpty()) return -1;
        else if (stack1.isEmpty()) {
            while (!stack2.isEmpty()) {
                stack1.add(stack2.pop());
            }
        }
        return stack1.pop();
    }
}

