package mm_test

import (
    "math"

    "github.com/rydrman/marshmallow"
)

//these are all constants to be used in testing

const (
    x = 2
    y = 3
    z = 4
    w = 5
)

var (

    //var negInf2 = new THREE.Vector2( math.Inf(-1), math.Inf(-1) )
    //var posInf2 = new THREE.Vector2( math.Inf(1), math.Inf(1) )

    //var zero2 = new THREE.Vector2()
    //var one2 = new THREE.Vector2( 1, 1 )
    //var two2 = new THREE.Vector2( 2, 2 )

    negInf3 = mm.NewVector3().Set(math.Inf(-1), math.Inf(-1), math.Inf(-1))
    posInf3 = mm.NewVector3().Set(math.Inf(1), math.Inf(1), math.Inf(1))

    zero3 = mm.NewVector3()
    one3  = mm.NewVector3().Set(1, 1, 1)
    two3  = mm.NewVector3().Set(2, 2, 2)
)
