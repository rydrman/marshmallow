/*
Package mm is the main package for the Marshmallow graphics library
for golang. It borrows shamlessly from three.js in it's layout and
overall simlicity goals.
*/
package mm

import (
    "runtime"

    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
)

func init() {

    runtime.LockOSThread()

    if err := glfw.Init(); err != nil {
        panic(err)
    }

    if err := gl.Init(); err != nil {
        panic(err)
    }

}
