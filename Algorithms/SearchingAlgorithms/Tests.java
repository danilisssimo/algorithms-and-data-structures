package SearchingAlgorithms;

import SearchingAlgorithms.LinearSearch.LinearSearch;
import SearchingAlgorithms.BinarySearch.BinarySearch;
import java.util.*;

class TestCase {
    
    private int[] array;
    private int searching;
    private int result;

    public TestCase(int[] testArray, int testSearching, int result) {
        this.array = testArray;
        this.searching = testSearching;
        this.result = result;
    }

    public int[] getTestArray() { return array; }
    public int getTestSearching() { return searching; }
    public int getResult() { return result; }

}

public class Tests {

    public static TestCase[] linearTestCases = {
        new TestCase(new int[]{5, 2, 4, 6, 1, 3}, 1, 4),  // Искомое число 1 находится на индексе 4
        new TestCase(new int[]{5, 2, 4, 6, 1, 3}, 6, 3),  // Искомое число 6 находится на индексе 3
        new TestCase(new int[]{5, 2, 4, 6, 1, 3}, 3, 5),  // Искомое число 3 находится на индексе 5
        new TestCase(new int[]{5, 2, 4, 6, 1, 3}, 7, -1), // Искомое число 7 отсутствует в массиве
        new TestCase(new int[]{}, 1, -1), // Пустой массив, искомое число отсутствует
        new TestCase(new int[]{5}, 5, 0), // Искомое число 5 находится на индексе 0
        new TestCase(new int[]{5}, 1, -1), // Искомое число 1 отсутствует в массиве
        new TestCase(new int[]{5, 2, 4, 6, 1, 3, 2}, 2, 1), // Искомое число 2 находится на индексе 1 (первое вхождение)
        new TestCase(new int[]{5, 2, 4, 6, 1, 3, 2}, 8, -1), // Искомое число 8 отсутствует в массиве
        new TestCase(new int[]{-5, -2, -4, -6, -1, -3}, -1, 4), // Искомое число -1 находится на индексе 4
        new TestCase(new int[]{-5, -2, -4, -6, -1, -3}, 0, -1), // Искомое число 0 отсутствует в массиве
        new TestCase(new int[]{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 50, 4), // Искомое число 50 находится на индексе 4
        new TestCase(new int[]{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 110, -1) // Искомое число 110 отсутствует в массиве
    };

    public static TestCase[] binaryTestCases = {
        // Не срабатывает 2 тест
        new TestCase(new int[]{1, 2, 3, 4, 5, 6}, 1, 0),  // Искомое число 1 находится на индексе 0
        new TestCase(new int[]{1, 2, 3, 4, 5, 6}, 6, 5),  // Искомое число 6 находится на индексе 5
        new TestCase(new int[]{1, 2, 3, 4, 5, 6}, 3, 2),  // Искомое число 3 находится на индексе 2
        new TestCase(new int[]{1, 2, 3, 4, 5, 6}, 7, -1), // Искомое число 7 отсутствует в массиве
        new TestCase(new int[]{}, 1, -1), // Пустой массив, искомое число отсутствует
        new TestCase(new int[]{5}, 5, 0), // Искомое число 5 находится на индексе 0
        new TestCase(new int[]{5}, 1, -1), // Искомое число 1 отсутствует в массиве
        new TestCase(new int[]{1, 2, 2, 3, 4, 5}, 2, 1), // Искомое число 2 находится на индексе 1 (первое вхождение)
        new TestCase(new int[]{1, 2, 2, 3, 4, 5}, 6, -1), // Искомое число 6 отсутствует в массиве
        new TestCase(new int[]{-6, -5, -4, -3, -2, -1}, -1, 5), // Искомое число -1 находится на индексе 5
        new TestCase(new int[]{-6, -5, -4, -3, -2, -1}, 0, -1), // Искомое число 0 отсутствует в массиве
        new TestCase(new int[]{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 50, 4), // Искомое число 50 находится на индексе 4
        new TestCase(new int[]{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 110, -1), // Искомое число 110 отсутствует в массиве
        new TestCase(new int[]{1, 3, 5, 7, 9}, 5, 2), // Искомое число 5 находится на индексе 2
        new TestCase(new int[]{1, 3, 5, 7, 9}, 8, -1), // Искомое число 8 отсутствует в массиве
        new TestCase(new int[]{2, 4, 6, 8, 10, 12}, 8, 3), // Искомое число 8 находится на индексе 3
        new TestCase(new int[]{2, 4, 6, 8, 10, 12}, 1, -1)  // Искомое число 1 отсутствует в массиве
    };

    
    public static void main(String[] args) {
        linearSearchTest();
        binarySearchTest();
    }

    public static void binarySearchTest() {
        System.out.println("\n--- TEST FOR BINARY SEARCH ---");
        for (int i = 0; i < binaryTestCases.length; i++) {
            int[] currentTest = binaryTestCases[i].getTestArray().clone();
            int result = BinarySearch.indexOf(currentTest, binaryTestCases[i].getTestSearching());
            boolean testCaseResult = (result == binaryTestCases[i].getResult());
            System.out.printf(
                "Array: %s, Searching: %d, Result: %d, TestResult: %s\n", 
                Arrays.toString(binaryTestCases[i].getTestArray()),
                binaryTestCases[i].getTestSearching(),
                result,
                Boolean.toString(testCaseResult)
            );
        }
    }

    public static void linearSearchTest() {
        System.out.println("\n--- TEST FOR LINEAT SEARCH ---");
        for (int i = 0; i < linearTestCases.length; i++) {
            int[] currentTest = linearTestCases[i].getTestArray().clone();
            int result = LinearSearch.indexOf(currentTest, linearTestCases[i].getTestSearching());
            boolean testCaseResult = (result == linearTestCases[i].getResult());
            System.out.printf(
                "Array: %s, Searching: %d, Result: %d, TestResult: %s\n", 
                Arrays.toString(linearTestCases[i].getTestArray()),
                linearTestCases[i].getTestSearching(),
                result,
                Boolean.toString(testCaseResult)
            );
        }
    }

   
}
