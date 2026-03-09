class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:

        MOD = 10**9 + 7
        n = zero + one

        fact = [1]*(n+1)
        invFact = [1]*(n+1)

        for i in range(1,n+1):
            fact[i] = fact[i-1]*i % MOD

        invFact[n] = pow(fact[n], MOD-2, MOD)

        for i in range(n-1,-1,-1):
            invFact[i] = invFact[i+1]*(i+1) % MOD

        def C(n,k):
            if k<0 or k>n:
                return 0
            return fact[n]*invFact[k]%MOD*invFact[n-k]%MOD

        def F(N,K):
            if K<=0 or K>N:
                return 0

            ans=0
            maxJ=(N-K)//limit

            for j in range(maxJ+1):

                term=C(K,j)*C(N-j*limit-1,K-1)%MOD

                if j%2:
                    ans=(ans-term)%MOD
                else:
                    ans=(ans+term)%MOD

            return ans

        maxK=min(zero,one+1)

        oneWays=[0]*(maxK+3)

        for k in range(1,maxK+2):
            oneWays[k]=F(one,k)

        ans=0

        for k in range(1,maxK+1):

            z=F(zero,k)

            o=(oneWays[k-1]+2*oneWays[k]+oneWays[k+1])%MOD

            ans=(ans+z*o)%MOD

        return ans