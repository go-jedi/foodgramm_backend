### generate swagger:
- `swag init -g cmd/app/main.go`

### migrations:

#### create:
- `migrate create -ext sql -dir migrations -seq users_table`
- `migrate create -ext sql -dir migrations -seq user_excluded_products_table`
- `migrate create -ext sql -dir migrations -seq user_create_fn`
- `migrate create -ext sql -dir migrations -seq recipes_table`
- `migrate create -ext sql -dir migrations -seq recipes_create_index`
- `migrate create -ext sql -dir migrations -seq user_free_recipes_table`
- `migrate create -ext sql -dir migrations -seq user_free_recipes_create_index`
- `migrate create -ext sql -dir migrations -seq subscriptions_table`
- `migrate create -ext sql -dir migrations -seq subscription_history_table`
- `migrate create -ext sql -dir migrations -seq subscription_create_fn`
- `migrate create -ext sql -dir migrations -seq users_create_index`
- `migrate create -ext sql -dir migrations -seq subscriptions_exists_fn`
- `migrate create -ext sql -dir migrations -seq subscriptions_get_by_telegram_id_fn`
- `migrate create -ext sql -dir migrations -seq recipe_create_fn`
- `migrate create -ext sql -dir migrations -seq recipe_of_days_table`
- `migrate create -ext sql -dir migrations -seq promo_codes_table`
- `migrate create -ext sql -dir migrations -seq promo_code_uses_table`
- `migrate create -ext sql -dir migrations -seq promo_code_create_fn`

#### execute:
- `migrate -database postgresql://admin:test@localhost:54321/foodgrammm_db?sslmode=disable -path migrations up`
- `migrate -database postgresql://admin:test@localhost:54321/foodgrammm_db?sslmode=disable -path migrations down`

#### build application:
- `go build -ldflags="-s -w" -trimpath -buildvcs=false -o app cmd/app/main.go`

#### run application in systemd:
- `cd /etc/systemd/system`
- `создать openai-service.service`
- `sudo systemctl daemon-reload`
- `sudo systemctl start openai-service.service`
- `sudo systemctl status openai-service.service`
- `sudo systemctl enable openai-service.service`

#### Включить порт:
- `sudo ufw allow 50051/tcp`
- `sudo ufw reload`
- если при выполнении команды: sudo ss -tuln | grep 50051 у вас показывается:
tcp    LISTEN  0       4096         127.0.0.1:50051        0.0.0.0:*
, то это указывает, что сервис будет доступен только внутри сервера через localhost. 
- Если нужно, чтобы можно было отправлять запросы из внешних источников:
- в сервисе указываем при запуске http сервера :50051
- выполняем sudo ss -tuln | grep 50051 и должно быть в ответе такой результат:
  tcp    LISTEN  0       4096           0.0.0.0:50051        0.0.0.0:*

#### remove all local branch without main:
- `git branch | grep -v "main" | xargs git branch -D`