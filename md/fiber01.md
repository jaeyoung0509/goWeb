```go
app := fiber.New()

app.Prefork = true // enable prefork

app.Get("/", func(c *fiber.Ctx) {
  c.Send(fmt.Sprintf("Hi, I'm worker #%v", os.Getpid()))
  // => Hi, I'm worker #16858
  // => Hi, I'm worker #16877
  // => Hi, I'm worker #16895
})

app.Listen(8080)
```

  * what is  ```prefork```
  ```
    프리포크가 뭔가요?
     prefork 기능을 활성화하면 동일한 포트에서 수신 대기하는 여러 go 프로세스가 생성됩니다. 
     Nginx에는 소켓 샤딩에 대한 훌륭한 기사가 있습니다. 
     이 사진은 같은 기사에서 가져온 것입니다 👇
    ```

