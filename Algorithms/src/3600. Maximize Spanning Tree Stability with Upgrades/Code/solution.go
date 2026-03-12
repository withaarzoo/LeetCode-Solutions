func maxStability(n int, edges [][]int, k int) int {

	parent:=make([]int,n)
	rank:=make([]int,n)

	for i:=0;i<n;i++{
		parent[i]=i
	}

	var find func(int) int
	find = func(x int) int{
		if parent[x]!=x{
			parent[x]=find(parent[x])
		}
		return parent[x]
	}

	union := func(a,b int) bool{
		ra,rb:=find(a),find(b)

		if ra==rb{
			return false
		}

		if rank[ra]<rank[rb]{
			ra,rb=rb,ra
		}

		parent[rb]=ra

		if rank[ra]==rank[rb]{
			rank[ra]++
		}

		return true
	}

	comp:=n
	mandatoryMin:=int(^uint(0)>>1)

	var optional [][]int

	for _,e:=range edges{

		u,v,s,m:=e[0],e[1],e[2],e[3]

		if m==1{

			if !union(u,v){
				return -1
			}

			comp--
			if s<mandatoryMin{
				mandatoryMin=s
			}

		}else{
			optional=append(optional,e)
		}
	}

	sort.Slice(optional,func(i,j int) bool{
		return optional[i][2]>optional[j][2]
	})

	var used []int

	for _,e:=range optional{

		if union(e[0],e[1]){
			used=append(used,e[2])
			comp--

			if comp==1{
				break
			}
		}
	}

	if comp>1{
		return -1
	}

	sort.Ints(used)

	ans:=mandatoryMin

	for _,w:=range used{

		val:=w

		if k>0{
			val*=2
			k--
		}

		if ans==int(^uint(0)>>1){
			ans=val
		}else if val<ans{
			ans=val
		}
	}

	return ans
}