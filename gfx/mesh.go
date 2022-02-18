package gfx

import (
//	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"

	"tartarus.xyz/qt"
)

type Mesh struct {
	verts []Vert
	indices []uint32

	// Number of verts may not necessarily be len(verts)
	numVerts int
	numIndices int

	transform *Transform
	modelMatrix qt.Mat4
	vao, vbo, ebo uint32
}

func NewMeshFromPrimitive(prim *Primitive, tr *Transform) *Mesh {
	m := Mesh{
		verts: prim.verts,
		indices: nil,
		numVerts: len(prim.verts),
		numIndices: 0,
		transform: tr,
	}

	m.InitMeshBuffers()
	m.updateModelMatrix()
	return &m
}

// Removes mesh VAOs, VBOs, and EBOs from memory.
func (m *Mesh) Free() {
	gl.DeleteVertexArrays(1, &m.vao)
	gl.DeleteBuffers(1, &m.vbo)
	if (m.numIndices > 0) {
		gl.DeleteBuffers(1, &m.ebo)
	}
}

func (m *Mesh) InitMeshBuffers() {
	gl.GenVertexArrays(1, &m.vao)
	gl.GenBuffers(1, &m.vbo)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, m.numVerts * BYTES_IN_A_VERTEX, gl.Ptr(m.verts), gl.STATIC_DRAW)

	gl.BindVertexArray(m.vao)

	if (m.numIndices > 0) {
		gl.GenBuffers(1, &m.ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, m.numIndices * 4, gl.Ptr(m.indices), gl.STATIC_DRAW)
	}

	// Attributes
	gl.EnableVertexAttribArray(0) // Position
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)

	gl.EnableVertexAttribArray(1) // Color
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)

	// Unbind
	gl.BindVertexArray(0)
	gl.EnableVertexAttribArray(0)
}

func (m *Mesh) updateModelMatrix() {
	// Reset
	m.modelMatrix = qt.Identity4()
	// Scale
	m.modelMatrix = m.modelMatrix.SetScale(m.transform.Scale)
	// Move to its position in the world
	m.modelMatrix = m.modelMatrix.SetPosition(m.transform.Pos)
	// Rotate around mesh origin
	m.modelMatrix = m.modelMatrix.RotateAboutPivot(qt.AxisAngleToMat3(qt.Vec3{1,0,0}, m.transform.EulerRot[0]), m.transform.Origin.Sub(m.transform.Pos))
	m.modelMatrix = m.modelMatrix.RotateAboutPivot(qt.AxisAngleToMat3(qt.Vec3{0,1,0}, m.transform.EulerRot[1]), m.transform.Origin.Sub(m.transform.Pos))
	m.modelMatrix = m.modelMatrix.RotateAboutPivot(qt.AxisAngleToMat3(qt.Vec3{0,0,1}, m.transform.EulerRot[2]), m.transform.Origin.Sub(m.transform.Pos))
}

func (m *Mesh) updateUniforms(program uint32) {
	loc := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(loc, 1, false, &m.modelMatrix[0])
}

// "program" is the GLint shader ID to throw this mesh's uniforms into when rendering.
func (m *Mesh) Draw(program uint32) {
	m.updateModelMatrix()
	m.updateUniforms(program)
//	m.UpdateAnimations(program)

	gl.BindVertexArray(m.vao)
	if (m.numIndices > 0) {
		gl.DrawElements(gl.TRIANGLES, int32(m.numIndices), gl.UNSIGNED_INT, nil)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(m.numVerts))
	}
	gl.BindVertexArray(0)
}
