package easy;

/**
 * Created by lwd at 2021/7/15
 *
 * @Description:
 */
public class 剑指Offer_17_打印从1到最大的n位数 {
    public int[] printNumbers(int n) {
        StringBuilder maxStr = new StringBuilder();
        while (n-- > 0) {
            maxStr.append(9);
        }
        int maxNum = Integer.parseInt(maxStr.toString());
        int[] ans = new int[maxNum];
        for (int i = 1; i < maxNum; ++i) {
            ans[i - 1] = i;
        }
        return ans;
    }
}
