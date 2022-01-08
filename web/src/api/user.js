import { Get, Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

export function createUser(data) {
  return new Promise((resolve, reject) => {
    Post("/user/register", data)
      .then((response) => {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (_, msg) => {
            ElMessage.error(msg);
            reject(msg);
          }
        );
      })
      .catch((err) => {
        reject(err);
      });
  });
}
