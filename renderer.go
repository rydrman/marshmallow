package mm

type Renderer interface {
    Render(*Scene, Projector)
}
