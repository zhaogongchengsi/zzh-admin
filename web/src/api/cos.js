import { Get, Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

function get_temporary_key(data) {
  return new Promise(function (resolve, reject) {
    Post("/cos/get_temporary_key")
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
