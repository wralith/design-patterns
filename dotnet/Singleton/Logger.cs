namespace Singleton;

public sealed class Logger
{
    private static Logger? _instance;

    public static Logger GetInstance()
    {
        return _instance ??= new Logger();
    }

    public void Log()
    {
        //your business code
    }
}