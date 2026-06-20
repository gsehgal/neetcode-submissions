type UnionFind struct {
	par map[int]int
	rank map[int]int
}

func NewUnion(n int) *UnionFind {
	uf := &UnionFind{
		par : make(map[int]int),
		rank : make(map[int]int),
	}
	for i:=0;i<n;i++ {
		uf.par[i]=i
		uf.rank[i]=1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if x==uf.par[x] {
		return x
	}
	for x!=uf.par[x] {
		uf.par[x] = uf.par[uf.par[x]]
		x=uf.par[x]
	}
	return x
}

func (uf *UnionFind) Union(x,y int) {
	x = uf.Find(x)
	y = uf.Find(y)

	if (x==y){
		return
	}
	if (uf.rank[x]>uf.rank[y]){
		uf.par[y]=x
		uf.rank[x]+=uf.rank[y]
	}else {
		uf.par[x]=y
		uf.rank[y]+=uf.rank[x]
	}
}

func accountsMerge(accounts [][]string) [][]string {
	uf := NewUnion(len(accounts))
	emailToAcc := make(map[string]int)
	for i, account := range accounts {
		for j:=1;j<len(account);j++ {
			email := account[j]
			idx,ok := emailToAcc[email]
			if !ok {
				emailToAcc[email]=i
			} else {
				uf.Union(i,idx)
			}
		}
	}
	emailGroup := make(map[int][]string)
	for email, idx := range emailToAcc {
		fmt.Printf("email %s : %d\n", email, idx)
		root := uf.Find(idx)
		emailGroup[root] = append(emailGroup[root],email)
	}
	var res [][]string
	for idx,emails := range emailGroup {
		name := accounts[idx][0]
		sort.Strings(emails)
		row := append([]string{name},emails...)
		res = append(res,row)
	}
	return res
}
