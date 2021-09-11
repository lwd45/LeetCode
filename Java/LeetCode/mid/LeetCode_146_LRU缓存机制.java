package mid;

import java.util.HashMap;
import java.util.LinkedList;

/**
 * Created by lwd at 2021/9/11
 *
 * @Description:
 */
public class LeetCode_146_LRU缓存机制 {

    class LRUCache {
        class LinkedNode {
            int val;
            int key;

            LinkedNode pre;
            LinkedNode next;

            public LinkedNode() {
            }

            public LinkedNode(int key, int val) {
                this.key = key;
                this.val = val;
            }
        }

        private int capacity;
        private int size;
        private LinkedNode head, tail;
        private HashMap<Integer, LinkedNode> map;

        public LRUCache(int capacity) {
            this.capacity = capacity;
            size = 0;
            head = new LinkedNode();
            tail = new LinkedNode();
            head.next = tail;
            tail.pre = head;
            map = new HashMap<>();
        }

        public int get(int key) {
            LinkedNode node = map.getOrDefault(key, null);
            if (node != null) {
                removeNode(node);
                moveToHead(node);
                return node.val;
            }
            return -1;
        }

        public void put(int key, int value) {
            LinkedNode exit = map.getOrDefault(key, null);
            if (exit == null) {
                LinkedNode node = new LinkedNode(key, value);
                if (size == capacity) {
                    LinkedNode lastNode = tail.pre;
                    map.remove(lastNode.key);
                    removeNode(lastNode);
                } else {
                    size++;
                }
                map.put(key, node);
                moveToHead(node);
            } else {
                removeNode(exit);
                moveToHead(exit);
                exit.val = value;
            }
        }

        private void moveToHead(LinkedNode node) {
            node.next = head.next;
            node.next.pre = node;
            node.pre = head;
            head.next = node;
        }

        public void removeNode(LinkedNode node) {
            node.next.pre = node.pre;
            node.pre.next = node.next;
        }
    }

}
