import { ElLoading } from "element-plus";
// import router from "@/routers/index";
let loading = null;
import { JudgeRequestStatus } from "@/utils";
import cos from "cos-js-sdk-v5";
import { Post } from "./index.js";
import { ElMessage } from "element-plus";
import { Post } from "@/service/index.js";

export function get_temporary_key(data) {
  return new Promise(function (resolve, reject) {
    Post("/cos/get_temporary_key", data, false)
      .then(function (response) {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (code, msg) => {
            ElMessage.error(msg);
            reject({});
          }
        );
      })
      .catch(function (err) {
        reject(err);
      });
  });
}

export function createCosInstance(data) {
  return new cos({
    getAuthorization: async function (options, callback) {
      try {
        let temp_key = await get_temporary_key(data);
        const { Credentials, ExpiredTime, StartTime } = temp_key;
        callback({
          ...Credentials,
          StartTime: StartTime,
          ExpiredTime: ExpiredTime,
        });
      } catch (err) {
        ElMessage.error(`权限获取失败，联系管理员`);
      }
    },
  });
}
