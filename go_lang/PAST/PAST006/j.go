package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"container/heap"
	"strconv"
	"math"
)



func main() {
	// fmt.Fscan(read,&hoge,&fuga,...)
	// fmt.Fprintln(write,hoge,fuga,...)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()
	// read := bufio.NewReader(os.Stdin)　// <-どちらか一方のみ 
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	c := float64(readInt())
	x := make([]float64,n)
	ans := float64(0)
	for i:=0;i<n;i++{
		x[i] = float64(readInt())
		y := float64(readInt())
		ans += (c-y)*(c-y)
	}
	low := -1e5-7
	up := 1e5+7
	xsum := 1e18
	for itr:=0;itr<555;itr++{
		mid1 := (up*2.0+low)/3.0
		mid2 := (up+low*2.0)/3.0
		sum1 := float64(0)
		sum2 := float64(0)
		for i:=0;i<n;i++{
			sum1 += (mid1-x[i])*(mid1-x[i])
			sum2 += (mid2-x[i])*(mid2-x[i])
		}
		xsum=math.Min(xsum,math.Min(sum1,sum2))
		if sum1>=sum2{
			up=mid1
		}else{
			low=mid2
		}
	}
	ans += xsum
	fmt.Fprintln(write,ans)
}


// io -------------------------
var sc = bufio.NewScanner(os.Stdin)
func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
// --------------------------------

// chmax(&a,b)として使う
func chmax(a *int, b int) bool {
	if *a>=b{
		return false
	}else{
		*a=b
		return true
	}
}

func chmin(a *int, b int) bool {
	if *a<=b{
		return false
	}else{
		*a=b
		return true
	}
}

// ソートされたkeyを返す。keyの型だけ指定する必要がある
func MapSortingString(mp interface{}) []string{
	keys := reflect.ValueOf(mp).MapKeys()
    res := make([]string, len(keys))
    for i, key := range keys {
        res[i] = key.String()
    }
    sort.Strings(res)
    return res
}
func MapSortingInt(mp interface{}) []int{
	keys := reflect.ValueOf(mp).MapKeys()
    res := make([]int, len(keys))
    for i, key := range keys {
        res[i] = int(key.Int())
    }
    sort.Ints(res)
    return res
}

type Array struct{
	a,b int
	// HOW TO SORT
	// x := make([]Array,n)
	// sort.Slice(x, func(i,j int) bool {
	// 	return x[i].a < x[j].a
	// })
}

type Edge struct {
	to, cost int
}
type WeightedGraph struct {
	n int
	v [][]Edge
}
func NewWeightedGraph(n int) WeightedGraph {
	g := WeightedGraph{n,make([][]Edge,n)}
	return g
}
func (g *WeightedGraph) add_directed_edge(from, to, cost int){
	g.v[from] = append(g.v[from],Edge{to,cost})
}
func (g *WeightedGraph) add_undirected_edge(u, v, cost int){
	g.v[u] = append(g.v[u],Edge{v,cost})
	g.v[v] = append(g.v[v],Edge{u,cost})
}
type State struct { dist, p int }
type Dijkstra []State
func (pq Dijkstra) Len() int { return len(pq) }
func (pq Dijkstra) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq Dijkstra) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *Dijkstra) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *Dijkstra) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}
func (g *WeightedGraph) shortest_path(s int)[]int{
	n := g.n
	INF := 100000000000000000+7
	dp := make([]int,n)
	for i:=0;i<n;i++ {
		dp[i] = INF
	}
	pq := new(Dijkstra)
	heap.Push(pq,State{0,s})
	for pq.Len() != 0 {
		cur := heap.Pop(pq).(State)
		if(dp[cur.p] < cur.dist){ continue }
		for i:=0;i<len(g.v[cur.p]);i++{
			nv := g.v[cur.p][i].to
			ndist := cur.dist + g.v[cur.p][i].cost
			if dp[nv] > ndist{
				dp[nv] = ndist
				heap.Push(pq,State{ndist,nv})
			}
		}
	}
	return dp
}
