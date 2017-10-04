using System;
using RabbitMQ.Client;
namespace MyNameSpace
{
    class Program
    {
        static void Main(string[] args)
        {
            var a=new RabbitMQ.Client.ConnectionFactory();
            Console.WriteLine("Hello World!");
        }
    }
}
