package SortingAlgorithms.insertionSort;

public class InsertionSort {

   public static void sorted(int[] array) {
      for(int index = 1; index < array.length; ++index) {
         int num = array[index];
         int j = index - 1;
         while (j >= 0 && num < array[j]) {
            int tmp = array[j];
            array[j] = num;
            array[j+1] = tmp;
            j--;
         }
      }

   }
}
