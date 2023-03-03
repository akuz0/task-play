# Как запустить.
1. Запустить db <br>
 > docker run --name go-db -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=digdb -d postgres:13.3

2. Накатить миграцию  <br>
 > migrate -source file://db/migrations/postgresql -database "postgres://root:pass@localhost:5434/digdb?sslmode=disable" up
3. Запуcтить
 > go run cmd/app/main.go -c cmd/app/config.yaml

Вводная по проетку:
Конфиг файл лежит по пути: 
 > cmd/app/config.yaml


Пример работы:
![example](https://s9.gifyu.com/images/ezgif.com-video-to-gife2f310a8ab1f320f.gif)

SQL:
select * from playlist
![image](https://user-images.githubusercontent.com/91601482/222738813-61c77c70-e608-444d-9d90-4b7cc8c2a485.png)
