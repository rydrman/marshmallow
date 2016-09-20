package mm_test

import (
    "testing"

    "github.com/rydrman/marshmallow"
)

func TestIterateDepth(t *testing.T) {

    objs := []*mm.Object{
        mm.NewObject(),
        mm.NewObject(),
        mm.NewObject(),
        mm.NewObject(),
    }

    objs[0].Add(objs[1])
    objs[0].Add(objs[3])
    objs[1].Add(objs[2])

    i := 0
    for o := range mm.IterateDepth(objs[0]) {

        if o != objs[i] {

            t.Error("depth iteration yielded object in wrong order")
            t.Errorf("expected: %v", objs[i])
            t.Errorf("got: %v", o)

        }

        i++

    }

}
