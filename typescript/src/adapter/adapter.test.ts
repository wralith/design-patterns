import { AdaptToNewDatabase, OldDatabase } from "./adapter"

const mockOldDatabase: OldDatabase = {
  create: jest.fn(),
  update: jest.fn(),
  has: jest.fn(),
}

describe("OldDatabase to NewDatabase Adapter", () => {
  const oldDatabaseUpdateSpy = jest.spyOn(mockOldDatabase, "update")
  const oldDatabaseCreateSpy = jest.spyOn(mockOldDatabase, "create")
  const adapted = AdaptToNewDatabase(mockOldDatabase)

  it("Should execute create method of old database class if database has entry with given id", () => {
    adapted.upsert({ id: "1", post: "test" })
    expect(oldDatabaseCreateSpy).toHaveBeenCalledWith({ id: "1", post: "test" })
  })

  it("Should execute update method of old database class if database entry with given id doesn't exists", () => {
    mockOldDatabase.has = jest.fn().mockReturnValue(true)
    adapted.upsert({ id: "1", post: "test" })
    expect(oldDatabaseUpdateSpy).toHaveBeenCalledWith({ id: "1", post: "test" })
  })
})
