package mm

type Scene struct {
    *Object

    BackgroundColor *Color

    AutoUpdate bool
}

func NewScene() *Scene {

    return &Scene{
        Object: NewObject(),

        BackgroundColor: Colors("Black"),

        AutoUpdate: true,
    }

}
