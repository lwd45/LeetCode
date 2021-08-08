package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/8/8
 *
 * @Description:
 */
public class 剑指Offer_60_n个骰子的点数 {
    public double[] dicesProbability(int n) {
        double[] temp = new double[6 * n + 1];
        Arrays.fill(temp, 1.0 / 6);

        for (int i = 2; i <= n; ++i) {
            int minIndex = i - 1, maxIndex = 6 * (i - 1);
            double[] temp2 = new double[6 * n + 1];
            for (int j = 1; j <= 6; j++) {
                for (int k = minIndex; k <= maxIndex; k++) {
                    temp2[k + j] += temp[k] * (1.0 / 6);
                }
            }
            temp = temp2;
        }
        double[] ans = new double[5 * n + 1];
        for (int i = 0, j = n; i < ans.length; i++, j++) {
            ans[i] = temp[j];
        }
        return ans;
    }
}
