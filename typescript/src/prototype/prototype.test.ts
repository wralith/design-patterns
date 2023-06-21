import { DataFrame } from "./prototype"

test("Clone DataFrame", () => {
  const prototype = new DataFrame("Test", {
    dataMap: new Map([
      ["name", "test"],
    ]),
  })

  const clone = prototype.clone()
  expect(clone).not.toEqual(prototype)
  expect(clone.data).not.toEqual(prototype.data)
  expect(clone.data.dataMap.get("name")).toEqual(prototype.data.dataMap.get("name"))

  clone.data.dataMap.set("version", 2)
  expect(prototype.data.dataMap.has("version")).toBeFalsy()
})