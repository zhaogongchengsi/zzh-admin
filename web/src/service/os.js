import { ElLoading } from "element-plus";
// import router from "@/routers/index";
let loading = null;
import { JudgeRequestStatus } from "@/utils";
import cos from "cos-js-sdk-v5";
import { Post } from "./index.js";
import { ElMessage } from "element-plus";

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

export async function createCosInstance(data) {
  try {
    let temp_key = await get_temporary_key(data);
    const cosInstance = new cos({
      getAuthorization: function (options, callback) {
        const { Credentials, ExpiredTime, StartTime } = temp_key;
        callback({
          TmpSecretId: Credentials.TmpSecretId,
          TmpSecretKey: Credentials.TmpSecretKey,
          SecurityToken: Credentials.Token,
          StartTime: StartTime,
          ExpiredTime: ExpiredTime,
        });
      },
    });
    return cosInstance;
  } catch (err) {
    console.error("创建实例", err);
  }
}
