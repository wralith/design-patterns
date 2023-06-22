import { Logger } from "./singleton"

test("When execute newLogger, it should return the same instance of Logger every time", () => {
  const instance1 = Logger.newLogger()
  const instance2 = Logger.newLogger()

  expect(instance1).toBe(instance2)
  expect(instance1 === instance2).toBeTruthy()
})
