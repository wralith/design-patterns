export class Logger {
  private static instance: Logger

  private constructor() {}

  // If Logger haven't been instantiated yet, do it!
  // Return the same instance of Logger whenever newLogger executed
  public static newLogger(): Logger {
    if (!Logger.instance) {
      Logger.instance = new Logger()
    }
    return Logger.instance
  }

  // Because why not? Example intensifies
  public info(message: string): void {
    const now = new Date()
    console.log(`${now.toLocaleDateString()} ${now.toLocaleTimeString()} - INFO: ${message}`)
  }
}
