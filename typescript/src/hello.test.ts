import { addHello } from "./hello"

test("addHello function should add 'Hello' before the given word", () => {
  const want = "Hello World"
  const got = addHello("World")
  expect(got).toEqual(want)
})
