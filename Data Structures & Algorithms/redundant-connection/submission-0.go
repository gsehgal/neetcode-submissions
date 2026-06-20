type UnionFind struct {
	par map[int]int
	rank map[int]int
}

func NewUnion(n int) *UnionFind {
    uf := &UnionFind{
        par : make(map[int]int),
        rank : make(map[int]int),
    }
    for i := 1; i<=n;i++ {
        uf.par[i]=i
        uf.rank[i]=0
    }
    return uf
}

func (uf *UnionFind) Find(x int) int {
    if uf.par[x] == x {
        return x
    }
    for x!=uf.par[x] {
        x=uf.par[x]
    }
    return x
}

func (uf *UnionFind) Union(x,y int) bool {
    x = uf.Find(x)
    y = uf.Find(y)

    if (x==y){
      return false
    }
    if (uf.rank[x]>uf.rank[y]) {
        uf.par[y]=x
        uf.rank[x]+=uf.rank[y]
    } else{
        uf.par[x]=y
        uf.rank[y]+=uf.rank[x]
    }
    return true
}

func findRedundantConnection(edges [][]int) []int {
    uf := NewUnion(len(edges))
    for _,edge := range edges {
        x,y := edge[0],edge[1]
        if (!uf.Union(x,y)){
            return edge
        }
    }
    return nil
         
}
