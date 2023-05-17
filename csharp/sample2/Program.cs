// See https://aka.ms/new-console-template for more information

using System;

class Program
{
    static void Main()
    {
        object[] values = { "hello", 42, true, new int[] { 1, 2, 3 }, 'a'  };
        foreach (object val in values)
        {
            
            switch (val)
            {
                case string str:
                    Console.WriteLine($"String of length {str.Length}");
                    break;
                case int i when i > 30:
                    Console.WriteLine($"Integer {i} is greater than 30");
                    break;
                case bool b:
                    Console.WriteLine($"Boolean value: {b}");
                    break;
                case int[] arr when arr.Length > 2:
                    Console.WriteLine($"Array with {arr.Length} elements");
                    break;
                case char c:
                    Console.WriteLine($"Character '{c}'");
                    break;
                default:
                    Console.WriteLine("Unknown type");
                    break;
            }
        }

        // int[] numbers = { 1, 2, 3 };

        // Console.WriteLine(numbers is [1, 2, 3]);  // True
        // Console.WriteLine(numbers is [1, 2, 4]);  // False
        // Console.WriteLine(numbers is [1, 2, 3, 4]);  // False
        // Console.WriteLine(numbers is [0 or 1, <= 2, >= 3]);  // True

        // 定义一个名为 person 的元组
        var person = (name: "Tom", age: 30);

        // 使用元组模式匹配来获取元组中的字段值
        switch (person)
        {
            case var p when p.age < 18:
                Console.WriteLine("Too young to vote.");
                break;
            case var p when p.age >= 18 && p.age <= 60:
                Console.WriteLine("Can vote.");
                break;
            case var p when p.age > 60:
                Console.WriteLine("Too old to vote.");
                break;
            default:
                Console.WriteLine("Invalid age.");
                break;
        }
    }
}
