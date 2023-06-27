namespace Logger;

public class ConsoleLogger
{
    private static ConsoleLogger? _instance;

    public static ConsoleLogger GetInstance() => _instance ??= new ConsoleLogger();

    public void Log(string message)
    {
        Console.WriteLine($"Log Message: {message}");
    }
}