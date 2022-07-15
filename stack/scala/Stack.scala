class Stack[A] {
    case class Node[A](elm: A, next: Option[Node[A]]) {
    }
    var head:Option[Node[A]] = None
    def push(elm: A):Unit = {
        this.head match {
            case None => this.head = Some(Node(elm, None))
            case Some(e) => {
                val curHead = e
                this.head = Some(Node(elm, Some(e)))
            }
        }
    }
    def pop:A = this.head.get.elm
    def print_elm = println(this.head)
}

// ref: https://codereview.stackexchange.com/questions/253095/stack-implementation-in-scala
// Listを使ったstack. こっちのがscalaっぽい？
class StackWithList[A] {
    private var list: List[A] = List[A]()
    def push(elm: A) = {
        this.list = elm :: list
    }
    def pop: Option[A] = {
        if (this.list.isEmpty) {
            None
        } else {
            val head = this.list.head
            this.list = this.list.tail
            Some(head)
        }
    }
    def print: Unit = println(this.list)
}

// main
object Main {
    def main(args: Array[String]):Unit = {
        val s: Stack[Int] = new Stack()
        s.push(3)
        s.push(4)
        s.push(5)
        s.print_elm
        println(s.pop)

        val s2: StackWithList[Int] = new StackWithList()
        s2.list
        s2.push(1)
        s2.print
        s2.push(2)
        s2.print
        println(s2.pop)
        s2.print
        s2.pop
        s2.print
    }
}
