package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x        int
	y        int
	distance int
}

type Line struct {
	p1        Point
	p2        Point
	direction string // vertical or horizontal
	positive  bool   // line moves in positive direct or negative
}

func getLine(p1 Point, p2 Point) Line {
	var newLine Line
	newLine.positive = true
	if p1.x == p2.x {
		// vertical line
		newLine.direction = "vertical"
		if p2.y >= p1.y {
			newLine.p1 = p1
			newLine.p2 = p2
		} else {
			newLine.p1 = p2
			newLine.p2 = p1
			newLine.positive = false
		}
	} else {
		// horizontal line
		newLine.direction = "horizontal"
		if p2.x >= p1.x {
			newLine.p1 = p1
			newLine.p2 = p2
		} else {
			newLine.p1 = p2
			newLine.p2 = p1
			newLine.positive = false
		}
	}
	return newLine
}

func getPaths(moves []string) []Line {
	// start at [0,0] then map through (range)
	var allPaths []Line
	var startPoint Point
	startPoint.x = 0
	startPoint.y = 0
	for _, move := range moves {
		var byteArr = ([]byte)(move)
		// get position of to move
		pos, err := strconv.Atoi(string(byteArr[1:len(byteArr)]))
		if err != nil {
			panic(err)
		}
		var direction string = string([]byte{move[0]})
		var endPoint Point
		var newLine Line
		if direction == "R" {
			endPoint.x = startPoint.x + pos
			endPoint.y = startPoint.y
			newLine = getLine(startPoint, endPoint)
		} else if direction == "L" {
			endPoint.x = startPoint.x - pos
			endPoint.y = startPoint.y
			newLine = getLine(startPoint, endPoint)
		} else if direction == "U" {
			endPoint.x = startPoint.x
			endPoint.y = startPoint.y + pos
			newLine = getLine(startPoint, endPoint)
		} else if direction == "D" {
			endPoint.x = startPoint.x
			endPoint.y = startPoint.y - pos
			newLine = getLine(startPoint, endPoint)
		} else {
			panic("Bad Direction")
		}
		allPaths = append(allPaths, newLine)
		startPoint = endPoint
		// fmt.Println(newLine, startPoint)
	}
	return allPaths
}

func getLineIntersect(line1 Line, line2 Line) Point {
	// Starting intersect is just starting point (no intersect)
	var intersect Point
	intersect.x = 0
	intersect.y = 0

	if line1.direction != line2.direction {
		var horizontal Line
		var vertical Line
		if line1.direction == "horizontal" {
			horizontal = line1
			vertical = line2
		} else {
			horizontal = line2
			vertical = line1
		}
		if (horizontal.p1.y < vertical.p2.y && horizontal.p1.y > vertical.p1.y) &&
			(vertical.p1.x < horizontal.p2.x && vertical.p1.x > horizontal.p1.x) ||
			(horizontal.p1.x == vertical.p1.x || horizontal.p2.x == vertical.p1.x) ||
			(vertical.p1.y == horizontal.p1.y || vertical.p2.y == horizontal.p1.y) {
			intersect.x = vertical.p1.x
			intersect.y = horizontal.p1.y
			// fmt.Println(`intersection`, line1, line2)
		}
	}
	return intersect
}

func getWireIntersections(wire1 []Line, wire2 []Line) []Point {
	var allIntersect []Point
	for _, wire1Path := range wire1 {
		for _, wire2Path := range wire2 {
			var newIntersect Point = getLineIntersect(wire1Path, wire2Path)
			if newIntersect.x != 0 || newIntersect.y != 0 {
				// fmt.Println(`intersection`, newIntersect)
				allIntersect = append(allIntersect, newIntersect)
			}
		}
	}
	return allIntersect
}

func getClosestIntersect(intersections []Point) Point {
	// fmt.Println(intersections)
	var closer Point = intersections[0]
	for _, intersect := range intersections {
		var closerDist float64 = math.Abs(float64(closer.x)) + math.Abs(float64(closer.y))
		var intersectDist float64 = math.Abs(float64(intersect.x)) + math.Abs(float64(intersect.y))
		if intersectDist < closerDist {
			closer = intersect
		}
	}
	return closer
}

