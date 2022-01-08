import { Get, Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

export function createRole(data) {
  return new Promise((resolve, reject) => {
    Post("role/create_role", data)
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
