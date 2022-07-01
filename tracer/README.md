# Tracer Example
## Call Relations

```
┌────────────────┐    http call    ┌───────────────────────────┐      rpc call   ┌───────────────┐
│                ├─────────────────┤                           ├────────────────►│               │
│   Hertz-Client │                 │ Hertz-Server/Kitex-Client │                 │ Kitex-Server  │
│                │◄────────────────┤                           │◄────────────────┤               │
└────────────────┘                 └───────────────────────────┘                 └───────────────┘
```

## opentracing
### HOW-TO-RUN
1. install docker
2. run jaeger all-in-one   
   `sh jaeger_run.sh`
3. run Kitex server   
   `sh kitex_server_run.sh`
4. open another terminal and run Hertz server  
   `sh hertz_server_run.sh`
5. open another terminal and run Hertz client   
   `sh hertz_client_run.sh`
### MONITORING
You can then navigate to http://localhost:16686 to access the Jaeger UI. (You can visit [Monitor Jaeger](https://www.jaegertracing.io/docs/1.24/monitoring/) for details)   
For more information about hertz tracer, please click [tracer](https://www.cloudwego.io/zh/docs/hertz/tutorials/service-governance/tracing/)