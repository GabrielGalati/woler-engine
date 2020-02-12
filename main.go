package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	log "github.com/sirupsen/logrus"
)

var vertices = []float32{
	-0.5, -0.5, 0.0,
	-0.5, 0.5, 0.0,
	0.5, 0.5, 0.0,
	0.5, -0.5, 0.0,
}

var rectangle = []uint{
	0, 1, 2,
	2, 3, 0,
}

func init() {
	runtime.LockOSThread()
}

func main() {
	window := initGLFW()
	defer glfw.Terminate()

	program := initOpenGL()

	var vertexBuffer uint32
	var elementBuffer uint32
	var vertexArray uint32

	gl.GenBuffers(1, &vertexBuffer)
	gl.GenBuffers(1, &elementBuffer)
	gl.GenVertexArrays(1, &vertexArray)

	gl.BindVertexArray(vertexArray)

	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(rectangle), gl.Ptr(rectangle), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	if err := gl.GetError(); err != 0 {
		log.Error(err)
	}

	for !window.ShouldClose() {
		gl.ClearColor(0.5, 0.5, 0.5, 0.5)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)

		gl.BindVertexArray(vertexArray)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		//gl.DrawArrays(gl.TRIANGLES, 0, 3)

		glfw.PollEvents()
		window.SwapBuffers()

		if err := gl.GetError(); err != 0 {
			log.Error(err)
		}
	}

	glfw.Terminate()
}
