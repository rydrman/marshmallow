package mm

type Projector interface {
    Positioner
    GetProjectionMatrix() *Matrix4
}
