import java.io.File;
import java.util.ArrayList;
import java.util.Scanner;

class Day1 {
    public static void main(String args[]) throws Exception {
        File file = new File("Input.txt");

        Scanner sc = new Scanner(file);

        ArrayList<Integer> list1 = new ArrayList<>();
        ArrayList<Integer> list2 = new ArrayList<>();

        while (sc.hasNextLine()) {
            String[] nums = sc.nextLine().split("   ");
            int n1 = Integer.parseInt(nums[0]);
            int n2 = Integer.parseInt(nums[1]);
            list1.add(n1);
            list2.add(n2);
        }
        sc.close();
        
        // Collections.sort(list1);
        // Collections.sort(list2);
        // int dif = 0;
        // for (int i = 0; i < list1.size(); i++) {
        //     dif += Math.abs(list1.get(i)-list2.get(i));
        // }
        // System.out.println(dif);

        int similarity_score = 0;

        for (int num : list1) {
            similarity_score += getTimes(num, list2);
        }
        
        System.out.println(similarity_score);

    }

    public static int getTimes(int n, ArrayList<Integer> list) {
        int k = 0;
        for (int checkNum : list) {
            if (n == checkNum) {
                k++;
            }
        }
        return k*n;
    }
}