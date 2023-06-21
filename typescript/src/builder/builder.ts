export class SQLBuilder {
    private statement = ""

    public select(...columns: string[]): SQLBuilder {
        this.statement += `SELECT ${columns.join(", ")} `
        return this
    }

    public from(table: string): SQLBuilder {
        this.statement += `FROM ${table} `
        return this
    }

    public where(condition: string): SQLBuilder {
        this.statement += `WHERE ${condition} `
        return this
    }

    public build(): string {
        return this.statement.slice(0, -1)
    }
}
