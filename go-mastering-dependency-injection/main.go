package main

import "fmt"

type SaftyPlacer interface {
	PlaceSafeties()
}

// ICE
// Sandy rocks
// Concrete
type RockClimber struct {
	rocksClimbed int
	kind         int
	sp           SaftyPlacer
}

func newRockClimber(sp SaftyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct {
	// db
	// data
	// api
}

func (sp *IceSafetyPlacer) PlaceSafeties() {
	fmt.Println("placing my ICE safeties...")
}

type NOPSafetyPlacer struct{}

func (sp *NOPSafetyPlacer) PlaceSafeties() {
	fmt.Println("placing my NOP safeties...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.PlaceSafeties()
	}
}

func main() {
	rc := newRockClimber(&IceSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
