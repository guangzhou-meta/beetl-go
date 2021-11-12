DemoTest
===
```sql
select #{page('*')} from sys_user 
where 1=1
-- @if(!isEmpty(name)){
 and name=#{name}
-- @}
-- @pageIgnoreTag(){
 order by createTime
-- @}
```

DemoTest2
===
```sql
select #{page('*')} from sys_user 
where 1=1
-- @if(!isEmpty(name)){
 and name=#{name}
-- @}
-- @pageIgnoreTag(){
 order by createTime
-- @}
```
