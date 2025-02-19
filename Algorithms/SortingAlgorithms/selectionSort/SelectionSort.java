package SortingAlgorithms.selectionSort;

public class SelectionSort {
    
    public static void sorted(int[] array) {
        for (int index = 0; index < array.length; index++) {
            int currentMinIndex = index;
            int subIndex = index + 1;
            while (subIndex < array.length) {
                if (array[subIndex] < array[currentMinIndex]) {
                    currentMinIndex = subIndex;
                }
                subIndex++;
            }
            if (index != currentMinIndex) {
                array[index] += array[currentMinIndex];
                array[currentMinIndex] = array[index] - array[currentMinIndex];
                array[index] -= array[currentMinIndex];
            }
        }
    }
}