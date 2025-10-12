// Java implementation
import java.util.*;

public class Solution {
    static final int MOD = 1000000007;
    public int magicalSum(int m, int k, int[] nums) {
        int n = nums.length;
        long[][] C = new long[m+1][m+1];
        for(int i=0;i<=m;i++){
            C[i][0] = C[i][i] = 1;
            for(int j=1;j<i;j++){
                C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD;
            }
        }
        long[][] powA = new long[n][m+1];
        for(int i=0;i<n;i++){
            powA[i][0] = 1;
            for(int t=1;t<=m;t++){
                powA[i][t] = (powA[i][t-1] * (nums[i] % MOD)) % MOD;
            }
        }

        int M = m;
        long[][][] cur = new long[M+1][M+1][M+1];
        cur[M][0][0] = 1;

        for(int i=0;i<n;i++){
            long[][][] nxt = new long[M+1][M+1][M+1];
            for(int r=0;r<=M;r++){
                for(int carry=0; carry<=M; carry++){
                    for(int ones=0; ones<=M; ones++){
                        long val = cur[r][carry][ones];
                        if(val==0) continue;
                        for(int t=0;t<=r;t++){
                            int newr = r - t;
                            int sum = carry + t;
                            int bit = (sum & 1);
                            int newones = ones + bit;
                            if(newones > M) continue;
                            int newcarry = sum >> 1;
                            long mult = (C[r][t] * powA[i][t]) % MOD;
                            long add = (val * mult) % MOD;
                            nxt[newr][newcarry][newones] += add;
                            if(nxt[newr][newcarry][newones] >= MOD) nxt[newr][newcarry][newones] -= MOD;
                        }
                    }
                }
            }
            cur = nxt;
        }

        long ans = 0;
        for(int carry=0; carry<=M; carry++){
            for(int ones=0; ones<=M; ones++){
                long val = cur[0][carry][ones];
                if(val==0) continue;
                int extra = Integer.bitCount(carry);
                if(ones + extra == k){
                    ans = (ans + val) % MOD;
                }
            }
        }
        return (int)ans;
    }
}
