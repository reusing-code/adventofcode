package main

import "fmt"

type cube struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
	on     bool
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func puzzle1(input []string) int {
	regionQueue := make([]cube, 0)
	for _, str := range input {
		var c cube
		var cmd string
		_, err := fmt.Sscanf(str, "%v x=%v..%v,y=%v..%v,z=%v..%v", &cmd, &c.x1, &c.x2, &c.y1, &c.y2, &c.z1, &c.z2)
		if err != nil {
			panic(err)
		}
		if cmd == "on" {
			c.on = true
		} else {
			c.on = false
		}
		regionQueue = append(regionQueue, c)
	}

	onRegions := make([]cube, 0)
	for len(regionQueue) > 0 {
		curregion := regionQueue[0]
		regionQueue = regionQueue[1:]
		overlap := false
		for i, onRegion := range onRegions {
			if curregion.x1 >= onRegion.x1 && curregion.x2 <= onRegion.x2 &&
				curregion.y1 >= onRegion.y1 && curregion.y2 <= onRegion.y2 &&
				curregion.z1 >= onRegion.z1 && curregion.z2 <= onRegion.z2 {
				// completely enclosed in onRegion, ignore if on
				if curregion.on {
					overlap = true
					break
				} else {
					if curregion.x1 == onRegion.x1 && curregion.x2 == onRegion.x2 &&
						curregion.y1 == onRegion.y1 && curregion.y2 == onRegion.y2 &&
						curregion.z1 == onRegion.z1 && curregion.z2 == onRegion.z2 {
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
					} else if onRegion.x2 > curregion.x2 && onRegion.x1 < curregion.x2 {
						newReg := onRegion
						onRegion.x2 = curregion.x2
						newReg.x1 = curregion.x2 + 1
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
						onRegions = append(onRegions, onRegion)
						onRegions = append(onRegions, newReg)
					} else if curregion.x1 > onRegion.x1 && curregion.x1 < onRegion.x2 {
						newReg := onRegion
						onRegion.x2 = curregion.x1
						newReg.x1 = curregion.x1 + 1
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
						onRegions = append(onRegions, onRegion)
						onRegions = append(onRegions, newReg)
					} else if onRegion.y2 > curregion.y2 && onRegion.y1 < curregion.y2 {
						newReg := onRegion
						onRegion.y2 = curregion.y2
						newReg.y1 = curregion.y2 + 1
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
						onRegions = append(onRegions, onRegion)
						onRegions = append(onRegions, newReg)
					} else if curregion.y1 > onRegion.y1 && curregion.y1 < onRegion.y2 {
						newReg := onRegion
						onRegion.y2 = curregion.y1
						newReg.y1 = curregion.y1 + 1
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
						onRegions = append(onRegions, onRegion)
						onRegions = append(onRegions, newReg)
					} else if onRegion.z2 > curregion.z2 && onRegion.z1 < curregion.z2 {
						newReg := onRegion
						onRegion.z2 = curregion.z2
						newReg.z1 = curregion.z2 + 1
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
						onRegions = append(onRegions, onRegion)
						onRegions = append(onRegions, newReg)
					} else if curregion.z1 > onRegion.z1 && curregion.z1 < onRegion.z2 {
						newReg := onRegion
						onRegion.z2 = curregion.z1
						newReg.z1 = curregion.z1 + 1
						onRegions = append(onRegions[:i], onRegions[i+1:]...)
						onRegions = append(onRegions, onRegion)
						onRegions = append(onRegions, newReg)
					} else {
						panic("should not happen")
					}
					overlap = true
					break
				}
			}
			if curregion.x1 > onRegion.x1 && curregion.x1 < onRegion.x2 && curregion.x2 > onRegion.x2 {
				// x1 overlap
				overlap = true
				newReg := curregion
				curregion.x2 = onRegion.x2
				newReg.x1 = onRegion.x2 + 1
				newQueue := make([]cube, 0, len(regionQueue)+2)
				newQueue = append(newQueue, curregion)
				newQueue = append(newQueue, newReg)
				newQueue = append(newQueue, regionQueue...)
				regionQueue = newQueue
			} else if curregion.x2 > onRegion.x1 && curregion.x2 < onRegion.x2 && curregion.x1 < onRegion.x1 {
				// x2 overlap
				overlap = true
				newReg := curregion
				curregion.x2 = onRegion.x1
				newReg.x1 = onRegion.x1 + 1
				newQueue := make([]cube, 0, len(regionQueue)+2)
				newQueue = append(newQueue, curregion)
				newQueue = append(newQueue, newReg)
				newQueue = append(newQueue, regionQueue...)
				regionQueue = newQueue
			} else if curregion.y1 > onRegion.y1 && curregion.y1 < onRegion.y2 && curregion.y2 > onRegion.y2 {
				// y1 overlap
				overlap = true
				newReg := curregion
				curregion.y2 = onRegion.y2
				newReg.y1 = onRegion.y2 + 1
				newQueue := make([]cube, 0, len(regionQueue)+2)
				newQueue = append(newQueue, curregion)
				newQueue = append(newQueue, newReg)
				newQueue = append(newQueue, regionQueue...)
				regionQueue = newQueue
			} else if curregion.y2 > onRegion.y1 && curregion.y2 < onRegion.y2 && curregion.y1 < onRegion.y1 {
				// y2 overlap
				overlap = true
				newReg := curregion
				curregion.y2 = onRegion.y1
				newReg.y1 = onRegion.y1 + 1
				newQueue := make([]cube, 0, len(regionQueue)+2)
				newQueue = append(newQueue, curregion)
				newQueue = append(newQueue, newReg)
				newQueue = append(newQueue, regionQueue...)
				regionQueue = newQueue
			} else if curregion.z1 > onRegion.z1 && curregion.z1 < onRegion.z2 && curregion.z2 > onRegion.z2 {
				// z1 overlap
				overlap = true
				newReg := curregion
				curregion.z2 = onRegion.z2
				newReg.z1 = onRegion.z2 + 1
				newQueue := make([]cube, 0, len(regionQueue)+2)
				newQueue = append(newQueue, curregion)
				newQueue = append(newQueue, newReg)
				newQueue = append(newQueue, regionQueue...)
				regionQueue = newQueue
			} else if curregion.z2 > onRegion.z1 && curregion.z2 < onRegion.z2 && curregion.z1 < onRegion.z1 {
				// z2 overlap
				overlap = true
				newReg := curregion
				curregion.z2 = onRegion.z1
				newReg.z1 = onRegion.z1 + 1
				newQueue := make([]cube, 0, len(regionQueue)+2)
				newQueue = append(newQueue, curregion)
				newQueue = append(newQueue, newReg)
				newQueue = append(newQueue, regionQueue...)
				regionQueue = newQueue
			}
		}
		if !overlap {
			if curregion.on {
				onRegions = append(onRegions, curregion)
			}
		}
		count := 0
		for _, region := range onRegions {
			volume := 1
			volume *= clamp(region.x2, -50, 50) - clamp(region.x1, -50, 50) + 1
			volume *= clamp(region.y2, -50, 50) - clamp(region.y1, -50, 50) + 1
			volume *= clamp(region.z2, -50, 50) - clamp(region.z1, -50, 50) + 1
			count += volume
		}
	}

	count := 0
	for _, region := range onRegions {
		volume := 1
		volume *= clamp(region.x2, -50, 50) - clamp(region.x1, -50, 50) + 1
		volume *= clamp(region.y2, -50, 50) - clamp(region.y1, -50, 50) + 1
		volume *= clamp(region.z2, -50, 50) - clamp(region.z1, -50, 50) + 1
		count += volume
	}

	return count
}

func puzzle2(input []string) int {
	return 1
}
