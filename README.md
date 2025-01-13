# golang-redis-example

- SQL code to add a table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fio VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- Running redis without password 
```bash
docker run -d --name redis_noauth -p 6379:6379 redis
redis-cli -h 127.0.0.1 -p 6379
```

- Running redis with password 
```bash
docker run -d --name redis_auth -p 6379:6379 -e "REDIS_PASSWORD=my_password" redis redis-server --requirepass my_password
redis-cli -h 127.0.0.1 -p 6379 -a my_password
```

- Redis commands
```bash
keys *      # see all keys
get my_key  # view the value of the key
del my_key  # delete key from redis
flushall    # clear all data
```
