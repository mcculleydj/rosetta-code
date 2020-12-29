class StackSet {
  stacks: Stack[]
  index: number
  stackCapacity: number

  constructor(stackCapacity: number) {
    this.stacks = []
    this.index = 0
    this.stackCapacity = stackCapacity
  }

  toString(): string {
    let stacks = ''

    this.stacks.forEach((s, i) => {
      stacks += `${i}: ${s.toString()}\n`
    })

    return stacks
  }

  pop(): number {
    if (this.stacks.length === 0) {
      throw 'stack set is empty'
    }

    while (!this.stacks[this.index].top) {
      this.stacks.pop()
      this.index--

      if (this.index === -1) {
        this.index = 0
        throw 'stack set is empty'
      }
    }

    return this.stacks[this.index].pop()
  }

  push(data: number) {
    if (this.stacks.length === 0) {
      this.stacks.push(new Stack())
      this.push(data)
    } else if (this.stacks[this.index].length == this.stackCapacity) {
      this.index++
      this.stacks.push(new Stack())
      this.push(data)
    } else {
      this.stacks[this.index].push(data)
    }
  }

  popAt(i: number): number {
    if (i >= this.stacks.length) {
      throw 'index out of bounds'
    }

    if (!this.stacks[i].top) {
      throw 'stack is empty'
    }

    return this.stacks[i].pop()
  }
}

class Stack {
  top: Frame
  length: number

  constructor() {
    this.top = null
    this.length = 0
  }

  toString() {
    const frames: string[] = []
    let top = this.top

    while (top) {
      frames.push(top.toString())
      top = top.next
    }

    if (frames.length > 0) {
      const stackStrings = ['top', ...frames, 'bottom']
      return stackStrings.join(' -> ')
    }

    return 'empty stack'
  }

  push(data: number) {
    const frame = new Frame(data, this.top)
    this.top = frame
    this.length++
  }

  pop(): number {
    if (!this.top) {
      throw 'stack is empty'
    }

    const frame = this.top
    this.top = this.top.next
    this.length--

    return frame.data
  }
}

class Frame {
  data: number
  next: Frame

  constructor(data: number, next?: Frame) {
    this.data = data
    this.next = next
  }

  toString(): string {
    return `[ ${this.data} ]`
  }
}

const ss = new StackSet(5)
ss.push(1)
ss.push(2)
ss.push(3)
ss.push(4)
ss.push(5)
ss.push(1)
ss.push(2)
ss.push(3)
ss.push(4)
ss.push(5)
ss.push(1)
ss.push(2)
ss.push(3)
ss.push(4)
ss.push(5)

console.log(ss.toString())

console.log(ss.popAt(1))
console.log(ss.popAt(1))
console.log(ss.popAt(1))
console.log(ss.popAt(1))
console.log(ss.popAt(1))

console.log(ss.toString())

console.log(ss.pop())
console.log(ss.pop())
console.log(ss.pop())
console.log(ss.pop())
console.log(ss.pop())

console.log(ss.toString())

console.log(ss.pop())
console.log(ss.toString())
