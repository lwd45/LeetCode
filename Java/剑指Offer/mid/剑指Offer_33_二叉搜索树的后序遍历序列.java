package mid;

/**
 * Created by lwd at 2021/7/29
 *
 * @Description:
 */
public class 剑指Offer_33_二叉搜索树的后序遍历序列 {

    public boolean verifyPostorder(int[] postorder) {
        return verifyPostorder(postorder, 0, postorder.length - 1);
    }

    public boolean verifyPostorder(int[] postorder, int l, int r) {
        if (l > r) return true;

        int root = postorder[r];

        int i = l;
        while (i < r && postorder[i] < root) i++;

        int j = i;
        while (j < r) {
            if (postorder[j] < root) return false;
            j++;
        }

        return verifyPostorder(postorder, l, i - 1) && verifyPostorder(postorder, i, r - 1);
    }

}
