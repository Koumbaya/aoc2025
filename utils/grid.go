package utils

import (
	"fmt"
	"strings"
)

type Pos struct {
	Y int // rows
	X int // columns (yes I know)
}

type Grid struct {
	G          []string
	VisitedDir map[PosDir]int // to keep track of visits by direction.
	Visited    map[Pos]int    // to keep track of visits in any direction.
}

func NewGrid(in []string) *Grid {
	return &Grid{
		G:          in,
		VisitedDir: make(map[PosDir]int),
		Visited:    make(map[Pos]int),
	}
}

func (g *Grid) Print() {
	for y := 0; y < len(g.G); y++ {
		fmt.Println(g.G[y])
	}
	fmt.Println()
}

// Get return the value at position p. Returns empty string if out of bounds.
func (g *Grid) Get(p Pos) string {
	if g.Valid(p) {
		return string(g.G[p.Y][p.X])
	}
	return ""
}

// Set sets values s at position p. Not super efficient.
func (g *Grid) Set(p Pos, s string) {
	if len(s) > 1 {
		panic(s)
	}
	if g.Valid(p) {
		line := g.G[p.Y]
		g.G[p.Y] = line[:p.X] + s + line[p.X+1:]
	}
}

// CountChainsDirections starts from p and goes in every Directions given while advancing through the string.
// e.g. (Pos{Y:0, X:3}, dirs[E,W], "TOP") on a grid "POTOP" would return 2.
func (g *Grid) CountChainsDirections(p Pos, dirs []Direction, chain string) int {
	if !g.IsSymbol(p, chain[0]) {
		return 0
	}
	count := 0
	for _, dir := range dirs {
		count += g.CountChainsDirection(p, dir, chain[1:])
	}

	return count
}

func (g *Grid) CountChainsDirection(p Pos, d Direction, chain string) int {
	p = d.DirFunc(p)
	if !g.Valid(p) {
		return 0
	}

	if g.IsSymbol(p, chain[0]) {
		if len(chain) == 1 {
			return 1 // full chain
		}
		return g.CountChainsDirection(p, d, chain[1:]) // continue chain
	}
	return 0 // broken chain
}

// CountChainsParallel counts chains going in given direction and changing directions.
// e.g. for a grid:
// A B C
// B B C
// (Pos{Y:0, X:0}, all, "ABC") would return 4.
func (g *Grid) CountChainsParallel(p Pos, d []Direction, chain string) []Pos {
	if g.IsSymbol(p, chain[0]) {
		if len(chain) == 1 {
			return []Pos{p} // full chain
		} // continue chain
	} else {
		return []Pos{} // broken chain
	}

	nbs := g.GetPosFromDirs(p, d)
	c := make(chan []Pos)
	defer close(c)
	valid := 0
	ret := make([]Pos, 0)
	for _, n := range nbs {
		if g.Valid(n) {
			valid++
			go func(x Pos) {
				c <- g.CountChainsParallel(x, d, chain[1:])
			}(n)
		}
	}

	for range valid {
		ret = append(ret, <-c...)
	}

	return ret
}

func (g *Grid) SameYorX(p1, p2 Pos) bool {
	return p1.X == p2.X || p1.Y == p2.Y
}

func (g *Grid) Find(s string) Pos {
	for y := 0; y < len(g.G); y++ {
		for x := 0; x < len(g.G[0]); x++ {
			if string(g.G[y][x]) == s {
				return Pos{
					Y: y,
					X: x,
				}
			}
		}
	}
	return Pos{}
}

func (g *Grid) FindAll(val string) []Pos {
	res := make([]Pos, 0)
	for y := 0; y < len(g.G); y++ {
		for x := 0; x < len(g.G[0]); x++ {
			p := Pos{X: x, Y: y}
			s := g.Get(p)
			if s == val {
				res = append(res, p)
			}
		}
	}
	return res
}

func (g *Grid) CountOccurrencesColumn(x int, s string) int {
	r := 0
	for y := 0; y < len(g.G); y++ {
		if string(g.G[y][x]) == s {
			r++
		}
	}
	return r
}

func (g *Grid) CountOccurrencesRow(y int, s string) int {
	r := 0
	for x := 0; x < len(g.G[0]); x++ {
		if string(g.G[y][x]) == s {
			r++
		}
	}
	return r
}

