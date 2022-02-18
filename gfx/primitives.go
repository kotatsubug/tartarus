package gfx

import (
	"tartarus.xyz/qt"
)

const BYTES_IN_A_VERTEX int = 4*6
type Vert struct {
	position qt.Vec3
	color qt.Vec3
}

type Primitive struct {
	verts []Vert
	indices []uint32
}

func Cube() *Primitive {
	return &Primitive{verts: cubeVertices, indices: nil}
}

var cubeVertices = []Vert{
	//  X, Y, Z, U, V
	// Bottom
	Vert{qt.Vec3{-1.0, -1.0, -1.0},  qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, 1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, 1.0},   qt.Vec3{1.0, 0.0, 0.0}},

	// Top
	Vert{qt.Vec3{-1.0, 1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, 1.0},     qt.Vec3{1.0, 0.0, 0.0}},

	// Front
	Vert{qt.Vec3{-1.0, -1.0, 1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, 1.0},     qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},

	// Back
	Vert{qt.Vec3{-1.0, -1.0, -1.0},  qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0},    qt.Vec3{1.0, 0.0, 0.0}},

	// Left
	Vert{qt.Vec3{-1.0, -1.0, 1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, -1.0},  qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, 1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},

	// Right
	Vert{qt.Vec3{1.0, -1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0},   qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0},    qt.Vec3{1.0, 0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, 1.0},     qt.Vec3{1.0, 0.0, 0.0}},
}

/*
var cubeVertices = []Vert{
	//  X, Y, Z, U, V
	// Bottom
	Vert{qt.Vec3{-1.0, -1.0, -1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, 1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0}, qt.Vec2{1.0, 1.0}},
	Vert{qt.Vec3{-1.0, -1.0, 1.0}, qt.Vec2{0.0, 1.0}},

	// Top
	Vert{qt.Vec3{-1.0, 1.0, -1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{1.0, 1.0, 1.0}, qt.Vec2{1.0, 1.0}},

	// Front
	Vert{qt.Vec3{-1.0, -1.0, 1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0}, qt.Vec2{1.0, 1.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, 1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0}, qt.Vec2{1.0, 1.0}},

	// Back
	Vert{qt.Vec3{-1.0, -1.0, -1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0}, qt.Vec2{1.0, 1.0}},

	// Left
	Vert{qt.Vec3{-1.0, -1.0, 1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, -1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{-1.0, -1.0, 1.0}, qt.Vec2{0.0, 1.0}},
	Vert{qt.Vec3{-1.0, 1.0, 1.0}, qt.Vec2{1.0, 1.0}},
	Vert{qt.Vec3{-1.0, 1.0, -1.0}, qt.Vec2{1.0, 0.0}},

	// Right
	Vert{qt.Vec3{1.0, -1.0, 1.0}, qt.Vec2{1.0, 1.0}},
	Vert{qt.Vec3{1.0, -1.0, -1.0}, qt.Vec2{1.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{1.0, -1.0, 1.0}, qt.Vec2{1.0, 1.0}},
	Vert{qt.Vec3{1.0, 1.0, -1.0}, qt.Vec2{0.0, 0.0}},
	Vert{qt.Vec3{1.0, 1.0, 1.0}, qt.Vec2{0.0, 1.0}},
}*/