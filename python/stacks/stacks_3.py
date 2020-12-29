from stacks_2 import Stack

# for stack class definition see Stack Set

# reverse one stack into the other as necessary
# pushing only onto the back stack
# and popping only from the front stack


class Queue:
    def __init__(self):
        self.front = Stack()
        self.back = Stack()

    def __repr__(self):
        return f'front: {self.front}\nback: {self.back}'

    def add(self, data):
        if not self.back.top:
            self.transfer(self.front, self.back)
        self.back.push(data)

    def remove(self):
        if not self.front.top:
            self.transfer(self.back, self.front)
            if not self.front.top:
                raise Exception('queue is empty')

        return self.front.pop()

    def transfer(self, source, destination):
        while source.top:
            destination.push(source.pop())
