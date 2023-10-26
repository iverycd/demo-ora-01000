# ORA-01000: maximum open cursors exceeded

## A Demo For Scene reproduction

![img.png](img.png)


```sql
select SQL_TEXT,CURSOR_TYPE,LAST_SQL_ACTIVE_TIME from V$OPEN_CURSOR where sid= and USER_NAME='ADMIN';
```


![img_1.png](img_1.png)