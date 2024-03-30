package synchronizers

import (
    "fmt"
)
type RClone struct {
    
}

func (m *RClone) Execute() {
    fmt.Println("MyStruct running")
}
