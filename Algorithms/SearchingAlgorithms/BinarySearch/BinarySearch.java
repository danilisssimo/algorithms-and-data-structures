package SearchingAlgorithms.BinarySearch;

public class BinarySearch {
    
    public static int indexOf(int[] array, int searching) {
        int leftIndex = 0;
        int rightIndex = array.length - 1;
        int middleIndex = array.length/2;
        while (leftIndex <= rightIndex) {
            if (array[middleIndex] == searching) return middleIndex;

            if (array[middleIndex] > searching) rightIndex = middleIndex; 
            else leftIndex = middleIndex;

            int diffOfIndexies = rightIndex - leftIndex;
            if (diffOfIndexies > 1) middleIndex = (diffOfIndexies)/2 + leftIndex;
            else if (array[leftIndex] == searching) return leftIndex;
            else if (array[rightIndex] == searching) return rightIndex;
            else break;
        }
        return -1;
    }
}
