package sort;

import sort.insertionSort.InsertionSort;
import sort.mergeSort.MergeSort;
import sort.selectionSort.SelectionSort;
import sort.quickSort.QuickSort;
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
        new TestCase(new int[]{}, new int[]{}), // Пустой массив
        new TestCase(new int[]{42}, new int[]{42}), // Массив из одного элемента
        new TestCase(new int[]{1, 2, 3, 4, 5}, new int[]{1, 2, 3, 4, 5}), // Уже отсортированный массив
        new TestCase(new int[]{5, 4, 3, 2, 1}, new int[]{1, 2, 3, 4, 5}), // Массив, отсортированный в обратном порядке
        new TestCase(new int[]{3, 1, 2, 3, 1}, new int[]{1, 1, 2, 3, 3}), // Массив с повторяющимися элементами
        new TestCase(new int[]{-3, -1, -2, -4}, new int[]{-4, -3, -2, -1}), // Массив с отрицательными числами
        new TestCase(new int[]{-3, 0, 2, -1, 4}, new int[]{-3, -1, 0, 2, 4}), // Массив с отрицательными, положительными числами и нулем
        new TestCase(new int[]{9, 3, 7, 5, 6, 4, 8, 2, 1}, new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9}), // Массив с большим количеством элементов
        new TestCase(new int[]{7, 7, 7, 7}, new int[]{7, 7, 7, 7}), // Массив с одинаковыми элементами
        new TestCase(new int[]{1000000, 999999, 1000001}, new int[]{999999, 1000000, 1000001}) // Массив с большими числами
    };

    
    public static void main(String[] args) {
        insertionSortTest();
        mergeSortTest();
        selectionSortTest();
        quickSortTest();
    }

    public static void insertionSortTest() {
        System.out.println("\n--- TEST FOR INSERTION SORT ---");
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

    public static void mergeSortTest() {
        System.out.println("\n--- TEST FOR MERGE SORT ---");
        for (int i = 0; i < testCases.length; i++) {
            int[] currentTest = testCases[i].getTest().clone();
            currentTest = MergeSort.sorted(currentTest);
            boolean testCaseResult = Arrays.equals(currentTest, testCases[i].getResult());
            System.out.printf(
                "Array: %s, Sorted: %s, Result: %s\n", 
                Arrays.toString(testCases[i].getTest()),
                Arrays.toString(currentTest),
                Boolean.toString(testCaseResult)
            );
        }
    }

    public static void selectionSortTest() {
        System.out.println("\n--- TEST FOR SELECTION SORT ---");
        for (int i = 0; i < testCases.length; i++) {
            int[] currentTest = testCases[i].getTest().clone();
            SelectionSort.sorted(currentTest);
            boolean testCaseResult = Arrays.equals(currentTest, testCases[i].getResult());
            System.out.printf(
                "Array: %s, Sorted: %s, Result: %s\n", 
                Arrays.toString(testCases[i].getTest()),
                Arrays.toString(currentTest),
                Boolean.toString(testCaseResult)
            );
        }
    }

    public static void quickSortTest() {
        System.out.println("\n--- TEST FOR QUICK SORT ---");
        for (int i = 0; i < testCases.length; i++) {
            int[] currentTest = testCases[i].getTest().clone();
            currentTest = QuickSort.sorted(currentTest);
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
