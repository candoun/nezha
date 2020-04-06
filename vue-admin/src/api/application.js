import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/v1/application/list',
    method: 'get',
    params: query
  })
}

export function fetchApplication(id) {
  return request({
    url: '/api/v1/application/detail/' + id,
    method: 'get'
  })
}


export function createApplication(data) {
  return request({
    url: '/api/v1/application',
    method: 'post',
    data
  })
}

export function updateApplication(data) {
  return request({
    url: '/application/update',
    method: 'post',
    data
  })
}
