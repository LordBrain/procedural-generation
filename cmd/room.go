package buildings

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

//RoomStruct for building rooms
type RoomStruct struct {
	X1     int
	X2     int
	Y1     int
	Y2     int
	Width  int
	Height int
}

//LayoutStruct for holding a bunch of rooms
type LayoutStruct struct {
	Room   []RoomStruct
	Blocks [][]int
	Edges  RoomStruct
}
type Point struct {
	X, Y int
}

func NewTown(mapSize, roomSize, buildings int) *LayoutStruct {
	rand.Seed(time.Now().UnixNano())
	block := make([][]int, mapSize)
	for i := 0; i < mapSize; i++ {
		block[i] = make([]int, mapSize)
	}
	layout := &LayoutStruct{
		Room:   []RoomStruct{},
		Blocks: block,
		Edges:  RoomStruct{X1: 1, Y1: 1, X2: mapSize, Y2: mapSize},
	}
	for i := 1; i <= buildings; i++ {
		layout.AddRoom(roomSize, mapSize)
	}

	return layout
}

//Intersects Rooms. Make sure they do not over lap.
func (room1 *RoomStruct) Intersects(room2 *RoomStruct) bool {
	var x, y bool
	if (room2.X1 >= room1.X1) && (room2.X1 >= room1.X2) {
		x = true
	}
	if (room2.Y1 >= room1.Y1) && (room2.Y1 >= room1.Y2) {
		y = true
	}
	if (x) && (y) {
		fmt.Println("BUUUUUUUUUUUMP")
		return true
	} else {
		return false
	}
}

//AddRoom returns the x/y cords and the center
func (a *LayoutStruct) AddRoom(roomSize, mapSize int) {
	//Build a room I think.

	//Values Grid
	x := randomInt(1, mapSize)
	y := randomInt(1, mapSize)
	width := randomInt(4, roomSize)
	height := randomInt(4, roomSize)
	var x1, x2, y1, y2 int

	//Build the walls
	x1 = x
	x2 = x + width
	y1 = y
	y2 = y + height
	newRoom := RoomStruct{X1: x1, X2: x2, Y1: y1, Y2: y2, Width: width, Height: height}

	doIntersect := false
	for _, existingRoom := range a.Room {
		if (&newRoom).Intersects(&existingRoom) {
			doIntersect = true
			break
		}
	}
	if !doIntersect {
		a.Room = append(a.Room, newRoom)

	}

}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//PrintRoom displays the rooms made.
func PrintRoom(r *LayoutStruct) {
	numRooms := len(r.Room)

	for i := 0; i < numRooms; i++ {

		topWall := r.Room[i].X2 - r.Room[i].X1
		sideGap := topWall - 2
		sideLenght := r.Room[i].Y2 - r.Room[i].Y1
		topGap := sideLenght - 2

		for i := 0; i <= topWall; i++ {
			fmt.Print("#")
		}
		fmt.Println()
		for x := 0; x <= topGap; x++ {
			fmt.Print("#")
			for y := 0; y <= sideGap; y++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
			fmt.Println()
		}
		for i := 0; i <= topWall; i++ {
			fmt.Print("#")
		}
		// fmt.Printf("\n%v\n", r.Room[i])
		fmt.Println()
	}

}

func (l *LayoutStruct) CreateImage() {
	width := 200
	height := 200

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}
	numRooms := len(l.Room)

	for i := 0; i < numRooms; i++ {
		// Set color for each pixel.
		for x := l.Room[i].X1; x < l.Room[i].X2; x++ {
			for y := l.Room[i].Y1; y < l.Room[i].Y2; y++ {

				img.Set(x, y, cyan)

			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
