using Aspire.Hosting.Dapr;

var builder = DistributedApplication.CreateBuilder(args);

builder.AddProject<Projects.checkout>("checkout")
    .WithDaprSidecar(new DaprSidecarOptions
    {
        AppId = "checkout"
    });

builder.AddProject<Projects.order_processor>("order-processor")
    .WithDaprSidecar("order-processor");

builder.Build().Run();
