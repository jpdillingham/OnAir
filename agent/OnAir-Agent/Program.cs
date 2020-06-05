using OpenCvSharp;
using System;

namespace OnAir_Agent
{
    class Program
    {
        static void Main(string[] args)
        {
            var cap = new VideoCapture();

            for (int i = 0; i < 5; i++)
            {
                if (cap.Open(i) && !cap.Grab())
                {
                    Console.WriteLine($"Camera at index {i} in use");
                    return;
                }
            }

            Console.WriteLine($"No cameras in use.");
        }
    }
}
