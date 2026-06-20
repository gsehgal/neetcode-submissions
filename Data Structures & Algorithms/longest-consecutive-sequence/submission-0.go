type UnionFind struct {
	par map[int]int
	rank map[int]int
}

func NewUnion(nums []int) *UnionFind {
	uf := &UnionFind {
		par: make(map[int]int),
		rank: make(map[int]int),
	}
	for _,num := range nums {
		uf.par[num]=num
		uf.rank[num]=1
	}
	return uf
}

func (uf *UnionFind) Find(x int) (int,error) {
	_,ok := uf.par[x]
	if (!ok) {
		return -1,fmt.Errorf("Not found")
	}
    for x!=uf.par[x] {
		uf.par[x] = uf.par[uf.par[x]]
		x = uf.par[x]
	}	
	return x,nil
}

func (uf *UnionFind) Union(x,y int) {
	rootX ,_ := uf.Find(x)
	rootY ,err := uf.Find(y)
	if (err!=nil){
		return
	}
	if (rootX==rootY) {
		return
	}
	if (uf.rank[rootX]>uf.rank[rootY]){
		uf.par[rootX]=rootY
		uf.rank[rootX]+=uf.rank[rootY]
	} else {
		uf.par[rootY]=rootX
		uf.rank[rootY]+=uf.rank[rootX]
	}
}

func longestConsecutive(nums []int) int {
		uf := NewUnion(nums)
		for _,num := range nums {
			uf.Union(num,num-1)
			uf.Union(num,num+1)
		}
		root := make(map[int]int)
		visited := make(map[int]bool)
		for _,num := range nums {
			if _,ok := visited[num]; ok {
				continue
			}
			visited[num]=true
			r,_ := uf.Find(num)
			root[r] += 1
		}
		m := 0
		for _,count := range root{
			if count>m {
				m = count
			}
		}
		return m
		    
		

}
