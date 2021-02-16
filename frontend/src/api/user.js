import request from "@/utils/request";
import api from "@/api/index";

export const login = data => {
  return request({
    url: api.login.loginUrl,
    method: 'post',
    data
  })
}

export const logout = data => {
  return request({
    url: api.login.logoutUrl,
    method: 'post',
    data
  })
}

export const infoUser = data => {
  return request({
    url: api.user.info,
    method: 'post',
    data
  })
}

export const updateUser = data => {
  return request({
    url: api.user.update,
    method: 'post',
    data
  })
}
