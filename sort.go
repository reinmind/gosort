package gosort

//Point has X,Y,Z
type Point struct {
	X int32
	Y int32
	Z int32
}

//PointArray array of Points
type PointArray []Point

//Len length of PointArray
func (pa PointArray) Len() int {
	return len(pa)
}

//Less compare two Point
func (pa PointArray) Less(i, j int) bool {
	if pa[i].X < pa[j].X {
		return true
	} else if pa[i].X == pa[j].X && pa[i].Y < pa[j].Y {
		return true
	} else if pa[i].X == pa[j].Y && pa[i].Y == pa[j].Y && pa[i].Z < pa[j].Z {
		return true
	}
	return false
}

//Swap swap two Point location
func (pa PointArray) Swap(i, j int) {
	pa[i], pa[j] = pa[j], pa[i]
}
