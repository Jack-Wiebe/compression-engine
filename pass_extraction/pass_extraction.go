package pass_extraction

import (
	"compression-engine/util"
	"fmt"
)
type Pixel = util.Pixel
type Color = util.Color

//var pass_map_x map[Color][][]int // {color: [[pass_count,y],[x,y],[x,y]]}
// {color: [[pass_count,y],[x,y],[x,y]]}
//
//	{
//		pass_count: [pixel, pixel],
//		pass_count: [[x,y], [x,y]]
//	}
//var pass_template map[int][]Pos
//var pass_map_y map[Color][][]int

//Data struct for pass maps
//var pass_map_x map[Color]map[int]Pixel //color -> pass -> pixel
//var pass_map_y map[Color]map[int]Pixel

//load Image
func Test() {

	image,_ := util.LoadImage("test/test_image.png")

	pass_map_x := make(map[Color]map[int]Pixel)

	for x, line := range image{
		for y, c := range line{

			fmt.Println(c)
			pixel := Pixel{X:x, Y:y}

			if pass_map_x[c] == nil {
					pass_map_x[c] = make(map[int]Pixel)
			}
			pass_map_x[c][x] = pixel

		}
	}

	fmt.Println(pass_map_x)

}


// scan both dirextions
// group pixels by color value as key in pass map





