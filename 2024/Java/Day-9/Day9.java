import java.io.File;
import java.io.IOException;
import java.util.Scanner;

public class Day9 {
    public static void main(String[] args) {
        try {
            File file = new File("Input.txt");
            Scanner scanner = new Scanner(file);

            // Reading and storing the input
            String input = "";
            while (scanner.hasNextLine()) {
                input += scanner.nextLine();
            }

            // Representing as individual blocks
            String fileBlocks = "";
            boolean isEmpty = false;
            int k = -1;
            for (int i = 0; i < input.length(); i++) {
                int d = Integer.parseInt(String.valueOf(input.charAt(i)));
                String fillWith;
                if (!isEmpty) {
                    k++;
                    fillWith = String.valueOf(k);
                } else {
                    fillWith = ".";
                }
                for (int j = 0; j < d; j++) {
                    fileBlocks += fillWith;
                }
                isEmpty = !isEmpty;
            }

            System.out.println(fileBlocks);

            for (int i = fileBlocks.length() - 1; i > 0; i--) {

                if (fileBlocks.charAt(i) == '.') {
                    continue;
                }

                System.out.println(fileBlocks.charAt(i));
                fileBlocks = fileBlocks.replaceFirst("\\.", String.valueOf(fileBlocks.charAt(i)));

                System.out.println(fileBlocks);
            }

            // Closing scanner to prevent resource leak
            scanner.close();
        } catch (IOException e) {
        }

    }
}
