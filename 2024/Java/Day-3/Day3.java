import java.io.File;
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

        // String regex = "mul\\((\\d{1,3}),(\\d{1,3})\\)";
        String regex = "(mul\\((\\d{1,3}),(\\d{1,3})\\)|don't\\(\\)|do\\(\\))";
        Pattern pattern = Pattern.compile(regex);
        Matcher matcher = pattern.matcher(str);

        int sum = 0;
        boolean flag = true;

        while (matcher.find()) {
            String match = matcher.group(0);

            if (match.equals("do()")) {
                flag = true;
            }
            else if (match.equals("don't()")) {
                flag = false;
            }
            else {
                if (flag) {
                    int x = Integer.parseInt(matcher.group(2));
                    int y = Integer.parseInt(matcher.group(3));
                    sum += x * y;
                }
            }
        }
        
        System.out.println(sum);
    }
}
