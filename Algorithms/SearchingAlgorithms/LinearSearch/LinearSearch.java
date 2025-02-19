package SearchingAlgorithms.LinearSearch;

public class LinearSearch {

    public static int indexOf(int[] array, int searching) {
        for (int index = 0; index < array.length; index++) {
            if (searching == array[index]) {
                return index;
            }
        }
        return -1;
    }
}
