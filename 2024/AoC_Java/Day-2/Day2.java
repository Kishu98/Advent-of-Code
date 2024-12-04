import java.io.File;
import java.util.ArrayList;
import java.util.Scanner;

class Day2 {
    public static void main(String[] args) throws Exception {
        File file = new File("Input.txt");

        Scanner sc = new Scanner(file);
        int total = 0;

        while (sc.hasNextLine()) {
            String[] nums = sc.nextLine().split(" ");
            ArrayList<Integer> arr = new ArrayList<>();
            for (String num : nums) {
                arr.add(Integer.valueOf(num));
            }
            if (validate(arr) || validate2(arr)) {
                total++;
            }
        }
        sc.close();

        System.out.println(total);
    }

    public static boolean validate2(ArrayList<Integer> arr) {
        for (int num : arr) {
            ArrayList<Integer> copy = new ArrayList<>(arr);
            copy.remove(Integer.valueOf(num));
            if (validate(copy)) {
                return true;
            }
        }
        return false;
    }

    public static boolean validate(ArrayList<Integer> arr) {
        return (checkInc(arr) || checkDec(arr)) && checkRange(arr);
    }

    public static boolean checkInc(ArrayList<Integer> arr) {
        for (int i = 0; i < arr.size()-1; i++) {
            if (arr.get(i) > arr.get(i+1)) {
                return false;
            }
        }
        return true;
    }

    public static boolean checkDec(ArrayList<Integer> arr) {
        for (int i = 0; i < arr.size()-1; i++ ) {
            if (arr.get(i) < arr.get(i+1)) {
                return false;
            }
        }
        return true;
    }

    public static boolean checkRange(ArrayList<Integer> arr) {
        for (int i = 0; i < arr.size()-1; i++) {
            if (!(((Math.abs(arr.get(i) - arr.get(i+1)) >= 1) && (Math.abs(arr.get(i) - arr.get(i+1)) <= 3)))) {
                return false;
            }
        }
        return true;
    }
}