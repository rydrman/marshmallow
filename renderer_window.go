package mm

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
)

type WindowRenderer struct {
    title  string
    window *glfw.Window
}

func NewWindowRenderer(w, h int, title string) (*WindowRenderer, error) {

    r := &WindowRenderer{
        title: title,
    }

    err := r.setupWindow(w, h)
    if nil != err {
        return nil, err
    }

    return r, nil

}

func (r *WindowRenderer) setupWindow(w, h int) error {

    if r.window != nil {

        r.window.Destroy()

    }

    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 2)
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    //glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    //glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    win, err := glfw.CreateWindow(w, h, r.title, nil, nil)
    if err != nil {
        return err
    }
    r.window = win

    return nil

}

func (r *WindowRenderer) Render(scene Node, camera Projector) {

    r.window.MakeContextCurrent()
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    r.window.SwapBuffers()

}
