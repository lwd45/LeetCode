package hard;

/**
 * Created by lwd at 2021/8/28
 *
 * @Description:
 */
public class LeetCode_42_接雨水 {
    public int trap(int[] height) {
        int left = 0, right = height.length - 1;
        int leftMax = 0, rightMax = height.length - 1;
        int ans = 0;

        while (left <= right) {
            if (height[leftMax] > height[rightMax]) {
                if (height[right] >= height[rightMax]) {
                    rightMax = right;
                } else {
                    ans += height[rightMax] - height[right];
                }
                right--;
            } else {
                if (height[left] >= height[leftMax]) {
                    leftMax = left;
                } else {
                    ans += height[leftMax] - height[left];
                }
                left++;
            }
        }
        return ans;
    }
}