func getPointDistance(distance int, point Point, points []Point) int {
	var newDistance int = distance + 1
	for _, checkPoint := range points {
		if checkPoint.x == point.x && checkPoint.y == point.y {
			newDistance = checkPoint.distance
			break
		}
	}
	return newDistance
}

func mapWirePoints(wire []Line) []Point {
	var path []Point = []Point{}
	var distance int = 0
	for _, wireSegment := range wire {
		var startPoint Point
		var endPoint Point
		if wireSegment.positive {
			startPoint = wireSegment.p1
			endPoint = wireSegment.p2
			if wireSegment.direction == "horizontal" {
				var y int = startPoint.y // always the same if horizontal
				for i := startPoint.x + 1; i <= endPoint.x; i++ {
					var x int = i
					var currentPoint Point
					currentPoint.x = x
					currentPoint.y = y
					distance = getPointDistance(distance, currentPoint, path)
					currentPoint.distance = distance
					path = append(path, currentPoint)
				}
			} else {
				var x int = startPoint.x // always the same if vertical
				for i := startPoint.y + 1; i <= endPoint.y; i++ {
					var y int = i
					var currentPoint Point
					currentPoint.x = x
					currentPoint.y = y
					distance = getPointDistance(distance, currentPoint, path)
					currentPoint.distance = distance
					path = append(path, currentPoint)
				}
			}
		} else {
			// negative direction
			startPoint = wireSegment.p2
			endPoint = wireSegment.p1
			if wireSegment.direction == "horizontal" {
				var y int = startPoint.y // always the same if horizontal
				for i := startPoint.x - 1; i >= endPoint.x; i-- {
					var x int = i
					var currentPoint Point
					currentPoint.x = x
					currentPoint.y = y
					distance = getPointDistance(distance, currentPoint, path)
					currentPoint.distance = distance
					path = append(path, currentPoint)
				}
			} else {
				var x int = startPoint.x // always the same if vertical
				for i := startPoint.y - 1; i >= endPoint.y; i-- {
					var y int = i
					var currentPoint Point
					currentPoint.x = x
					currentPoint.y = y
					distance = getPointDistance(distance, currentPoint, path)
					currentPoint.distance = distance
					path = append(path, currentPoint)
				}
			}
		}

	}
	return path
}

func getPoints(point Point, points []Point) []Point {
	var newDistance int = 1
	if len(points) > 0 {
		newDistance = points[len(points)-1].distance + 1
	}
	for _, checkPoint := range points {
		if checkPoint.x == point.x && checkPoint.y == point.y {
			newDistance = checkPoint.distance
			break
		}
	}
	point.distance = newDistance
	points = append(points, point)
	return points
}

func getWirePoints(moves []string) []Point {
	// start at [0,0] then map through (range)
	var points []Point
	var startPoint Point
	startPoint.x = 0
	startPoint.y = 0

	for _, move := range moves {
		var byteArr = ([]byte)(move)
		// get position of to move
		pos, err := strconv.Atoi(string(byteArr[1:len(byteArr)]))
		if err != nil {
			panic(err)
		}
		var direction string = string([]byte{move[0]})
		// fmt.Println("d:", direction, pos)
		var lastPoint Point = startPoint
		for i := 1; i <= pos; i++ {
			var nextPoint Point
			if direction == "R" {
				nextPoint.x = lastPoint.x + 1
				nextPoint.y = lastPoint.y
			} else if direction == "L" {
				nextPoint.x = lastPoint.x - 1
				nextPoint.y = lastPoint.y
			} else if direction == "U" {
				nextPoint.x = lastPoint.x
				nextPoint.y = lastPoint.y + 1
			} else if direction == "D" {
				nextPoint.x = lastPoint.x
				nextPoint.y = lastPoint.y - 1
			} else {
				panic("Bad Direction")
			}
			points = getPoints(nextPoint, points)
			lastPoint = points[len(points)-1]
		}
		startPoint = lastPoint
	}
	return points
}

