import java.io.BufferedReader;
import java.io.FileReader;

public class FileUtil {
    public static int countLines(String filename) throws Exception {
        try (BufferedReader br = new BufferedReader(new FileReader(filename))) {
            int count = 0;
            while (br.readLine() != null) count++;
            return count;
        }
    }
}