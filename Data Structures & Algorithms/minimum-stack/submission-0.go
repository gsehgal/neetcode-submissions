type MinStack struct {
    stack []int
    mstack []int
}

func Constructor() MinStack {
    s := MinStack {
        stack : make([]int,0),
        mstack: make([]int,0),
    }
    return s
}

func (this *MinStack) Push(val int) {
    if len(this.stack)==0 {
        this.stack = append(this.stack,val)
        this.mstack = append(this.mstack,val)
    } else {
        if val <= this.mstack[len(this.mstack)-1] {
            this.mstack = append(this.mstack,val)
        }
        this.stack = append(this.stack,val)           
    }
}

func (this *MinStack) Pop() {
    top := this.stack[len(this.stack)-1]
    if top == this.mstack[len(this.mstack)-1] {
        this.mstack = this.mstack[:len(this.mstack)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
   return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
   return this.mstack[len(this.mstack)-1]
}
