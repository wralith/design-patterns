interface Cloneable<T> {
  clone(): T;
}

// I am nesting for the sake of nesting
type Data = {
  dataMap: Map<string, any>
}

// Dataframe takes data in the constructor which is an object type so passed by reference
export class DataFrame implements Cloneable<DataFrame> {
  name: string
  data: Data

  constructor(name: string, data: Data) {
    this.name = name
    this.data = data
  }

  clone(): DataFrame {
    // structuredClone is a native function that clones objects
    // If you want to use something else, go ahead like lodash
    // or your own implementation, which will be easy for this example
    return structuredClone(this)
  }
}
