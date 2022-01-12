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
select * from sys_user 
where 1=1
-- @if(!isEmpty(name)){
 and name=#{name}
-- @}
-- @if(!isEmpty(phone)||!isEmpty(email)){
 and (
    -- @if(!isEmpty(phone)){
        phone = #{phone}
    -- @}
    -- @if(!isEmpty(email)){
        -- @if(!isEmpty(phone)){
            or
        -- @}
        email = #{email}
    -- @}
    )
-- @}
```

