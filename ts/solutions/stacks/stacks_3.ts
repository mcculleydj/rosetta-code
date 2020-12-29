import { Stack } from './stacks_2'

class Queue {
  front: Stack
  back: Stack

  constructor() {
    this.front = new Stack()
    this.back = new Stack()
  }

  toString(): string {
    return `Front: ${this.front.toString()}\nBack: ${this.back.toString()}`
  }

  private transfer(source: Stack, destination: Stack): void {
    while (source.top) {
      destination.push(source.pop())
    }
  }

  add(data: number): void {
    if (!this.back.top) {
      this.transfer(this.front, this.back)
    }
    this.back.push(data)
  }

  remove(): number {
    if (!this.front.top) {
      this.transfer(this.back, this.front)

      if (!this.front.top) {
        throw 'queue is empty'
      }
    }

    return this.front.pop()
  }
}
