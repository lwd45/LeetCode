package mid;

/**
 * Created by lwd at 2021/9/26
 *
 * @Description:
 */
public class LeetCode_287_寻找重复数 {
    public int findDuplicate(int[] nums) {
        int length = nums.length;
        for (int i = 0; i < length; i++) {
            if (nums[i] == i + 1) continue;
            while (nums[i] != i + 1 && nums[nums[i] - 1] != nums[i]) {
                int t = nums[nums[i] - 1];
                nums[nums[i] - 1] = nums[i];
                nums[i] = t;
            }
            if (nums[i] != i + 1 && nums[nums[i] - 1] == nums[i]) return nums[i];
        }
        return 0;
    }
}
