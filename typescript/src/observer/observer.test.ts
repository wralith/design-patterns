import { Publisher, Subscriber } from "./observer"

const sub1: Subscriber = {
  update: jest.fn(),
}

const sub2: Subscriber = {
  update: jest.fn(),
}

const sub3: Subscriber = {
  update: jest.fn(),
}

test("Observer", () => {
  const broker = new Publisher(sub1, sub2)
  broker.sub(sub3)
  broker.pub({ data: "Hello" })

  expect(sub1.update).toHaveBeenCalledWith({ data: "Hello" })
  expect(sub2.update).toHaveBeenCalledWith({ data: "Hello" })
  expect(sub3.update).toHaveBeenCalledWith({ data: "Hello" })
})