func (g *Grid) CountOccurrences(s string) int {
	r := 0
	for y := 0; y < len(g.G); y++ {
		for x := 0; x < len(g.G[0]); x++ {
			if string(g.G[y][x]) == s {
				r++
			}
		}
	}
	return r
}

func (g *Grid) IsSymbol(p Pos, t uint8) bool {
	return g.G[p.Y][p.X] == t
}

func (g *Grid) IsOneOfSymbol(p Pos, symb string) bool {
	if !g.Valid(p) {
		return false
	}
	if strings.Contains(symb, string(g.G[p.Y][p.X])) {
		return true
	}
	return false
}

// Neighbors returns the immediate 8 neighbors of p
func (g *Grid) Neighbors(p Pos) []Pos {
	nbs := make([]Pos, 8)
	dirs := GetAllDirs()
	nbs[0] = dirs[NW].DirFunc(p)
	nbs[1] = dirs[N].DirFunc(p)
	nbs[2] = dirs[NE].DirFunc(p)
	nbs[3] = dirs[E].DirFunc(p)
	nbs[4] = dirs[SE].DirFunc(p)
	nbs[5] = dirs[S].DirFunc(p)
	nbs[6] = dirs[SW].DirFunc(p)
	nbs[7] = dirs[W].DirFunc(p)
	return nbs
}

type DirVal struct {
	DirName
	Val string
}

func (g *Grid) NeighborsVals(p Pos) map[DirName]string {
	dirs := GetAllDirs()
	vals := make(map[DirName]string)
	for _, d := range dirs {
		vals[d.DirName] = g.Get(d.DirFunc(p))
	}

	return vals
}

func (g *Grid) NeighborsWithVal(p Pos, val string) map[DirName]bool {
	nbsVal := g.NeighborsVals(p)
	vals := make(map[DirName]bool)
	for d, v := range nbsVal {
		if v == val {
			vals[d] = true
		}
	}
	return vals
}

func (g *Grid) GetPosFromDirs(p Pos, d []Direction) []Pos {
	r := make([]Pos, len(d))
	for i := range d {
		r[i] = d[i].DirFunc(p)
	}

	return r
}

// NeighborsIn return Neighbors if their value is one of the character in the s.
func (g *Grid) NeighborsIn(p Pos, s string) []Pos {
	nbs := g.Neighbors(p)
	out := make([]Pos, 0)
	for _, n := range nbs {
		if g.IsOneOfSymbol(n, s) {
			out = append(out, n)
		}
	}

	return out
}

// CountNeighborsIn count 1 for each Neighbors if their value is one of the character in the s.
func (g *Grid) CountNeighborsIn(p Pos, s string) int {
	count := 0
	nbs := g.Neighbors(p)
	for _, n := range nbs {
		if g.IsOneOfSymbol(n, s) {
			count++
		}
	}

	return count
}

// Valid returns false if p is out of bounds.
func (g *Grid) Valid(p Pos) bool {
	if p.X < 0 || p.Y < 0 || p.Y >= len(g.G) || p.X >= len((g.G)[0]) {
		return false
	}
	return true
}

// Visit store a position & direction as visited.
func (g *Grid) Visit(p PosDir) {
	g.VisitedDir[p]++
	g.Visited[p.Pos]++
}

// SetVisitedCount forces the visited count for a position & direction.
func (g *Grid) SetVisitedCount(p PosDir, count int) {
	g.VisitedDir[p] = count
	g.Visited[p.Pos] = count
}

// ClearVisit removes a position & direction from the list of visited tiles.
func (g *Grid) ClearVisit(p PosDir) {
	delete(g.VisitedDir, p)
	delete(g.Visited, p.Pos)
}

// IsVisited returns true if the position & direction is marked as visited.
func (g *Grid) IsVisited(p PosDir) bool {
	_, ok := g.VisitedDir[p]
	return ok
}

// IsVisitedAny returns true if the position is marked as visited.
func (g *Grid) IsVisitedAny(p Pos) bool {
	_, ok := g.Visited[p]
	return ok
}

// CountVisited count how many Pos were visited.
func (g *Grid) CountVisited() int {
	return len(g.Visited)
}

