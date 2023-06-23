type Message = { data: string }

// We want implementation of Subscriber to have update method
export interface Subscriber {
  update(message: Message): void
}

// Broker class is responsible for managing subscribers
export class Publisher {
  private subscribers: Subscriber[]

  constructor(...subscribers: Subscriber[]) {
    this.subscribers = subscribers ?? []
  }

  sub(subscriber: Subscriber) {
    this.subscribers.push(subscriber)
  }

  pub(message: Message) {
    this.subscribers.forEach((subscriber) => subscriber.update(message))
  }
}
