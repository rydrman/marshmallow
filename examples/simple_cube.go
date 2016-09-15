package main

import "github.com/rydrman/marshmallow"

func main() {

    win, err := mm.NewWindowRenderer(500, 500)

    scene := mm.NewScene()

    geo := mm.NewCubeGeometry(1, 1, 1)
    mat := mm.NewBasicMaterial(mm.MaterialProps{
        Color: mm.Color().FromHex(0x001111),
    })

    mesh := mm.NewMesh(geo, mat)

    scene.Add(mesh)

    cam := mm.NewPerspectiveCamera(1)

    win.Render(scene, cam)

}
