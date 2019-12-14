package main

import (
	"fmt"
)

type xyz struct {
	x, y, z int
}

type moon struct {
	pos      xyz
	velocity xyz
}

func main() {
	moons := [4]moon{
		moon{pos: xyz{16, -11, 2}, velocity: xyz{0, 0, 0}},
		moon{pos: xyz{0, -4, 7}, velocity: xyz{0, 0, 0}},
		moon{pos: xyz{6, 4, -10}, velocity: xyz{0, 0, 0}},
		moon{pos: xyz{-3, -2, -4}, velocity: xyz{0, 0, 0}},
	}

	xstates := make(map[string]bool)
	ystates := make(map[string]bool)
	zstates := make(map[string]bool)

	xperoid := -1
	yperoid := -1
	zperoid := -1

	
	step := 0
	for {
		if xperoid != -1 && yperoid != -1 && zperoid != -1 {
			fmt.Println("Found peroid")
			break
		}

		// Mutates gravity
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				p1, p2 := moons[i].pos, moons[j].pos
				v1, v2 := moons[i].velocity, moons[j].velocity
				
				v1x, v2x := getIncrements(p1.x, p2.x)
				v1y, v2y := getIncrements(p1.y, p2.y)
				v1z, v2z := getIncrements(p1.z, p2.z)
				
				v1.x += v1x
				v1.y += v1y
				v1.z += v1z
				v2.x += v2x
				v2.y += v2y
				v2.z += v2z
				
				moons[i].velocity = v1
				moons[j].velocity = v2
			}
		}
		
		for k, moon := range moons {
			moon.pos.x += moon.velocity.x
			moon.pos.y += moon.velocity.y
			moon.pos.z += moon.velocity.z
			
			moons[k] = moon
		}
		
		ok := false

		if xperoid == -1 {
			s := getStateString([]int{moons[0].pos.x, moons[1].pos.x, moons[2].pos.x, moons[3].pos.x, moons[0].velocity.x, moons[1].velocity.x, moons[2].velocity.x, moons[3].velocity.x})
			_, ok = xstates[s]
			if ok {
				fmt.Printf("Found peroid x at step %d\n", step)
				xperoid = step
			}
			xstates[s] = true
		}
		
		if yperoid == -1 {
			s := getStateString([]int{moons[0].pos.y, moons[1].pos.y, moons[2].pos.y, moons[3].pos.y, moons[0].velocity.y, moons[1].velocity.y, moons[2].velocity.y, moons[3].velocity.y})
			_, ok = ystates[s]
			if ok {
				fmt.Printf("Found peroid y at step %d\n", step)
				yperoid = step
			}
			ystates[s] = true
		}
		
		if zperoid == -1 {
			s := getStateString([]int{moons[0].pos.z, moons[1].pos.z, moons[2].pos.z, moons[3].pos.z, moons[0].velocity.z, moons[1].velocity.z, moons[2].velocity.z, moons[3].velocity.z})
			_, ok = zstates[s]
			if ok {
				fmt.Printf("Found peroid z step %d\n", step)
				zperoid = step
			}
			zstates[s] = true
		}
		step++
	}

	fmt.Println(lcm(lcm(xperoid, yperoid), zperoid))

	// sum := 0
	// for _, moon := range moons {
	// 	sum += totalEnergy(moon)
	// }
	// fmt.Println(sum)
}

func lcm(a, b int) int {
	return (a * b) / gcf(a, b)
}

func gcf(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return a
}

func getStateString(v []int) string {
	return fmt.Sprintf("%d%d%d%d%d%d%d%d", v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])
}

func totalEnergy(moon moon) int {
	pe := abs(moon.pos.x) + abs(moon.pos.y) + abs(moon.pos.z)
	ke := abs(moon.velocity.x) + abs(moon.velocity.y) + abs(moon.velocity.z)

	return pe * ke
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getIncrements(x, y int) (int, int) {
	if x < y {
		return 1, -1
	} else if x > y {
		return -1, 1
	}
	return 0, 0
}

func printMoons(moons [4]moon, step int) {
	fmt.Printf("After %d steps:\n", step)
	for _, moon := range moons {
		x1, y1, z1 := moon.pos.x, moon.pos.y, moon.pos.z
		x2, y2, z2 := moon.velocity.x, moon.velocity.y, moon.velocity.z

		fmt.Printf("pos=<x=%4d, y=%4d, z=%4d>, vel=<x=%4d, y=%4d, z=%4d>\n", x1, y1, z1, x2, y2, z2)
	}
	fmt.Println()
}
