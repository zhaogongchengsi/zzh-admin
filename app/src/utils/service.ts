import axios, { AxiosRequestConfig } from 'axios';
import type { AxiosInstance } from 'axios'
import { state } from "@/types/response"

const baseURL = "http://localhost:8000/api/v1/"
const http:AxiosInstance = axios.create({
  baseURL,
  timeout: 3000
})


// 添加请求拦截器
http.interceptors.request.use(function (config:AxiosRequestConfig) {
    // 在发送请求之前做些什么
    const token = localStorage.getItem('z_token')
    if (token) {
      config.headers['z_token'] = token;
    }
    
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
http.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    return response;
}, function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
});


export function Get (url: string) {
  return http.get(url)
}

export function Post (url: string, data: unknown) {
  return http.post(url, data)
}

type reqfun = (arg?: any) => Boolean

export function confirmStatus (state: state, fn?:reqfun) :boolean {
  if (state.code === 200) {
    if (fn) {
      fn()
    }
    return true
  }
  return false
}

export default http;