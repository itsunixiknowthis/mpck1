package mpck1

import "testing"
import "github.com/itsunixiknowthis/mpck1/b"

func TestA(t *testing.T) {
    if 123 != A { t.Errorf("!") }
    if "lol" != B { t.Errorf("?") }
    if 444 != C { t.Errorf("# %v", C) }
    if 444 != b.X { t.Errorf("$ %v", C) }
}

