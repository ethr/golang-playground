// Interfaces and their implementations are not as closely linked as they are in
// say Java or C++. There's no clear way for interfaces to inherit from each
// other or extend each other.
// In Java you might write the following code as:
//
// public interface SubmitAble {}
// public interface Closable {}
// public interface SubmitServer extends SubmitAble, Closable {}
// public class SubmitServerImpl implements SubmitServer {}
//
// Through this you're saying SubmitServerImpl "is a" SubmitAble and a Closable.
//
// In go this link is very vague, is more like "SubmitServerImpl" can be treated
// like a "Closer" as it implements the functions required for a Closer. This is
// far weaker than saying it is a Closer, it might have the same method
// signature for some other reason

package main

import "fmt"

type Submitter interface {
        Submit(s string) string
}

type Closer interface {
        Close()
}

// Duplicates methods from Submitter and Closer because inheritance is messed up
// in Go when compared with traditional languages like C++ or Java
type SubmitServer interface {
        Submit(s string) string
        Close()
}

type SubmitServerImpl struct {
        something int
}

func (this *SubmitServerImpl) Submit(s string) string {
        return s
}

func (this *SubmitServerImpl) Close() {
        fmt.Println("Closed")
}

func NewSubmitServerImpl() *SubmitServerImpl {
        s := new(SubmitServerImpl)
        return s
}

func main() {
        var server SubmitServer = NewSubmitServerImpl()
        fmt.Println(server.Submit("foo"))
        server.Close()
}
