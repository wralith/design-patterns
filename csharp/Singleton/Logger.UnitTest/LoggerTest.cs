namespace Logger.UnitTest;

public class LoggerTest
{
    [Fact]
    public void GetInstance_ReturnsSameInstance()
    {
        var instance1 = ConsoleLogger.GetInstance();
        var instance2 = ConsoleLogger.GetInstance();

        Assert.Same(instance1, instance2);
    }

    [Fact]
    public void Log_WritesMessageToConsole()
    {
        var logger = ConsoleLogger.GetInstance();
        string message = "test";
        string expected = $"Log Message: {message}";
        string actual;

        using (var writer = new StringWriter())
        {
            Console.SetOut(writer);
            logger.Log(message);

            actual = writer.ToString().Trim();
        }
        
        Assert.Equal(expected, actual);
    }
}