// CountVisitedTimes returns how many times a Pos has been visited in this direction.
func (g *Grid) CountVisitedTimes(pos PosDir) int {
	if i, ok := g.VisitedDir[pos]; !ok {
		return 0
	} else {
		return i
	}
}

// CountAllVisitedTimes return the number of tiles marked as visited * number of directions they were visited in.
func (g *Grid) CountAllVisitedTimes() int {
	r := 0
	for _, i := range g.VisitedDir {
		r += i
	}

	return r
}

// GetVisited return the Pos of all visited tiles, regardless of direction.
func (g *Grid) GetVisited() []Pos {
	r := make([]Pos, 0, len(g.Visited))
	for pos := range g.Visited {
		r = append(r, pos)
	}

	return r
}

func (g *Grid) ResetVisited() {
	g.VisitedDir = make(map[PosDir]int)
	g.Visited = make(map[Pos]int)
}

// WalkUntil steps through the grid in Direction d until one of the char in block is reached, or the edge is reached.
// It can return when it steps through a tile that's already been visited in this direction with blockVisited (indicating a loop).
// It marks Pos it can step on as VisitedDir.
func (g *Grid) WalkUntil(s Pos, d Direction, block string, blockVisited bool) (p Pos, edge bool, visit bool) {
	p = s
	for {
		s = d.DirFunc(s)
		if !g.Valid(s) {
			return p, true, false
		}
		if g.IsOneOfSymbol(s, block) {
			return p, false, false
		}
		if blockVisited && g.CountVisitedTimes(PosDir{Pos: s, DirName: d.DirName}) > 0 {
			return p, false, true
		}
		g.Visit(PosDir{Pos: s, DirName: d.DirName})
		p = s
	}
}

// Analyze return a map of items found on the grid and their positions.
// exclude - values to ignore (background)
func (g *Grid) Analyze(exclude string) map[string][]Pos {
	res := make(map[string][]Pos)
	for y := 0; y < len(g.G); y++ {
		for x := 0; x < len(g.G[0]); x++ {
			p := Pos{X: x, Y: y}
			s := g.Get(p)
			if strings.Contains(exclude, s) {
				continue
			}
			res[s] = append(res[s], p)
		}
	}
	return res
}

type Region struct {
	Name      string
	Count     int
	Perimeter int
	Corners   int
}

func (g *Grid) MapDistinctRegionsPerimeter(dirs []Direction) []Region {
	res := make([]Region, 0)
	for y := 0; y < len(g.G); y++ {
		for x := 0; x < len(g.G[0]); x++ {
			p := Pos{X: x, Y: y}
			v := g.Get(p)
			if !g.IsVisitedAny(p) {
				count, perimeter, corners := g.CountContinuousTiles(p, dirs, v)
				res = append(res, Region{Name: v, Perimeter: perimeter, Count: count, Corners: corners})
			}
		}
	}
	return res
}

func (g *Grid) CountContinuousTiles(start Pos, dirs []Direction, allowedVals string) (count, perimeter, corners int) {
	if !g.Valid(start) || !g.IsOneOfSymbol(start, allowedVals) {
		return 0, 0, 0
	}

	visitedPos := make(map[Pos]bool)
	cornersPos := make(map[Pos]struct{})
	visitedPos[start] = true
	count++
	queue := []Pos{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			nb := dir.DirFunc(curr)
			allowed := g.IsOneOfSymbol(nb, allowedVals)
			valid := g.Valid(nb)
			visited := visitedPos[nb]
			if !allowed {
				switch dir.DirName {
				case N, E, S, W:
					// count as perimeter, mark as visited to not count twice
					perimeter++
					visitedPos[nb] = true
				case NE, SE, SW, NW:
					if g.IsCorner(nb, allowedVals) {
						cornersPos[nb] = struct{}{}
					}
				}
			} else if valid && !visited {
				// mark visited on the map
				g.Visit(PosDir{Pos: Pos{
					Y: nb.Y,
					X: nb.X,
				}})
				// normal visit flow
				count++
				visitedPos[nb] = true
				queue = append(queue, nb)
			}
		}
	}

	return count, perimeter, len(cornersPos)
}

