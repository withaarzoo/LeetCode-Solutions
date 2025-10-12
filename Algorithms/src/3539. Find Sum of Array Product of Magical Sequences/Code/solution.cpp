/* C++ implementation */
#include <bits/stdc++.h>
using namespace std;
using int64 = long long;
const int MOD = 1000000007;

int64 modMul(int64 a, int64 b){ return (a*b) % MOD; }

class Solution {
public:
    int magicalSum(int m, int k, vector<int>& nums) {
        int n = nums.size();
        // Precompute combinations C up to m
        vector<vector<int64>> C(m+1, vector<int64>(m+1,0));
        for(int i=0;i<=m;i++){
            C[i][0] = C[i][i] = 1;
            for(int j=1;j<i;j++){
                C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD;
            }
        }
        // Precompute powA[i][t] = nums[i]^t mod MOD for t=0..m
        vector<vector<int64>> powA(n, vector<int64>(m+1,1));
        for(int i=0;i<n;i++){
            int64 a = nums[i] % MOD;
            for(int t=1;t<=m;t++){
                powA[i][t] = (powA[i][t-1] * a) % MOD;
            }
        }

        int M = m;
        // dp[r][carry][ones]
        vector<vector<vector<int64>>> cur(M+1, vector<vector<int64>>(M+1, vector<int64>(M+1,0)));
        cur[M][0][0] = 1;

        for(int i=0;i<n;i++){
            vector<vector<vector<int64>>> nxt(M+1, vector<vector<int64>>(M+1, vector<int64>(M+1,0)));
            for(int r=0;r<=M;r++){
                for(int carry=0; carry<=M; carry++){
                    for(int ones=0; ones<=M; ones++){
                        int64 val = cur[r][carry][ones];
                        if(val==0) continue;
                        for(int t=0;t<=r;t++){
                            int newr = r - t;
                            int sum = carry + t;
                            int bit = sum & 1;
                            int newones = ones + bit;
                            if(newones > M) continue;
                            int newcarry = sum >> 1;
                            int64 mult = (C[r][t] * powA[i][t]) % MOD;
                            int64 add = (val * mult) % MOD;
                            nxt[newr][newcarry][newones] += add;
                            if(nxt[newr][newcarry][newones] >= MOD) nxt[newr][newcarry][newones] -= MOD;
                        }
                    }
                }
            }
            cur.swap(nxt);
        }

        int64 ans = 0;
        for(int carry=0; carry<=M; carry++){
            for(int ones=0; ones<=M; ones++){
                int64 val = cur[0][carry][ones];
                if(val==0) continue;
                int extra = __builtin_popcount((unsigned)carry);
                if(ones + extra == k){
                    ans = (ans + val) % MOD;
                }
            }
        }
        return (int)ans;
    }
};
