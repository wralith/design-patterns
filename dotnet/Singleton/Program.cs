using Singleton;

var firstLogger = Logger.GetInstance();
var secondLogger = Logger.GetInstance();

if (firstLogger == secondLogger)
{
    Console.WriteLine("they are same.");
}
