package gfx

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)




const winWidth, winHeight int32 = 640, 480

func InitGBuffer() error {

	var gBuffer uint32
	gl.GenFramebuffers(1, &gBuffer)

	var gPosition, gNormal, gAlbedoSpec uint32
	// Position colorbuffer
	gl.GenTextures(1, &gPosition)
	gl.BindTexture(gl.TEXTURE_2D, gPosition)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB16F, winWidth, winHeight, 0, gl.RGB, gl.FLOAT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, gPosition, 0)
	// Normal colorbuffer
	gl.GenTextures(1, &gNormal)
	gl.BindTexture(gl.TEXTURE_2D, gNormal)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB16F, winWidth, winHeight, 0, gl.RGB, gl.FLOAT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT1, gl.TEXTURE_2D, gPosition, 0)
	// Color-and-specular colorbuffer
	gl.GenTextures(1, &gAlbedoSpec)
	gl.BindTexture(gl.TEXTURE_2D, gAlbedoSpec)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, winWidth, winHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT2, gl.TEXTURE_2D, gAlbedoSpec, 0)

	// Tell OpenGL which color attachments (of this framebuffer) we will use for rendering
	attachments := [3]uint32{gl.COLOR_ATTACHMENT0, gl.COLOR_ATTACHMENT1, gl.COLOR_ATTACHMENT2}
	gl.DrawBuffers(3, &attachments[0])

	// Create and attach depth buffer (renderbuffer)
	var rboDepth uint32
	gl.GenRenderbuffers(1, &rboDepth)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rboDepth)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT, winWidth, winHeight)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, rboDepth)

	// Finally, check if framebuffer is complete
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return fmt.Errorf("incomplete framebuffer")
	}


	return nil
}
















