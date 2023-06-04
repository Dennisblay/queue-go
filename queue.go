package main

import "fmt"

func main__() {
	q := Queue{}
	for i := 1; i <= 10; i++ {
		q.enqueue(i)
	}
	fmt.Println(q.size)
	q.DisplayList()
	fmt.Println("done")
	//a := q.dequeue()

}

type Queue struct {
	LinkedList
}

func (q *Queue) enqueue(data interface{}) {
	q.AddNode(data)
}

func (q *Queue) dequeue() interface{} {
	front := q.front()
	if !q.isEmpty() {
		q.Remove(q.head.data)
	}
	return front

}

func (q *Queue) front() interface{} {
	if q.isEmpty() {
		return nil
	}
	return q.head.data
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}