// uglyyyyyyy :(
// also does not work
func (g *Grid) IsCorner(p Pos, val string) bool {
	vs := g.NeighborsWithVal(p, val)
	switch len(vs) {
	case 8, 1, 4, 5, 6, 7: // 8 is inner patch
		return true
	case 0: // but why would it be zero ???
		return false
	case 2:
		switch {
		case vs[NE] && vs[E],
			vs[NW] && vs[N],
			vs[NE] && vs[N],
			vs[SW] && vs[W],
			vs[SW] && vs[S],
			vs[E] && vs[SW]:
			return false
		}
		return true
	case 3:
		switch {
		case vs[N] && vs[NE] && vs[E],
			vs[W] && vs[NW] && vs[N],
			vs[W] && vs[SW] && vs[S],
			vs[S] && vs[SE] && vs[E]:
			return true
		}
	}

	return false
}

// should only be used with a cardinal direction of 1 step.
func (g *Grid) TryMovePush(p Pos, dir Direction, blocks string, ignores string) bool {
	toMove := make([]Pos, 0)
	toMove = append(toMove, p)
	pp := p
	for {
		pp = dir.DirFunc(pp)
		if !g.Valid(pp) {
			panic("not implemented")
		}
		if g.IsOneOfSymbol(pp, blocks) {
			return false
		}
		toMove = append(toMove, pp)
		if g.IsOneOfSymbol(pp, ignores) {
			endSymbol := g.Get(pp)
			for i := len(toMove) - 1; i > 1; i-- {
				g.Set(toMove[i], g.Get(toMove[i-1]))
			}
			g.Set(toMove[0], endSymbol)
			return true
		}
	}
}

func (g *Grid) BFS(startPos Pos, targetPos Pos, dirs []Direction, walls string) []Pos {
	type Path struct {
		Pos   Pos
		Steps []Pos
	}

	visited := make(map[Pos]bool)
	queue := []Path{{Pos: startPos, Steps: []Pos{startPos}}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Pos == targetPos {
			return current.Steps
		}

		if visited[current.Pos] {
			continue
		}
		visited[current.Pos] = true

		for _, dir := range dirs {
			neighbor := dir.DirFunc(current.Pos)
			if g.Valid(neighbor) && !visited[neighbor] && !g.IsOneOfSymbol(neighbor, walls) {
				queue = append(queue, Path{
					Pos:   neighbor,
					Steps: append(current.Steps, neighbor),
				})
			}
		}
	}

	return nil
}

func (g *Grid) FindPathByLengthAndWeight(
	startPos Pos,
	endPos Pos,
	directions []Direction,
	maxLength int,
	maxWeight int,
	weights func(p Pos) int,
) []Pos {
	type Path struct {
		Pos    Pos
		Steps  []Pos
		Length int
		Weight int
	}

	visited := make(map[Pos]bool)
	queue := []Path{{Pos: startPos, Steps: []Pos{startPos}, Length: 0, Weight: 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.Pos] {
			continue
		}
		visited[current.Pos] = true

		if current.Pos == endPos && current.Length <= maxLength && current.Weight < maxWeight {
			return current.Steps
		}

		if current.Length >= maxLength {
			continue
		}

		// Enqueue valid neighbors
		for _, dir := range directions {
			neighbor := dir.DirFunc(current.Pos)
			if g.Valid(neighbor) {
				weight := weights(neighbor)
				newWeight := current.Weight + weight
				if newWeight < maxWeight {
					queue = append(queue, Path{
						Pos:    neighbor,
						Steps:  append(current.Steps, neighbor),
						Length: current.Length + 1,
						Weight: newWeight,
					})
				}
			}
		}
	}

	return nil
}

type DirName int

const (
	N DirName = iota
	NE
	E
	SE
	S
	SW
	W
	NW
	Custom = 99
)

type DirFunc func(Pos) Pos

type Direction struct {
	DirFunc
	DirName
}

type PosDir struct {
	Pos
	DirName
}

// DirectionFromPos transforms the path from Pos a to b into a Direction function.
func DirectionFromPos(a, b Pos) Direction {
	dx := b.X - a.X
	dy := b.Y - a.Y

	return Direction{
		DirFunc: func(p Pos) Pos {
			return Pos{X: p.X + dx, Y: p.Y + dy}
		},
		DirName: Custom,
	}
}

