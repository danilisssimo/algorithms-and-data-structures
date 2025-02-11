package SearchingAlgorithms.BinarySearch;

public class BinarySearch {
    
    public static int indexOf(int[] array, int searching) {
        int leftIndex = 0;
        int rightIndex = array.length - 1;
        int middleIndex = array.length/2;
        while (leftIndex <= rightIndex) {
            if (array[middleIndex] == searching) {
                return middleIndex;
            } else if (array[middleIndex] > searching) {
                rightIndex = middleIndex;
                middleIndex = (rightIndex - leftIndex)/2;
            } else {
                leftIndex = middleIndex;
                middleIndex = (rightIndex + leftIndex)/2;
            }
        }
        return -1;
    }
}
