package mapgen

type Coords struct {
	X int64
	Y int64
}

type DungeonMapPart struct {
	LeftTop     Coords
	RightBottom Coords
}

// func (dmp *DungeonMapPart) Width() int64 {
// 	return int64(uint64(dmp.LeftTop.X-dmp.RightBottom.X) &^ (1 << 63))
// }

// func (dmp *DungeonMapPart) Height() int64 {
// 	return int64(uint64(dmp.LeftTop.Y-dmp.RightBottom.Y) &^ (1 << 63))
// }