// Revert returns the inverse of the Direction operation.
func (d *Direction) Revert() Direction {
	// to keep consistent dirname return for standard direction
	switch d.DirName {
	case N, NE, E, SE, SW, W, NW:
		rotate180 := d.Rotate90()
		return rotate180.Rotate90()
	default:
	}

	r := d.DirFunc(Pos{X: 0, Y: 0})
	return Direction{
		DirFunc: func(p Pos) Pos {
			return Pos{X: p.X - r.X, Y: p.Y - r.Y}
		},
		DirName: Custom,
	}
}

// Rotate90 returns the next 90° position (clockwise).
func (d *Direction) Rotate90() Direction {
	switch d.DirName {
	case N:
		return GetAllDirs()[E]
	case NE:
		return GetAllDirs()[SE]
	case E:
		return GetAllDirs()[S]
	case SE:
		return GetAllDirs()[SW]
	case S:
		return GetAllDirs()[W]
	case SW:
		return GetAllDirs()[NW]
	case W:
		return GetAllDirs()[N]
	case NW:
		return GetAllDirs()[NE]
	default:
		panic("not implemented yet for Custom")
	}
}

// Rotate45 returns the next 45° position (clockwise).
func (d *Direction) Rotate45() Direction {
	switch d.DirName {
	case N:
		return GetAllDirs()[NE]
	case NE:
		return GetAllDirs()[E]
	case E:
		return GetAllDirs()[SE]
	case SE:
		return GetAllDirs()[S]
	case S:
		return GetAllDirs()[SW]
	case SW:
		return GetAllDirs()[W]
	case W:
		return GetAllDirs()[NW]
	case NW:
		return GetAllDirs()[N]
	default:
		panic("not implemented yet for Custom")
	}
}

// GetAllDirs return a map of the 8 basic Directions for easy access/reference.
// DirName is then a little bit redundant but we want to be able to switch on basic directions.
func GetAllDirs() map[DirName]Direction {
	return map[DirName]Direction{
		N:  {DirFunc: func(p Pos) Pos { return Pos{X: p.X, Y: p.Y - 1} }, DirName: N},
		NE: {DirFunc: func(p Pos) Pos { return Pos{X: p.X + 1, Y: p.Y - 1} }, DirName: NE},
		E:  {DirFunc: func(p Pos) Pos { return Pos{X: p.X + 1, Y: p.Y} }, DirName: E},
		SE: {DirFunc: func(p Pos) Pos { return Pos{X: p.X + 1, Y: p.Y + 1} }, DirName: SE},
		S:  {DirFunc: func(p Pos) Pos { return Pos{X: p.X, Y: p.Y + 1} }, DirName: S},
		SW: {DirFunc: func(p Pos) Pos { return Pos{X: p.X - 1, Y: p.Y + 1} }, DirName: SW},
		W:  {DirFunc: func(p Pos) Pos { return Pos{X: p.X - 1, Y: p.Y} }, DirName: W},
		NW: {DirFunc: func(p Pos) Pos { return Pos{X: p.X - 1, Y: p.Y - 1} }, DirName: NW},
	}
}

func GetOrdinals() map[DirName]Direction {
	return map[DirName]Direction{
		NW: {DirFunc: func(p Pos) Pos { return Pos{X: p.X - 1, Y: p.Y - 1} }, DirName: NW},
		SW: {DirFunc: func(p Pos) Pos { return Pos{X: p.X - 1, Y: p.Y + 1} }, DirName: SW},
		NE: {DirFunc: func(p Pos) Pos { return Pos{X: p.X + 1, Y: p.Y - 1} }, DirName: NE},
		SE: {DirFunc: func(p Pos) Pos { return Pos{X: p.X + 1, Y: p.Y + 1} }, DirName: SE},
	}
}

func GetCardinals() map[DirName]Direction {
	return map[DirName]Direction{
		N: {DirFunc: func(p Pos) Pos { return Pos{X: p.X, Y: p.Y - 1} }, DirName: N},
		S: {DirFunc: func(p Pos) Pos { return Pos{X: p.X, Y: p.Y + 1} }, DirName: S},
		W: {DirFunc: func(p Pos) Pos { return Pos{X: p.X - 1, Y: p.Y} }, DirName: W},
		E: {DirFunc: func(p Pos) Pos { return Pos{X: p.X + 1, Y: p.Y} }, DirName: E},
	}
}
