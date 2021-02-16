export default {
  redis: {
    serversUrl: '/redis/connection/list',
    dbsUrl: '/redis/db/list',
    keyListUrl: '/redis/key/list',
    keyInfoUrl: '/redis/key/info',
    updateUrl: '/redis/key/update',
    connectionInfoUrl: '/redis/connection/info',
    connectionUpdateUrl: '/redis/connection/update',
    connectionTestUrl: '/redis/connection/test',
    connectionDeleteUrl: '/redis/connection/delete',
    redisInfoUrl: '/redis/info',
    deleteKeyUrl: '/redis/key/delete'
  },
  login: {
    loginUrl: '/login',
    logoutUrl: '/logout'
  },
  user: {
    info: '/user/info',
    update: '/user/update'
  }
}
