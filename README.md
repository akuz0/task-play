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


https://user-images.githubusercontent.com/91601482/222741415-e017d329-994b-48c3-8f8f-1403af02a14b.mp4



SQL:
select * from playlist
![image](https://user-images.githubusercontent.com/91601482/222738813-61c77c70-e608-444d-9d90-4b7cc8c2a485.png)
