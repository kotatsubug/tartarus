package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
	
	"tartarus.xyz/qt"
	"tartarus.xyz/gfx"
	"tartarus.xyz/ecs"
)

const __fragmentShader string = `
#version 330

in vec3 vs_position;
in vec3 vs_color;

out vec4 fs_color;

void main() {
    fs_color = vec4(vs_color, 1.0);
}
` + "\x00"

const __vertexShader string = `
#version 330

layout (location = 0) in vec3 vertex_position;
layout (location = 1) in vec3 vertex_color;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

out vec3 vs_position;
out vec3 vs_color;

void main() {
    vs_position = vec4(model * vec4(vertex_position, 1.0)).xyz;
    vs_color = vertex_color;

    gl_Position = projection * camera * model * vec4(vertex_position, 1.0);
}
` + "\x00"

// TODO: Fetch these from a configuration file
const winTitle string = "Go-SDL2 + Go-GL"
const winWidth, winHeight int32 = 640, 480

func MessageCallback(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
	if id == 131169 || id == 131185 || id == 131218 || id == 131204 {
		return
	}

	fmt.Println("---------------------------")
	fmt.Println("OpenGL debug message (", id, "): ", message)

	switch source {
	case gl.DEBUG_SOURCE_API:              fmt.Println("source: api")
	case gl.DEBUG_SOURCE_WINDOW_SYSTEM:    fmt.Println("source: window system")
	case gl.DEBUG_SOURCE_SHADER_COMPILER:  fmt.Println("source: shader compiler")
	case gl.DEBUG_SOURCE_THIRD_PARTY:      fmt.Println("source: third party")
	case gl.DEBUG_SOURCE_APPLICATION:      fmt.Println("source: application")
	case gl.DEBUG_SOURCE_OTHER:            fmt.Println("source: other")
	}
	switch gltype {
	case gl.DEBUG_TYPE_ERROR:                fmt.Println("type: error")
	case gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR:  fmt.Println("type: deprecated behavior")
	case gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR:   fmt.Println("type: undefined behavior")
	case gl.DEBUG_TYPE_PORTABILITY:          fmt.Println("type: portability")
	case gl.DEBUG_TYPE_PERFORMANCE:          fmt.Println("type: performance")
	case gl.DEBUG_TYPE_MARKER:               fmt.Println("type: marker")
	}
	switch severity {
	case gl.DEBUG_SEVERITY_HIGH:          fmt.Println("severity: high")
	case gl.DEBUG_SEVERITY_MEDIUM:        fmt.Println("severity: medium")
	case gl.DEBUG_SEVERITY_LOW:           fmt.Println("severity: low")
	case gl.DEBUG_SEVERITY_NOTIFICATION:  fmt.Println("severity: notification")
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	// Init SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create window from SDL
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE); if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Force core profile
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4); err != nil { panic(err) }
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 2); err != nil { panic(err) }
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE); err != nil { panic(err) }

	// Create GL context
	context, err := window.GLCreateContext(); if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	// Init Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	gl.Enable(gl.DEBUG_OUTPUT)
	gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
	gl.DebugMessageCallback(MessageCallback, gl.Ptr(nil))
	gl.DebugMessageControl(gl.DONT_CARE, gl.DONT_CARE, gl.DONT_CARE, 0, nil, true)

	// Configure vertex and fragment shaders
//	program, err := gfx.NewProgram(
//		"assets/shaders/vertex.glsl",
//		"assets/shaders/fragment.glsl"); if err != nil {
//		panic(err)
//	}
	program, err := gfx.NewProgram(
		__vertexShader,
		__fragmentShader); if err != nil {
		panic(err)
	}
	
	gl.UseProgram(program)

	projection := qt.PerspectiveProjectionMat4(qt.Radians(45.0), float32(winWidth)/float32(winHeight), 0.1, 50.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := qt.LookAtV(qt.Vec3{3, 3, 3}, qt.Vec3{0, 0, 0}, qt.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	fmt.Printf("|%v,%v,%v,%v|\n|%v,%v,%v,%v|\n|%v,%v,%v,%v|\n|%v,%v,%v,%v|\n\n\n",
	camera[0], camera[4], camera[8],  camera[12],
	camera[1], camera[5], camera[9],  camera[13],
	camera[2], camera[6], camera[10], camera[14],
	camera[3], camera[7], camera[11], camera[15])

	gl.BindFragDataLocation(program, 0, gl.Str("fs_color\x00"))

	// Mesh init
	foo := gfx.NewTransform(qt.Vec3{0,0,0}, qt.Vec3{0,0,0}, qt.Vec3{0,0,0})
	mesh := gfx.NewMeshFromPrimitive(gfx.Cube(), &foo)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.2, 0.3, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LESS)

	// VSync
	err = sdl.GLSetSwapInterval(1); if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// ECS
	world := ecs.World{}

	renderingSystem := gfx.RenderableSystem{}
	_baseEntity := ecs.NewBaseEntity()
	_transform := gfx.TransformComponent{gfx.NewTransform(qt.Vec3{0,0,0}, qt.Vec3{0,0,0}, qt.Vec3{0,0,0})}
	_baseent := gfx.RenderableEntity{BaseEntity: &_baseEntity, TransformComponent: &_transform}
	renderingSystem.Add(_baseent.BaseEntity, _baseent.TransformComponent)

	world.AddSystem(&renderingSystem)

//	angle := 0.0
	var event sdl.Event
	var running bool = true
	for running {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update
		world.Update(0.125)

		foo.Pos = qt.Vec3{0, 0, foo.Pos[2] - 0.003}
		foo.Origin = qt.Vec3{0, 0, foo.Origin[2] - 0.003}
		foo.EulerRot = qt.Vec3{0, foo.EulerRot[1] * 1.01 + 0.01, 0}

		// Render
		gl.UseProgram(program)
		
		mesh.Draw(program)


		// Maintenance

		// Swap buffers
		window.GLSwap()

		// Poll events -- I use SDL for event handling instead of something like GLFW
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
				case *sdl.MouseMotionEvent:
					fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}
	}
	mesh.Free()
}
