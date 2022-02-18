package gfx

import (
	"tartarus.xyz/qt"
)

type Transform struct {
	Pos qt.Vec3
	Origin qt.Vec3
	EulerRot qt.Vec3 // TODO: Quaternion system
	Scale qt.Vec3
}

func NewTransform(pos, origin, eulerRot qt.Vec3) Transform {
	return Transform{Pos: pos, Origin: origin, EulerRot: eulerRot, Scale: qt.Vec3{1.0, 1.0, 1.0}}
}
