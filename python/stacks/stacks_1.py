class MinStack:
    def __init__(self, top=None, min_frame=None):
        self.top = top
        self.min_frame = min_frame

    def __repr__(self):
        frames = []
        top = self.top

        while top:
            frames.append(str(top))
            top = top.next

        # [] is falsy in Python
        if frames:
            return f'{" -> ".join(["top"] + frames + ["bottom"])}; Min: {self.min_frame.data if self.min_frame else "None"}'

        return 'empty stack'

    def push(self, data):
        frame = MinFrame(data, self.top)

        # assumes that "data" is comparable to other "data" using standard operators
        if self.min_frame:
            if data < self.min_frame.data:
                frame.previous_min_frame = self.min_frame
                self.min_frame = frame
        else:
            self.min_frame = frame

        self.top = frame

    def pop(self):
        if not self.top:
            raise Exception('stack is empty')

        if self.top is self.min_frame:
            self.min_frame = self.top.previous_min_frame

        frame = self.top
        self.top = self.top.next

        return frame.data

    def min(self):
        if not self.top:
            raise Exception('stack is empty')

        return self.min_frame.data


class MinFrame:
    def __init__(self, data=None, next=None, previous_min_frame=None):
        self.data = data
        self.next = next
        self.previous_min_frame = previous_min_frame

    def __repr__(self):
        return f'[ {self.data} ]'
