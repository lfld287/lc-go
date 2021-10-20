package queue

type(
	Queue struct {
		top *node
		rear *node
		length int
	}
	node struct {
		pre *node
		next *node
		value interface{}
	}
)

func New() *Queue{
	return &Queue{
		top:    nil,
		rear:   nil,
		length: 0,
	}
}

func (p *Queue) len() int {
	return p.length
}

func (p *Queue) Any() bool {
	return p.length > 0
}

func (p *Queue) Peek() interface{} {
	if p.top == nil {
		return nil
	}
	return p.top.value
}

func (p *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if p.length == 0 {
		p.top = n
		p.rear = p.top
	} else {
		n.pre = p.rear
		p.rear.next = n
		p.rear = n
	}
	p.length++
}

func (p *Queue) Pop() interface{} {
	if p.length == 0 {
		return nil
	}
	n := p.top
	if p.top.next == nil {
		p.top = nil
	} else {
		p.top = p.top.next
		p.top.pre.next = nil
		p.top.pre = nil
	}
	p.length--
	return n.value
}