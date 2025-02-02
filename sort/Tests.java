package sort;

import sort.insertionSort.InsertionSort;
import java.util.*;

class TestCase {
    
    private int[] test;
    private int[] result;

    public TestCase(int[] test, int[] result) {
        this.test = test;
        this.result = result;
    }

    public int[] getTest() { return test; }

    public int[] getResult() { return result; }

}

public class Tests {

    public static TestCase[] testCases = {
        new TestCase(new int[]{5, 2, 4, 6, 1, 3}, new int[]{1, 2, 3, 4, 5, 6}),
        new TestCase(new int[]{31,41,59,26,41,58}, new int[]{26,31,41,41,58,59}),
    };

    
    public static void main(String[] args) {
        for (int i = 0; i < testCases.length; i++) {
            int[] currentTest = testCases[i].getTest().clone();
            InsertionSort.sorted(currentTest);
            boolean testCaseResult = Arrays.equals(currentTest, testCases[i].getResult());
            System.out.printf(
                "Array: %s, Sorted: %s, Result: %s\n", 
                Arrays.toString(testCases[i].getTest()),
                Arrays.toString(currentTest),
                Boolean.toString(testCaseResult)
            );
        }
    }
}
