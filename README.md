# CNX Migrations

TODO

## Documentação

### Migrate

Realiza uma migração nos bancos de dados a partir de um arquivo `.sql` ou de uma query SQL. Durante o processo, são feitas cópias dos bancos de dados para caso seja necessário realizar o rollback, realizado através do comando 'cnx rollback'.

```powershell
cnx migrate my_migration.sql
```

#### Flags

##### --database | -D

Realiza a migração somente no banco de dado especificado.

```powershell
cnx migrate my_migration.sql --database database_1

cnx migrate my_migration.sql -D database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
cnx migrate my_migration.sql --database database_1 --database database_2

cnx migrate my_migration.sql -D database_1 -D database_2
```

##### --databases

Realiza a migração somente nos bancos de dados especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
cnx migrate my_migration.sql --databases database_1,database_2
```

##### --sql | -S

Especifica uma query SQL para realizar a migração em vez de utilizar um arquivo `.sql`.

```powershell
cnx migrate --sql "CREATE TABLE users (id INT)"

cnx migrate -S "CREATE TABLE users (id INT)"
```

> [!NOTE]
> As queries são validadas antes de serem executadas.

> [!TIP]  
> As flags podem ser combinadas em um só comando
> ```powershell
> cnx migrate --sql "CREATE TABLE users (id INT)" --databases database_1,database_2
> ```

> [!IMPORTANT]
> Caso sejam utilizadas as flags `--database` e `--databases` no memso comando, somente o(s) banco(s) de dados especificado(s) pela última flag serão considerados.
