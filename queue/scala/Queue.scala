import scala.reflect.ClassTag

// TODO: headとtailが、教会を超えるかどうかの判定をしなくてもいいようにしたい
// TODO: immutable vecではないデータ構造を使う
class Queue[A: ClassTag](size: Int) {
    private val ringArray = Vector.fill(size)()
    private var head = 0
    private var tail = 0

    def enqueue(elm: A): Unit = {
        val _ = this.ringArray.updated(tail, elm)
        this.tailBack
    }

    def dequeue: A = {
        this.headAdvance
        this.ringArray
    }
    
    private def tailBack: Unit = {
        if(this.tail == 0) {
            this.tail = size - 1
        } else {
            this.tail -= 1
        }
    }
    private def headAdvance: Unit = {
        if(this.head == size - 1) {
            this.head = 0
        } else {
            this.head += 1
        }
    }

    def queueSize: Int = this.ringArray.size
    def headPos: Int = this.head
    def tailPos: Int = this.tail
}

object Main {
    def main(args: Array[String]) = {
        println("hello world")
        val q = new Queue[Int](10)
        q.enqueue(3)
        println(q.tailPos)
    }
}
