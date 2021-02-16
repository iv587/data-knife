const Mock = require('mockjs2')
var data = {
  code: 1,
  msg: 'hello',
  data: {
    list: [
      {
        id: 1,
        name: 'hello',
        ip: '127.0.0.1',
        port: '6379'
      },
      {
        id: 2,
        name: 'hello1',
        ip: '127.0.0.2',
        port: '6379'
      },
      {
        id: 3,
        name: 'hello2',
        ip: '127.0.0.3',
        port: '6379'
      }
    ]
  }
}
Mock.mock('/api/redis/servers', 'post', data)

var dbData = {
  code: 1,
  msg: 'hello',
  data: {
    list: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
  }
}
Mock.mock('/api/redis/dbs', 'post', dbData)

var keyList = {
  code: 1,
  msg: 'hello',
  data: {
    list: [
      {
        key: 'test',
        type: 'String',
        ttl: '10'
      }
    ]
  }
}
Mock.mock('/api/redis/data', 'post', keyList)

