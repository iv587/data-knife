package redis

const keysInfoLuaScript = `
local keyList = {}
local count = tonumber(ARGV[1])
local dbSize = redis.call('DBSIZE')
if(dbSize <= count) then 
    if(ARGV[2]) then 
      keyList = redis.call('keys',ARGV[2])
    else
      keyList = redis.call('keys','*')
    end  
else
    if (ARGV[2]) then
      local cursor = 0
      local index = 0
      repeat
        local res = redis.call('SCAN', cursor, 'MATCH', ARGV[2], 'COUNT', count)
        for key,value in pairs(res[2]) do
            keyList[index+key] = value
        end
        index = #res[2]
        cursor = tonumber(res[1])
      until(cursor == 0)
    else
      keyList = redis.call('SCAN', 0, 'COUNT', count)[2]
    end
end
local keyInfoList = { }
local index = 1
for k,v in pairs(keyList) do
    local type = redis.call('TYPE', v)
    local ttl = redis.call('ttl', v)
    keyInfoList[index] = { v, ttl, type }
    index = index+1
end
return keyInfoList
`
