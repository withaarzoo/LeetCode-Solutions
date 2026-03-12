class Solution:
    def maxStability(self, n: int, edges: List[List[int]], k: int) -> int:

        parent=list(range(n))
        rank=[0]*n

        def find(x):
            if parent[x]!=x:
                parent[x]=find(parent[x])
            return parent[x]

        def union(a,b):
            ra,rb=find(a),find(b)

            if ra==rb:
                return False

            if rank[ra]<rank[rb]:
                ra,rb=rb,ra

            parent[rb]=ra

            if rank[ra]==rank[rb]:
                rank[ra]+=1

            return True


        comp=n
        mandatory_min=float('inf')

        optional=[]

        for u,v,s,m in edges:

            if m==1:
                if not union(u,v):
                    return -1

                comp-=1
                mandatory_min=min(mandatory_min,s)

            else:
                optional.append((u,v,s))


        optional.sort(key=lambda x:-x[2])

        used=[]

        for u,v,s in optional:

            if union(u,v):
                used.append(s)
                comp-=1
                if comp==1:
                    break

        if comp>1:
            return -1


        used.sort()

        ans=mandatory_min

        for w in used:

            val=w

            if k>0:
                val*=2
                k-=1

            if ans==float('inf'):
                ans=val
            else:
                ans=min(ans,val)

        return ans