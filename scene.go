package mm

type Scene struct {
    *Object

    AutoUpdate bool
}

func NewScene() *Scene {
    return &Scene{
        Object: NewObject(),

        AutoUpdate: true,
    }
}
