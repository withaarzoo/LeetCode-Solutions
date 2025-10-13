import java.util.*;

class Solution {
    public List<String> removeAnagrams(String[] words) {
        List<String> result = new ArrayList<>();
        String prevSig = "";

        for (String w : words) {
            char[] arr = w.toCharArray();
            Arrays.sort(arr);               // signature: sorted chars
            String sig = new String(arr);
            if (!sig.equals(prevSig)) {
                result.add(w);
                prevSig = sig;
            }
        }
        return result;
    }
}
