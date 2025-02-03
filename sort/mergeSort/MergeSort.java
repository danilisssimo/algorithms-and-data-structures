package sort.mergeSort;

import java.util.Arrays;

public class MergeSort {
    
    public static int[] sorted(int[] array) {
        if (array.length < 2) { return array; }
        if (array.length == 2) {
            if (array[1] < array[0]) {
                array[0] += array[1];
                array[1] = array[0] - array[1];
                array[0] -= array[1];
            }
            return array;
        }
        int middle = array.length/2;
        int[] left = Arrays.copyOfRange(array, 0, middle);
        int[] right = Arrays.copyOfRange(array, middle, array.length);
        left = sorted(left);
        right = sorted(right);
        int leftIndex = 0;
        int rightIndex = 0;
        int[] merged = new int[right.length+left.length];
        while (left.length + right.length > leftIndex + rightIndex) {
            if (rightIndex >= right.length || (leftIndex < left.length && left[leftIndex] <= right[rightIndex])) {
                merged[leftIndex+rightIndex] = left[leftIndex];
                leftIndex++;
            } else {
                merged[leftIndex+rightIndex] = right[rightIndex];
                rightIndex++;
            }
        }
        return merged;
    }
}