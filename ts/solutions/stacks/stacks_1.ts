class MinStack {
  top: MinFrame
  minFrame: MinFrame

  constructor(top?: MinFrame, minFrame?: MinFrame) {
    this.top = top
    this.minFrame = minFrame
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
      const min = this.minFrame ? this.minFrame.data : 'null'
      return `${stackStrings.join(' -> ')}; Min: ${min}`
    }

    return 'empty stack'
  }

  push(data: number) {
    const frame = new MinFrame(data, this.top)

    if (this.minFrame) {
      if (data < this.minFrame.data) {
        frame.previousMinFrame = this.minFrame
        this.minFrame = frame
      }
    } else {
      this.minFrame = frame
    }

    this.top = frame
  }

  pop(): number {
    if (!this.top) {
      throw 'stack is empty'
    }

    if (this.top === this.minFrame) {
      this.minFrame = this.top.previousMinFrame
    }

    const frame = this.top
    this.top = this.top.next

    return frame.data
  }

  min(): number {
    if (!this.top) {
      throw 'stack is empty'
    }

    return this.minFrame.data
  }
}

class MinFrame {
  data: number
  next: MinFrame
  previousMinFrame: MinFrame

  constructor(data?: number, next?: MinFrame, previousMinFrame?: MinFrame) {
    this.data = data
    this.next = next
    this.previousMinFrame = previousMinFrame
  }

  toString() {
    return `[ ${this.data} ]`
  }
}
