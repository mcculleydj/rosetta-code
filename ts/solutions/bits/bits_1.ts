/*
JS uses 32-bit bitwise operands. JS stores numbers as 64 bit floating point numbers, but all bitwise operations are performed on 32 bits binary numbers. Before a bitwise operation is performed, JS converts numbers to 32 bits signed integers.
*/

// O(n), where n is the number of bits necessary to represent
// the larger of the 2 numbers
function add(x: number, y: number): number {
  if (!Number.isSafeInteger(x) || !Number.isSafeInteger(y)) {
    throw 'arguments are not precision safe'
  }

  console.log(x, y)

  while (y !== 0) {
    const carry = x & y
    x ^= y
    const y_ = y
    y = carry << 1

    if ((y_ < 0 && y > 0) || (y_ > 0 && y < 0)) {
      throw 'integer overflow'
    }
  }

  return x
}
