import { SQLBuilder } from "./builder"

test("SQLBuilder", () => {
  const cases: { want: string, got: string }[] = [
    {
      want: "SELECT * FROM users",
      got: new SQLBuilder().select("*").from("users").build(),
    },
    {
      want: "SELECT * FROM users WHERE id = 1",
      got: new SQLBuilder().select("*").from("users").where("id = 1").build(),
    },
    {
      want: "SELECT username, email FROM users WHERE id = 1",
      got: new SQLBuilder().select("username", "email").from("users").where("id = 1").build(),
    },
  ]

  cases.forEach(c => expect(c.got).toEqual(c.want))
})