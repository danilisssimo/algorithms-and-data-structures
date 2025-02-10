package SearchingAlgorithms.BinarySearch;

public class BinarySearch {
    
    public static int indexOf(int[] array, int searching) {
        int leftIndex = 0;
        int rightIndex = array.length - 1;
        int middleIndex = array.length/2;
        while (rightIndex - leftIndex > 2) {
            if (array[middleIndex] == searching) {
                return middleIndex;
            } else if (array[middleIndex] > searching) {
                leftIndex = middleIndex;
                middleIndex = (rightIndex - leftIndex)/2;
            } else {
                rightIndex = middleIndex;
                middleIndex = (rightIndex - leftIndex)/2;
            }
        }
        return -1;
    }
}
