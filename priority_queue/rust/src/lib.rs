use std::fmt::Debug;


#[derive(Debug)]
struct Elm<T:Sized+Debug> {
    elm: T,
    priority: i32
}
impl<T:Sized+Debug> Elm<T> {
    fn new(elm: T, priority: i32) -> Self{
        Self {
            elm: elm,
            priority: priority
        }
    }
}

struct PriorityQueue<T:Sized+Debug> {
    head: usize,
    tail: usize,
    innerArray: Vec<Elm<T>>,
    queueSize: usize
}

impl<T:Sized+Debug> PriorityQueue<T> {
    // とりあえず要素0
    fn new(size: usize) -> Self {
        
        Self {
            head: 0,
            tail: 0,
            innerArray: Vec::new(),
            queueSize: size
        }
    }
    fn enqueue(&mut self, elm: T, priority: i32) -> () {
        // self.innerArray
        if self.innerArray.len() < self.queueSize {
            self.innerArray.push(Elm::new(elm, priority));
            self.innerArray.sort_by(|e1, e2| e1.priority.cmp(&e2.priority));
            self.increHead()
        } else {
            // とりあえずheadの次に追加する
            // TODO: tailを抜かす時のハンドリング
            self.innerArray[self.head] = Elm::new(elm, priority);
            self.increHead()
        }
    }
    // TODO: TがCopyを実装していない場合のimpl
    // fn dequeue(&mut self) -> Option<T> {
    // }
    fn headDecre(&mut self) {
        self.head-=1
    }
    fn increHead(&mut self) {
        self.head+=1
    }
    fn print(&self) {
        for i in 0 .. self.innerArray.len() {
            print!("{:?} ", self.innerArray[i]);
            println!("");
        }
    }
}
impl<T:Sized + Debug + Copy> PriorityQueue<T> {
    fn dequeue(&mut self) -> Option<T> {
        if self.head == self.tail {
            return None
        }
        let e =  self.innerArray[self.head - 1].elm;
        self.headDecre();
        Some(e)
    }
}

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut q:PriorityQueue<&str> = PriorityQueue::new(10);
        q.enqueue("motoyuki", 24);
        q.enqueue("yasu", 20);
        q.enqueue("haru", 22);
        q.print();
        println!("{:?}", q.dequeue());
        println!("{:?}", q.dequeue());
        println!("{:?}", q.dequeue());
        println!("{:?}", q.dequeue());
        println!("{:?}", q.dequeue());
        println!("{:?}", q.dequeue());
    }
}
