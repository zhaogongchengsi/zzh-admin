import axios from "axios";
import { ElLoading } from "element-plus";
import router from "@/routers/index";
let loading = null;

const baseURL = "http://localhost:8000/api/v1/";
const requestWhiteList = new Set(["/verify_code", "/user/login"]); // 是否挂token 的请求地址
const loadingWhiteList = new Set(["/verify_code"]); // 是否开启 loading

const Axios = axios.create({
  baseURL,
  timeout: 5000,
});
//  添加请求拦截
Axios.interceptors.request.use(
  function (config) {
    if (requestWhiteList.has(config.url)) {
      return config;
    }
    const token = window.localStorage.getItem("token");
    if (token) {
      config.headers = {
        z_token: token,
      };
    } else {
      router.push("/login");
    }
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    loading.close();
    return Promise.reject(error);
  }
);

// 添加响应拦截器
Axios.interceptors.response.use(
  function (response) {
    if (loading) {
      setTimeout(function () {
        loading.close();
      }, 200);
    }
    if (response.status !== 200) {
      return;
    }
    return response.data;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    loading.close();
    return Promise.reject(error);
  }
);

export function Get(url, load = true) {
  if (load) {
    loading = ElLoading.service();
  }
  return Axios.get(url);
}

export function Post(url, data, load = true) {
  if (load) {
    loading = ElLoading.service();
  }
  return Axios.post(url, data);
}
// 上传文件到本地服务器
export function UpLoad(url, file, load = true) {
  if (load) {
    loading = ElLoading.service();
  }
  let fromData = new FromData();
  fromData.append("file", file);
  return Axios.post(url, fromData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
}
