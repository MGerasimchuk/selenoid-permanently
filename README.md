Работает:
```bash
docker stop $(docker ps -q) ;\
cd autostart &&\
docker-compose up -d &&\
cd .. &&\
go run main.go
```

Не работает:
```bash
docker stop $(docker ps -q) ;\
cd start-permanently &&\
docker-compose up -d &&\
cd .. &&\
go run main.go
```