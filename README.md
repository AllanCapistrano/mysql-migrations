# MySQL Migrations

Executa migrações em bancos de dados [MySQL](https://www.mysql.com/) que estão sendo executados em um [container Docker](https://hub.docker.com/_/mysql).

## Documentação

Comandos:

- [Migrate](#migrate)
- [Rollback](#rollback)
- [Dump](#dump)
- [Clear](#clear)

### Migrate

Realiza uma migração nos bancos de dados a partir de um arquivo `.sql` ou de uma query SQL. Durante o processo, são feitas cópias dos bancos de dados para caso seja necessário realizar o rollback, realizado através do comando 'mm rollback'.

```powershell
mm migrate my_migration.sql
```

#### Flags

##### --database | -D

Realiza a migração somente no banco de dado especificado.

```powershell
mm migrate my_migration.sql --database database_1

mm migrate my_migration.sql -D database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
mm migrate my_migration.sql --database database_1 --database database_2

mm migrate my_migration.sql -D database_1 -D database_2
```

##### --databases

Realiza a migração somente nos bancos de dados especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
mm migrate my_migration.sql --databases database_1,database_2
```

> [!IMPORTANT]
> Caso sejam utilizadas as flags `--database` e `--databases` no mesmo comando, somente o(s) banco(s) de dados especificado(s) pela última flag serão considerados.

##### --sql | -S

Especifica uma query SQL para realizar a migração em vez de utilizar um arquivo `.sql`.

```powershell
mm migrate --sql "CREATE TABLE users (id INT)"

mm migrate -S "CREATE TABLE users (id INT)"
```

> [!NOTE]
> As *queries* são validadas antes de serem executadas.

##### --no-database
Realiza a migração em todos bancos de dados, exceto no especificado.

```powershell
mm migrate my_migration.sql --no-database database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
mm migrate my_migration.sql --no-database database_1 --no-database database_2
```

##### --no-databases
Realiza a migração em todos bancos de dados, exceto nos especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
mm migrate my_migration.sql --no-databases database_1,database_2
```

> [!TIP]  
> As flags podem ser combinadas em um só comando
> ```powershell
> mm migrate --sql "CREATE TABLE users (id INT)" --databases database_1,database_2
> ```

### Rollback

Realiza o rollback da migração mais recente a partir do último arquivo criado no diretório '.rollback'. Caso nenhum arquivo seja encontrado, o rollback não poderá ser realizado. Ao final do processo o arquivo de rollback é removido.

```powershell
mm rollback
```

##### --database | -D

Realiza o rollback somente no banco de dado especificado.

```powershell
mm rollback --database database_1

mm rollback -D database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
mm rollback --database database_1 --database database_2

mm rollback -D database_1 -D database_2
```

##### --databases

Realiza o rollback somente nos bancos de dados especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
mm rollback --databases database_1,database_2
```

> [!IMPORTANT]
> Caso sejam utilizadas as flags `--database` e `--databases` no mesmo comando, somente o(s) banco(s) de dados especificado(s) pela última flag serão considerados.

##### --no-database
Realiza o rollback em todos bancos de dados, exceto no especificado.

```powershell
mm rollback --no-database database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
mm rollback --no-database database_1 --no-database database_2
```

##### --no-databases
Realiza o rollback em todos bancos de dados, exceto nos especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
mm rollback --no-databases database_1,database_2
```

### Dump

Realiza o dump de todos os bancos de dados (baseados na *whitelist* e *blacklist*, se estiverem preenchidas) organizando-os em diferentes diretórios.

```powershell
mm dump
```

##### --database | -D

Realiza o dump somente do banco de dado especificado.

```powershell
mm dump --database database_1

mm dump -D database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
mm dump --database database_1 --database database_2

mm dump -D database_1 -D database_2
```

##### --databases

Realiza o dump somente dos bancos de dados especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
mm dump --databases database_1,database_2
```

> [!IMPORTANT]
> Caso sejam utilizadas as flags `--database` e `--databases` no mesmo comando, somente o(s) banco(s) de dados especificado(s) pela última flag serão considerados.

##### --no-database
Realiza o dump de todos bancos de dados, exceto no especificado.

```powershell
mm dump --no-database database_1
```

Para múltiplos bancos de dados, utilize a flag mais de uma vez.

```powershell
mm dump --no-database database_1 --no-database database_2
```

##### --no-databases
Realiza o dump de todos bancos de dados, exceto nos especificados. Para múltiplos bancos de dados, utilize **vírgulas** para separá-los.

```powershell
mm dump --no-databases database_1,database_2
```
### Clear

Remove todos os arquivos de rollback que estão presentes no diretório '.rollback'

```powershell
mm clear
```