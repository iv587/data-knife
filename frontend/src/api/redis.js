import request from '@/utils/request'
import api from '@/api/index'

export const getRedisServerList = data => {
  return request({
    url: api.redis.serversUrl,
    method: 'post',
    data
  })
}

export const getRedisDbs = data => {
  return request({
    url: api.redis.dbsUrl,
    method: 'post',
    data
  })
}

export const getRedisKeyList = data => {
  return request({
    url: api.redis.keyListUrl,
    method: 'post',
    data
  })
}

export const getRedisKeyInfo = data => {
  return request({
    url: api.redis.keyInfoUrl,
    method: 'post',
    data
  })
}

export const updateRedisData = data => {
  return request({
    url: api.redis.updateUrl,
    method: 'post',
    data
  })
}

export const connectionInfo = data => {
  return request({
    url: api.redis.connectionInfoUrl,
    method: 'post',
    data
  })
}

export const redisInfo = data => {
  return request({
    url: api.redis.redisInfoUrl,
    method: 'post',
    data
  })
}

export const connectionUpdate = data => {
  return request({
    url: api.redis.connectionUpdateUrl,
    method: 'post',
    data
  })
}

export const deleteKey = data => {
  return request({
    url: api.redis.deleteKeyUrl,
    method: 'post',
    data
  })
}

export const testConnection = data => {
  return request({
    url: api.redis.connectionTestUrl,
    method: 'post',
    data
  })
}

export const deleteConnection = data => {
  return request({
    url: api.redis.connectionDeleteUrl,
    method: 'post',
    data
  })
}
