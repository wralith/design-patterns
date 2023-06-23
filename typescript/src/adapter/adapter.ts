export type Entry = {
  id: string
  post: string
}

export interface OldDatabase {
  has(id: string): Entry
  create(entry: Entry): void
  update(entry: Entry): void
}

export interface NewDatabase {
  has(id: string): Entry
  upsert(entry: Entry): void
}

// Simple example, simple code...

export const AdaptToNewDatabase = (old: OldDatabase): NewDatabase => ({
  has: old.has,
  upsert: (entry) => (old.has(entry.id) ? old.update(entry) : old.create(entry)),
})
