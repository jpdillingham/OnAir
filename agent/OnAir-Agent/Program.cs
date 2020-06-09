using OpenCvSharp;
using System;
using System.Diagnostics;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;

namespace OnAir_Agent
{
    class Program
    {
        static async Task Main(
            string[] processes,
            int[] cameras,
            string webhook)
        {
            processes ??= Array.Empty<string>();
            cameras ??= Array.Empty<int>();

            if (processes.Any(process => Process.GetProcessesByName(process).Length > 0))
            {
                Console.WriteLine($"At least one watched process is running");

                if (cameras.Length > 0) {
                    foreach (var camera in cameras)
                    {
                        Console.WriteLine($"Checking camera {camera}");

                        var cap = new VideoCapture(camera);
                        if (cap.Open(camera) && !cap.Grab())
                        {
                            Console.WriteLine($"Camera at index {camera} in use");
                            await InvokeWebhook(webhook, on: true);
                            return;
                        }
                    }
                }
                else {
                    await InvokeWebhook(webhook, on: true);
                    return;
                }
            }

            await InvokeWebhook(webhook, on: false);
        }

        private async static Task InvokeWebhook(string url, bool on)
        {
            url = $"{url}/{(on ? "on" : "off")}";

            Console.WriteLine($"POST: {url}");

            var response = await new HttpClient().PostAsync(url, null);

            Console.WriteLine(response);
        }
    }
}
