class StackSet:
    def __init__(self, stack_capacity=None):
        self.stacks = []
        self.index = None
        self.stack_capacity = stack_capacity

    def __repr__(self):
        stacks = ""

        for i, s in enumerate(self.stacks):
            stacks += f'{i}: {s}\n'

        return stacks

    def pop(self):
        if self.index == None:
            raise Exception('stack set is empty')

        # find a non-empty stack
        # necessary to support popAt(i)
        while not self.stacks[self.index].top:
            # Python list built in pop
            self.stacks.pop()
            self.index -= 1

            if self.index == -1:
                self.index = None
                raise Exception('stack set is empty')

        return self.stacks[self.index].pop()

    def push(self, data):
        if self.index == None:
            self.index = 0
            self.stacks.append(Stack())
            self.push(data)
        elif self.stacks[self.index].length == self.stack_capacity:
            self.index += 1
            self.stacks.append(Stack())
            self.push(data)
        else:
            self.stacks[self.index].push(data)

    def popAt(self, i):
        if i >= len(self.stacks):
            raise IndexError

        if not self.stacks[i].top:
            raise Exception('stack is empty')

        return self.stacks[i].pop()


class Stack:
    def __init__(self, top=None):
        self.top = top
        self.length = 0

    def __repr__(self):
        frames = []
        top = self.top

        while top:
            frames.append(str(top))
            top = top.next

        if frames:
            return f'{" -> ".join(["top"] + frames + ["bottom"])}'

        return 'empty stack'

    def push(self, data):
        frame = Frame(data, self.top)
        self.top = frame
        self.length += 1

    def pop(self):
        if not self.top:
            raise Exception('stack is empty')

        frame = self.top
        self.top = self.top.next
        self.length -= 1

        return frame.data


class Frame:
    def __init__(self, data=None, next=None):
        self.data = data
        self.next = next

    def __repr__(self):
        return f'[ {self.data} ]'


ss = StackSet(5)
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

print(ss)

print(ss.popAt(1))
print(ss.popAt(1))
print(ss.popAt(1))
print(ss.popAt(1))
print(ss.popAt(1))

print(ss)

print(ss.pop())
print(ss.pop())
print(ss.pop())
print(ss.pop())
print(ss.pop())

print(ss)

print(ss.pop())
print(ss)