func removeDuplicates(points []Point) []Point {
	var newPoints []Point
	for i := 0; i < len(points); i++ {
		var notfound bool = true
		for _, point := range newPoints {
			if point.x == points[i].x && point.y == points[i].y {
				notfound = false
				break
			}
		}
		if notfound {
			newPoints = append(newPoints, points[i])
		}
	}
	return newPoints
}

func getWirePointIntersects(points1 []Point, points2 []Point) []Point {
	var intersects []Point
	for _, point1 := range points1 {
		for _, point2 := range points2 {
			if point1.x == point2.x && point1.y == point2.y {
				var intersect Point = point1
				intersect.distance = point1.distance + point2.distance
				intersects = append(intersects, intersect)
				break // don't add duplicates of duplicates
			}
		}
	}
	return intersects
}

func main() {

	var Wire1 []string = strings.Split("R997,D99,R514,D639,L438,D381,L251,U78,L442,D860,R271,U440,L428,U482,R526,U495,R361,D103,R610,D64,L978,U587,L426,D614,R497,D116,R252,U235,R275,D882,L480,D859,L598,D751,R588,D281,R118,U173,L619,D747,R994,U720,L182,U952,L49,D969,R34,D190,L974,U153,L821,U593,L571,U111,L134,U111,R128,D924,R189,U811,R100,D482,L708,D717,L844,U695,R277,D81,L107,U831,L77,U609,L629,D953,R491,D17,R160,U468,R519,D41,R625,D501,R106,D500,R473,D244,R471,U252,R440,U326,R710,D645,L190,D670,L624,D37,L46,D242,L513,D179,R192,D100,R637,U622,R322,U548,L192,D85,L319,D717,L254,D742,L756,D624,L291,D663,R994,U875,R237,U304,R40,D399,R407,D124,R157,D415,L405,U560,R607,U391,R409,U233,R305,U346,L233,U661,R213,D56,L558,U386,R830,D23,L75,D947,L511,D41,R927,U856,L229,D20,L717,D830,R584,U485,R536,U531,R946,U942,R207,D237,L762,U333,L979,U29,R635,D386,R267,D260,R484,U887,R568,D451,R149,U92,L379,D170,R135,U799,L617,D380,L872,U868,R48,U279,R817,U572,L728,D792,R833,U788,L940,D306,R230,D570,L137,U419,L429,D525,L730,U333,L76,D435,R885,U811,L937,D320,R152,U906,L461,U227,L118,U951,R912,D765,L638,U856,L193,D615,L347,U303,R317,U23,L139,U6,L525,U308,L624,U998,R753,D901,R556,U428,L224,U953,R804,D632,L764,U808,L487,U110,L593,D747,L659,D966,R988,U217,L657,U615,L425,D626,L194,D802,L440,U209,L28,U110,L564,D47,R698,D938,R13,U39,R703,D866,L422,D855,R535,D964,L813,D405,R116,U762,R974,U568,R934,U574,R462,D968,R331,U298,R994,U895,L204,D329,R982,D83,L301,D197,L36,U329,R144,U497,R300,D551,L74,U737,R591,U374,R815,U771,L681", ",")
	var Wire2 []string = strings.Split("L997,D154,R652,U379,L739,U698,R596,D862,L125,D181,R786,U114,R536,U936,L144,U936,R52,U899,R88,D263,R122,D987,L488,U303,R142,D556,L691,D769,L717,D445,R802,U294,L468,D13,R301,D651,L242,D767,R465,D360,L144,D236,R59,U815,R598,U375,R645,U905,L714,U440,R932,D160,L420,U361,L433,D485,L276,U458,R760,D895,R999,U263,R530,U691,L918,D790,L150,U574,R800,U163,R478,U112,L353,U30,L763,U239,L353,U619,R669,D822,R688,U484,L678,D88,R946,D371,L209,D175,R771,D85,R430,U16,R610,D326,R836,U638,L387,D996,L758,U237,L476,U572,L456,U579,L457,D277,L825,U204,R277,U267,L477,D573,L659,D163,L516,D783,R762,U146,L387,U700,R911,U335,L115,D887,R677,U312,R707,U463,L743,U358,L715,D603,R966,U21,L857,D680,R182,D977,L279,U196,R355,D624,L434,U410,R385,U47,L999,D542,L453,D735,R781,U115,R814,U110,R344,D139,R899,D650,L118,D774,L227,D140,L198,D478,R115,D863,R776,D935,R473,U722,R555,U528,L912,U268,R776,D223,L302,D878,R90,U52,L595,U898,L210,U886,R161,D794,L846,U404,R323,U616,R559,U510,R116,D740,L554,U231,R54,D328,L56,U750,R347,U376,L148,U454,L577,U61,L772,D862,R293,U82,L676,D508,L53,D860,L974,U733,R266,D323,L75,U218,L390,U757,L32,D455,R34,D363,L336,D67,R222,D977,L809,D909,L501,U483,L541,U923,R97,D437,L296,D941,L652,D144,L183,U369,L629,U535,L825,D26,R916,U131,R753,U383,L653,U631,R280,U500,L516,U959,R858,D830,R357,D87,R885,D389,L838,U550,R262,D529,R34,U20,L25,D553,L884,U806,L800,D988,R499,D360,R435,U381,R920,D691,R373,U714,L797,D677,L490,D976,L734,D585,L384,D352,R54,D23,R339,D439,L939,U104,L651,D927,L152", ",")

	// Wire1 = strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	// Wire2 = strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")

	// Wire1 = strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ",")
	// Wire2 = strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ",")

	// Wire1 = strings.Split("R8,U5,L5,D3", ",")
	// Wire2 = strings.Split("U7,R6,D4,L4", ",")

	// Wire1 = strings.Split("R8,U5,L5,D4", ",")
	// Wire2 = strings.Split("U7,R6,D4,L7,D2,R4", ",")

	var wire1Points = getWirePoints(Wire1)
	var wire2Points = getWirePoints(Wire2)
	// fmt.Println("points1: ", wire1Points)
	// fmt.Println("points2: ", wire2Points)

	var intersections []Point = getWirePointIntersects(removeDuplicates(wire1Points), removeDuplicates(wire2Points))
	// fmt.Println("intersections:", intersections)
	var smallestSteps int = intersections[0].distance
	for _, intersect := range intersections {
		if intersect.distance < smallestSteps {
			smallestSteps = intersect.distance
		}
	}
	fmt.Println("smallest steps:", smallestSteps)

	// fmt.Println("Manhattan Distance")
	// var wire1Paths []Line = getPaths(Wire1)
	// var wire2Paths []Line = getPaths(Wire2)

	// var wire1Points []Point = mapWirePoints(wire1Paths)
	// // fmt.Println("wire1Points", wire1Points)
	// var wire2Points []Point = mapWirePoints(wire2Paths)
	// // fmt.Println("wire2Points", wire2Points)
	// var intersections = getWireIntersections(wire1Paths, wire2Paths)
	// // fmt.Println("Intersections", intersections)

	// var smallestSteps int = 99999999999999
	// for _, intersect := range intersections {
	// 	var distance int = 0
	// 	var dist1 int = getPointDistance(distance, intersect, wire1Points)
	// 	var dist2 int = getPointDistance(distance, intersect, wire2Points)
	// 	var intesectionSteps int = dist1 + dist2
	// 	if intesectionSteps < smallestSteps {
	// 		smallestSteps = intesectionSteps
	// 	}
	// 	fmt.Println("steps:", intersect, intesectionSteps)
	// }
	// fmt.Println("smallest steps:", smallestSteps)

	// var closest = getClosestIntersect(intersections)
	// fmt.Println("Closest", closest)
	// fmt.Println("Answer", math.Abs(float64(closest.x))+math.Abs(float64(closest.y)))

}
