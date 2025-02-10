package SortingAlgorithms.quickSort;

import java.util.Arrays;

public class QuickSort {
    
    public static int[] sorted(int[] array) {
        if (array.length <= 1) { return array; }
        int supportIndex = array.length - 1;
        int index = 0;
        while (index < supportIndex) {
            if (array[index] > array[supportIndex]) {
                int supportNum = array[supportIndex];
                int toSwapEl = array[supportIndex - 1];
                array[supportIndex] = array[index];
                array[index] = toSwapEl;
                array[--supportIndex] = supportNum;
            } else {
                index++;
            }
        }
        int[] leftArray = Arrays.copyOfRange(array, 0, supportIndex);
        int[] rightArray = Arrays.copyOfRange(array, supportIndex+1, array.length);
        return getConcatArrays(sorted(leftArray), array[supportIndex], sorted(rightArray));
    }

    private static int[] getConcatArrays(int[] left, int middleEl, int[] right) {
        // Somrthing wrong with this method
        int[] newArray = new int[left.length + right.length + 1];
        int index = 0;
        while (index < left.length) {
            newArray[index] = left[index];
            index++;
        }
        newArray[left.length] = middleEl;
        index++;
        while (index - left.length - 1 < right.length) {
            newArray[index] = right[index - left.length - 1];
            index++;
        }
        return newArray;
    }
}
