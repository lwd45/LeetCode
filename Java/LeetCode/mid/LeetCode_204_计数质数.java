package mid;

import java.util.Arrays;

/**
 * Created by lwd at 2021/11/13
 *
 * @Description:
 */
public class LeetCode_204_计数质数 {
    public int countPrimes(int n) {
        int[] primes = new int[n];
        Arrays.fill(primes, 1);
        int ans = 0;
        for (int i = 2; i < n; i++) {
            if (primes[i] == 1) {
                ans++;
                for (long j = (long) i * i; j < n; j += i) {
                    primes[(int) j] = 0;
                }
            }
        }
        return ans;
    }
}
