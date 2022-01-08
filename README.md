# Transport-Pipe

A way to transport a unix pipe over TCP

### Serve a pipe (By default at 0.0.0.0:8899)
```
go run . -mode server -b 1M < '/home/stargazer/Videos/[1080P Full風] Hand in Hand - Hatsune Miku 初音ミク Project DIVA Arcade English lyrics Romaji subtitles [9SKA6PmcLuQ].webm'
```

### Consume it
```
go run . -address 127.0.0.1:8899 -b 1M | ffplay -
```

Very alpha. Under devlopment.
