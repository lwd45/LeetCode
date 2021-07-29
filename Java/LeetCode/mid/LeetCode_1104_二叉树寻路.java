package mid;

import com.sun.org.apache.xpath.internal.operations.Or;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * Created by lwd at 2021/7/29
 *
 * @Description:
 */
public class LeetCode_1104_二叉树寻路 {
    public List<Integer> pathInZigZagTree(int label) {
        int level = 1, min = 1, max = 1;
        while (max < label) {
            level++;
            min *= 2; // 记录每一层最小值方便计算偏移量
            max = max * 2 + 1; // 记录每一层最大值方便计算偏移量
        }

        List<Integer> ans = new ArrayList<>(); // 存储的是倒序结果
        ans.add(label);
        if (level % 2 == 1) { // 奇数层，正常排序
            boolean reverse = true; // 上一层是否逆序
            while (level > 1) {
                level--;
                min = min / 2;
                max = max / 2;
                int parent = label / 2;
                if (reverse) {
                    parent = (max - parent + min);//如果是逆序，那么根据偏移量求当前真实编号
                }
                ans.add(parent);
                label = label / 2;
                reverse = !reverse;
            }
        } else {//偶数层，倒排序
            label = max - label + min;
            boolean reverse = false; // 上一层是否逆序
            while (level > 1) {
                level--;
                min = min / 2;
                max = max / 2;
                int parent = label / 2;
                if (reverse) {
                    parent = max - parent + min;// 如果是逆序，那么根据偏移量求当前真实编号
                }
                ans.add(parent);
                label = label / 2;
                reverse = !reverse;
            }
        }
        Collections.reverse(ans);
        return ans;
    }
}
