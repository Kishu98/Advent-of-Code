
import java.io.File;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day3 {

    public static void main(String[] args) throws Exception {
        File file = new File("Input.txt");
        Scanner sc = new Scanner(file);

        String str = "";
        while (sc.hasNextLine()) {
            str += sc.nextLine();
        }
        sc.close();

        String regex = "mul\\(\\d{1,3},\\d{1,3}\\)|don't\\(\\)|do\\(\\)";
        boolean flag = true;

        Pattern pattern = Pattern.compile(regex);
        Matcher matcher = pattern.matcher(str);

        ArrayList<String> arr = new ArrayList<>();

        while (matcher.find()) {
            if ("don't()".equals(matcher.group())) {
                flag = false;
            } else if ("do()".equals(matcher.group())) {
                flag = true;
            } else {
                if (flag) {
                    String subRegex = "\\d{1,3}";
                    Pattern p = Pattern.compile(subRegex);
                    // System.out.println(matcher.group());
                    Matcher m = p.matcher(matcher.group());
                    String tstr = "";
                    while (m.find()) {
                        tstr += m.group() + " ";
                    }
                    if (tstr.length() > 0) {
                        arr.add(tstr.trim());
                    }
                }
            }
            // System.out.println(arr);
        }

        int sum = 0;

        for (int i = 0; i < arr.size(); i++) {
            String mult = arr.get(i);
            String[] nums = mult.split(" ");
            sum += Integer.parseInt(nums[0]) * Integer.parseInt(nums[1]);
        }

        System.out.println(sum);
    }
}
