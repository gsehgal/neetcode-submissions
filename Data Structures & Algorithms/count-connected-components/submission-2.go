type UnionFind struct{
	par map[int]int
	rank map[int]int
}

func NewUnion(n int) *UnionFind {
	uf := &UnionFind{
		par: make(map[int]int),
		rank: make(map[int]int),
	}
	for i:=0;i<n;i++ {
		uf.par[i]=i
		uf.rank[i]=1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int{
	for x!=uf.par[x] {
		uf.par[x]=uf.par[uf.par[x]]
		x=uf.par[x]
	}
	return x
}

func (uf *UnionFind) Union(x,y int){
	x = uf.Find(x)
	y = uf.Find(y)
	if x==y {
		return
	}
	if uf.rank[x]>uf.rank[y] {
		uf.par[y]=x
		uf.rank[x]+=uf.par[y]
	} else {
		uf.par[x]=y
		uf.rank[y]+=uf.par[x]
	}
}

func countComponents(n int, edges [][]int) int {
    uf := NewUnion(n)
	for _,edge := range edges {
		x,y := edge[0],edge[1]
		uf.Union(x,y)
	}
	root := make(map[int]int)
	for i:=0 ; i<n;i++ {
		r := uf.Find(i)
		root[r]=1
	}
	return len(root)
}
